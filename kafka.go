package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

type MsgSender struct {
	Producer *sarama.Producer
}

//Returns a new Sender
func NewMsgSender(brokers []string) *MsgSender {
	client, err := sarama.NewClient("client_id", brokers, sarama.NewClientConfig())
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to Kafka Brokers...")
	}
	producer, err := sarama.NewProducer(client, nil)
	if err != nil {
		panic(err)
	}
	s := &MsgSender{producer}
	return s
}

//Send msg to specific topic
func (s *MsgSender) Send(topic string, msg []byte) {
	s.Producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(msg)}
}
