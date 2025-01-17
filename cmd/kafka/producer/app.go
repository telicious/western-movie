package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/google/uuid"
	"github.com/hamba/avro/v2"
	"log"
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
	cfg := sarama.NewConfig()
	cfg.ClientID = "western-producer"
	cfg.Version = sarama.V3_3_2_0
	cfg.Metadata.AllowAutoTopicCreation = true
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	cfg.Producer.Retry.Max = 30
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Net.MaxOpenRequests = 1

	if err := cfg.Validate(); err != nil {
		log.Fatal(err)
	}

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
	producer, err := sarama.NewSyncProducer([]string{":49816"}, cfg)
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
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("[Message Sent] ", "topic:", "western-movies", " - key:", 1, " - msg:", data, " - partition:", partition, " - offset:", offset)
}
