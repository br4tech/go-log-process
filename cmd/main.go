package main

import (
	"fmt"
	"log"
	"time"

	"github.com/br4tech/go-log-process/internal/adapter"
	"github.com/br4tech/go-log-process/internal/domain/entities"
	"github.com/br4tech/go-log-process/internal/domain/services"
)

func main() {
	kafkaServer := []string{"localhost:9092"}
	topic := "logs"
	adapter := adapter.NewKafkaAdapter(kafkaServer, topic)
	service := services.NewLogService(adapter)

	for i := 1; i <= 10; i++ {
		logEntity := entities.Log{
			Timestamp: time.Now(),
			Level:     "error",
			Message:   fmt.Sprintf("Erro numero: %d", i),
		}
		err := service.Create(&logEntity)
		if err != nil {
			log.Fatalf("Ocorreu um erro: %s", err)
		}
	}
}
