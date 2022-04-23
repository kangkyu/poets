package main

import (
	"github.com/Shopify/sarama"
	"github.com/joeshaw/envdecode"
)

type AppConfig struct {
	Kafka struct {
		URL           string `env:"KAFKA_URL,required"`
		// TrustedCert   string `env:"KAFKA_TRUSTED_CERT,required"`
		// ClientCertKey string `env:"KAFKA_CLIENT_CERT_KEY,required"`
		// ClientCert    string `env:"KAFKA_CLIENT_CERT,required"`
		// Prefix        string `env:"KAFKA_PREFIX"`
		Topic         string `env:"KAFKA_TOPIC,default=poems"`
		// ConsumerGroup string `env:"KAFKA_CONSUMER_GROUP,default=poems-go"`
	}
}

func main() {
	appconfig := AppConfig{}
	envdecode.MustDecode(&appconfig)
	brokers := []string{appconfig.Kafka.URL}
	topic := appconfig.Kafka.Topic

	config := sarama.NewConfig()
	producer, err := sarama.NewAsyncProducer(brokers, config)

	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := "Hello world"
	producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
}
