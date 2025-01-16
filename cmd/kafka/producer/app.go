package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/hamba/avro/v2"
)

// WesternMovie - Struktur für das AVRO-Schema
type WesternMovie struct {
	ID          string   `avro:"id"`
	Title       string   `avro:"title"`
	Director    string   `avro:"director"`
	ReleaseYear int      `avro:"releaseYear"`
	Starring    []string `avro:"starring"`
}

func main() {
	// Kafka-Config
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0
	config.Producer.Return.Successes = true

	// AVRO-Schema
	schema, err := avro.Parse(`{
        "namespace": "com.western",
        "type": "record",
        "name": "WesternMovie",
        "fields": [
            {"name": "id", "type": "string"},
            {"name": "title", "type": "string"},
            {"name": "director", "type": "string"},
            {"name": "releaseYear", "type": "int"},
            {"name": "starring", "type": {"type": "array", "items": "string"}}
        ]
    }`)
	if err != nil {
		log.Fatal(err)
	}

	// Producer-Erstellung
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, config)
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Close()

	// UUID-Generator
	id := uuid.New().String()

	westernMovie := WesternMovie{
		ID:          id,
		Title:       "Der gute, der schlechte und der hässliche",
		Director:    "Sergio Leone",
		ReleaseYear: 1966,
		Starring:    []string{"Clint Eastwood", "Lee Van Cleef", "Aldo Giuffrè"},
	}

	data, err := avro.Marshal(schema, westernMovie)
	if err != nil {
		log.Fatal(err)
	}

	msg := &sarama.ProducerMessage{
		Topic: "western-movies",
		Value: sarama.ByteEncoder(data),
		Key:   sarama.StringEncoder(uuid.NewString()),
	}
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Nachricht mit ID '%s' erfolgreich an Kafka gesendet!\n", id)
}
