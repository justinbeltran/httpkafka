package main

import (
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

type KafkaHandler struct {
	MsgSender *MsgSender
}

//handler that sends message to producer's input channel
func (h *KafkaHandler) Handle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	topic := params["topic"]
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	h.MsgSender.Send(topic, body)
	w.WriteHeader(http.StatusNoContent)
}
