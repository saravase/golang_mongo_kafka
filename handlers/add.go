package handlers

import (
	"net/http"

	"github.com/saravase/golang_mongo_kafka/data"
)

type KeyPlant struct{}

// AddPlant used to insert the new plant data into the datastore
func (plant *Plant) AddPlant(res http.ResponseWriter, req *http.Request) {
	plant.logger.Printf("[DEBUG] Add the plant data")

	plantData := &data.Plant{}
	err := data.FromJSON(plantData, req.Body)
	if err != nil {
		plant.logger.Println("[ERROR] Deserializing plant", err)
		res.WriteHeader(http.StatusBadRequest)
		data.ToJSON("Add Failed", res)
		return
	}
	plant.logger.Printf("[DEBUG] Added plant data: %#v\n", plantData)

	data.SavePlantToKafka(plantData)

}
