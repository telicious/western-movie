package main

import (
	"context"
	"log"
	"time"

	"github.com/IBM/sarama"
)

type consumerGroupHandler struct{}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("Received message: %s\n", string(msg.Value))
		// Process the message as per your requirement here
		sess.MarkMessage(msg, "")
	}
	return nil
}

func main() {
	brokers := []string{"localhost:9093"}
	groupID := "consumer-group"

	config := sarama.NewConfig()
	config.Version = sarama.V2_0_0_0 // specify appropriate Kafka version
	config.Consumer.Offsets.AutoCommit.Enable = true
	config.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	ctx := context.Background()

	for {
		err := consumerGroup.Consume(ctx, []string{"post-likes"}, consumerGroupHandler{})
		if err != nil {
			log.Panicf("Error from consumer: %v", err)
		}
	}
}
