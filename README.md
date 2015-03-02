httpkafka
=========

Proof-of-Concept to expose HTTP endpoint to send JSON messages to a Kafka topic

Start
-----

* Download and Install Kafka
http://kafka.apache.org/documentation.html#quickstart

* Install Dependencies
```
$ go get github.com/gorilla/mux
$ go get github.com/Shopify/sarama
```

* Run
```
$ go build
$ ./httpkafka 
```

* Test sending message to "test"
```
$ curl -X POST http://localhost:3000/topics/test/messages -d "testing 123"
$ KAFKA_HOME/bin/kafka-console-consumer.sh --zookeeper localhost:2181 --topic test --from-beginning
```
