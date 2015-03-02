package main

import (
	"github.com/Shopify/sarama"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

type Sender struct {
	Producer *sarama.Producer
}

//handler that sends message to producer's input channel
func (s *Sender) Handler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	topic := params["topic"]
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	s.Producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(body)}
}
