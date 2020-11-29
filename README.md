# golang_mongo_kafka

## Golang:
   Go is a procedural programming language. It was developed in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google but launched in 2009 as an open-source programming language. Programs are assembled by using packages, for efficient management of dependencies. This language also supports environment adopting patterns alike to dynamic languages.

## Kafka:
   Kafka is often used in real-time streaming data architectures to provide real-time analytics. Since Kafka is a fast, scalable, durable, and fault-tolerant publish-subscribe messaging system, Kafka is used in use cases where JMS, RabbitMQ, and AMQP may not even be considered due to volume and responsiveness.    

## Kafka Environment Setup

    $ cd /<download path>/kafka-2.6.0-src
    $ ./gradlew jar -PscalaVersion=2.13.2
    $ bin/zookeeper-server-start.sh config/zookeeper.properties
    $ bin/kafka-server-start.sh config/server.properties

## Execution Steps
    $ git clone https://github.com/saravase/golang_mongo_kafka.git
    $ cd golang_mongo_kafka
    $ make rest-to-kafka
    $ http://localhost:9090/plant

        {
  		    "id": 1,
  	        "name": "Graphs",
  	        "category": "Fruit",
  	        "price": 300.00,
  	        "description": "Healthy Fruit"
        }
    $ make kafka-to-mongo
