package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Setup sarama producer and muxer
func main() {
	sender := NewMsgSender([]string{"localhost:9092"})
	handler := &KafkaHandler{sender}
	router := mux.NewRouter()
	router.HandleFunc("/topics/{topic:[\\w]+}/messages", handler.Handle).Methods("POST")
	http.Handle("/", router)
	log.Println("Listening on 3000...")
	http.ListenAndServe(":3000", nil)
}
