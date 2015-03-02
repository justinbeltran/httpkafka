package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Setup sarama producer and muxer
func main() {
	client, err := sarama.NewClient("client_id", []string{"localhost:9092"}, sarama.NewClientConfig())
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to Kafka Broker...")
	}
	defer client.Close()

	producer, err := sarama.NewProducer(client, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	s := &Sender{producer}
	r := mux.NewRouter()
	r.HandleFunc("/topics/{topic:[\\w]+}/messages", s.Handler).Methods("POST")
	http.Handle("/", r)
	log.Println("Listening on 3000...")
	http.ListenAndServe(":3000", nil)
}
