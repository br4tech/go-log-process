package port

import "github.com/segmentio/kafka-go"

type IMessage interface {
	Publish(topic string, message string) error
	Consume(topic string) (<-chan kafka.Message, error)
}
