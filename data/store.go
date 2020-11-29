package data

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"gopkg.in/mgo.v2"
)

type MongoStore struct {
	Session    *mgo.Session
	Host       string
	Database   string
	Username   string
	Password   string
	Collection string
}

// Receive plant data from kafka
func (ms *MongoStore) ReceiveFromKafka() {

	fmt.Println("Start receiving from Kafka")
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "group-id-1",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"plants-topic1"}, nil)

	for {
		msg, err := c.ReadMessage(-1)

		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
		fmt.Printf("Received from Kafka %s: %s\n", msg.TopicPartition, string(msg.Value))
		plant := string(msg.Value)
		ms.savePlantToMongo(plant)
	}

	c.Close()
}

// Insert plant data into the mongo collection
func (ms *MongoStore) savePlantToMongo(plantString string) {

	fmt.Println("Save to MongoDB")
	col := ms.Session.DB(ms.Database).C(ms.Collection)

	//Save data into Job struct
	var plant Plant
	b := []byte(plantString)
	err := json.Unmarshal(b, &plant)
	if err != nil {
		panic(err)
	}

	//Insert job into MongoDB
	errMongo := col.Insert(plant)
	if errMongo != nil {
		panic(errMongo)
	}

	fmt.Printf("Saved to MongoDB : %s", plantString)

}

// Insert plant data into kafka topic store
func SavePlantToKafka(plant *Plant) {

	fmt.Println("save to kafka")

	jps, err := json.Marshal(plant)

	js := string(jps)
	fmt.Print(js)

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		panic(err)
	}

	// Produce messages to topic (asynchronously)
	topic := "plants-topic1"
	for _, word := range []string{string(js)} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}
}
