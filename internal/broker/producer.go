package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/commerce-app-demo/order-service/internal/models/orders"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/samborkent/uuidv7"
)

type Producer struct {
	Producer *kafka.Producer
}

type OrderPlaced struct {
	OrderID    int     `json:"order_id"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
	Timestamp  string  `json:"timestamp"`
}

func NewKafkaProducer() (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"client.id":         "orderProducer",
		"acks":              "all",
	})

	if err != nil {
		return nil, fmt.Errorf("error creating kafka producer: %w", err)
	}

	return &Producer{Producer: p}, nil
}

func (p *Producer) PublishOrderPlaced(ctx context.Context, order *orders.OrderEntity) error {
	orderPlaced := entityToEvent(order)
	orderPlacedEvent, err := json.Marshal(orderPlaced)

	if err != nil {
		return err
	}

	topic := "yeah"
	if err := p.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          orderPlacedEvent,
		Key:            []byte(uuidv7.New().String()),
		Timestamp:      time.Now(),
		TimestampType:  kafka.TimestampCreateTime,
		Opaque:         nil,
		Headers: []kafka.Header{
			kafka.Header{Key: "sourceService", Value: []byte("order-service")},
		},
		LeaderEpoch: nil,
	}, nil); err != nil {
		return err
	}
	return nil
}

func entityToEvent(order *orders.OrderEntity) *OrderPlaced {
	return &OrderPlaced{
		OrderID:    order.Id,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
		Timestamp:  order.CreatedAt,
	}
}
