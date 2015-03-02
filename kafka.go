package main

import (
	"github.com/Shopify/sarama"
	"log"
)

type MsgSender struct {
	Client   *sarama.Client
	Producer *sarama.Producer
}

//Returns a new Sender
func NewMsgSender(brokers []string) *MsgSender {
	client, err := sarama.NewClient("httpkafka", brokers, sarama.NewClientConfig())
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected to Kafka Brokers...")
	}
	producer, err := sarama.NewProducer(client, nil)
	if err != nil {
		panic(err)
	} else {
		log.Println("Created Kafka Producer")
	}
	s := &MsgSender{client, producer}
	return s
}

//Send msg to specific topic
func (s *MsgSender) Send(topic string, msg []byte) {
	s.Producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(msg)}
}
