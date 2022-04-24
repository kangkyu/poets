package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
	env "github.com/Netflix/go-env"
)

type AppConfig struct {
	Kafka struct {
		URL           string `env:"KAFKA_URL,required=true"`
		// TrustedCert   string `env:"KAFKA_TRUSTED_CERT,required"`
		// ClientCertKey string `env:"KAFKA_CLIENT_CERT_KEY,required"`
		// ClientCert    string `env:"KAFKA_CLIENT_CERT,required"`
		// Prefix        string `env:"KAFKA_PREFIX"`
		Topic         string `env:"KAFKA_TOPIC,default=poems"`
		// ConsumerGroup string `env:"KAFKA_CONSUMER_GROUP,default=poems-go"`
	}
}

type SubmitBody struct {
	Message string `json:"message"`
}

type Application struct {
	Producer sarama.AsyncProducer
	Topic string
}

func main() {
	var appconfig AppConfig
	_, err := env.UnmarshalFromEnviron(&appconfig)
	if err != nil {
		log.Fatal(err)
	}

	brokers := []string{appconfig.Kafka.URL}
	topic := appconfig.Kafka.Topic

	producer := NewProducer(brokers)
	defer producer.Close()

	app := Application{
		Producer: producer,
		Topic: topic,
	}

	http.HandleFunc("/submit", app.SubmitHandler)
	log.Println("Listen on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (app Application) SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only", http.StatusMethodNotAllowed)
		return
	}

	var submit SubmitBody
	err := json.NewDecoder(r.Body).Decode(&submit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%s\n", submit.Message)

	msg := submit.Message
	app.Producer.Input() <- &sarama.ProducerMessage{
		Topic: app.Topic,
		Value: sarama.StringEncoder(msg),
	}
	w.WriteHeader(http.StatusCreated)
}

func NewProducer(brokers []string) sarama.AsyncProducer {
	config := sarama.NewConfig()
	producer, err := sarama.NewAsyncProducer(brokers, config)

	if err != nil {
		panic(err)
	}
	return producer
}
