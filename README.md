httpkafka
=========

Proof-of-Concept to expose HTTP endpoint to send JSON messages to a Kafka topic. Also a chance for me to write some more Go.

## Test it out!

### Download and Install Kafka
http://kafka.apache.org/documentation.html#quickstart

### Install Go Dependencies
```
$ go get github.com/gorilla/mux
$ go get github.com/Shopify/sarama
```

### Build and Start
```
$ go build
$ ./httpkafka 
```

### Test sending message to "test" kafka topic
```
$ curl -X POST http://localhost:3000/topics/test/messages -d "testing 123"
$ KAFKA_HOME/bin/kafka-console-consumer.sh --zookeeper localhost:2181 --topic test --from-beginning
```
