package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Setup sarama producer and muxer
func main() {
	log.Println("Starting up...")
	sender := NewMsgSender([]string{"localhost:9092"})
	handler := &KafkaHandler{sender}
	router := mux.NewRouter()
	router.HandleFunc("/topics/{topic:[\\w]+}/messages", handler.Handle).Methods("POST")
	http.Handle("/", router)
	log.Println("Waiting for requests on port 3000...")
	http.ListenAndServe(":3000", nil)
}
