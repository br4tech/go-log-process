package adapter

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/br4tech/go-log-process/internal/port"
	"github.com/segmentio/kafka-go"
)

type KafkaAdapter struct {
	Writer *kafka.Writer
}

func NewKafkaAdapter(kafkaBrokers []string, topic string) port.IMessage {
	// 1. Criar o tópico
	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaBrokers[0], topic, 0)
	if err != nil {
		log.Fatalf("failed to dial leader: %w", err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		log.Fatalf("failed to get controller: %w", err)
	}

	controllerConn, err := kafka.Dial("tcp", net.JoinHostPort(controller.Host, fmt.Sprint(controller.Port)))
	if err != nil {
		log.Fatalf("failed to dial controller: %w", err)
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{Topic: topic, NumPartitions: 1, ReplicationFactor: 1},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		log.Fatalf("failed to create topic: %w", err)
	}

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  kafkaBrokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	return &KafkaAdapter{Writer: w}
}

func (k *KafkaAdapter) Publish(topic string, message string) error {
	err := k.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(message),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to write messages: %w", err)
	}

	log.Printf(" [x] Published message to topic %s: %s\n", topic, message)
	return nil
}

func (k *KafkaAdapter) Consume(topic string) (<-chan kafka.Message, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  strings.Split(topic, ","),
		GroupID:  "consumer-group-id",
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	messageChan := make(chan kafka.Message)

	go func() {
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Printf("Error reading message: %s", err)
				close(messageChan)
				return
			}
			messageChan <- m
		}
	}()

	return messageChan, nil
}

func (k *KafkaAdapter) Close() error {
	if err := k.Writer.Close(); err != nil {
		return fmt.Errorf("failed to close writer: %w", err)
	}
	return nil
}
