package main

import (
	"golang_mongo_kafka/db"
)

func main() {

	//Create MongoDB session
	ms := db.NewMongoStore()
	ms.ReceiveFromKafka()

}
