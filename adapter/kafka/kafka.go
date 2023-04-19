package kafka


import (
	"context"
	"github.com/segmentio/kafka-go"
	"lykafe/news/config"
)

type Kafka struct {
	r *kafka.Reader
}

func NewKafka() *Kafka {
	reader := kafka.NewReader(kafka.ReaderConfig{
		// Brokers:   []string{"kafka:9092"},
		Brokers:   config.KafkaBrokers(),
		GroupID:   config.KafkaGroupId(),
		Topic:     config.KafkaTopicUpdateUser(),
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	return &Kafka{
		r: reader,
	}
}

func (k *Kafka) ReadMessage(ctx context.Context) (key []byte, value []byte, err error) {
	m, err := k.r.ReadMessage(ctx)
	key = m.Key
	value = m.Value
	return
}
