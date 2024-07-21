package services

import (
	"encoding/json"
	"log"

	"github.com/br4tech/go-log-process/internal/domain/entities"
	"github.com/br4tech/go-log-process/internal/port"
)

type LogService struct {
	message port.IMessage
}

func NewLogService(message port.IMessage) *LogService {
	return &LogService{
		message: message,
	}
}

func (service *LogService) Create(log *entities.Log) error {
	logJSON, err := json.Marshal(log)
	if err != nil {
		return err
	}

	if err := service.message.Publish("logs", string(logJSON)); err != nil {
		return err
	}

	return nil
}

func (service *LogService) ListAll() error {
	message, err := service.message.Consume("logs")
	if err != nil {
		return err
	}

	for msg := range message {
		var logEntity entities.Log
		if err := json.Unmarshal(msg.Value, &logEntity); err != nil {
			return err
		}

		log.Printf("Received log: %+v\n", logEntity)
	}

	return nil
}
