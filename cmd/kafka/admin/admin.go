package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type topic struct {
	name        string
	partition   int32
	replication int16
	settings    map[string]*string
}

var (
	messageRetentionMinute = "60000"
	messageRetentionHour   = "3600000"
)

var topics = []topic{{
	name:        "western-topic-1",
	partition:   1,
	replication: 1,
	settings: map[string]*string{
		"retention.ms":        &messageRetentionMinute,
		"delete.retention.ms": &messageRetentionMinute,
	},
}}

func main() {
	config := sarama.NewConfig()
	config.Version = sarama.V3_3_2_0

	clusterAdmin, err := sarama.NewClusterAdmin([]string{":49816"}, config)
	if err != nil {
		log.Fatalln(err)
	}
	defer clusterAdmin.Close()

	clusterTopics, err := clusterAdmin.ListTopics()
	if err != nil {
		log.Fatalln(err)
	}

	for _, topic := range topics {
		if _, ok := clusterTopics[topic.name]; ok {
			continue
		}

		err := clusterAdmin.CreateTopic(topic.name, &sarama.TopicDetail{
			NumPartitions:     topic.partition,
			ReplicationFactor: topic.replication,
			ConfigEntries:     topic.settings,
		}, false)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("> TOPIC: %s PARTITION: %d REPLICA: %d\n", topic.name, topic.partition, topic.replication)
	}
}
