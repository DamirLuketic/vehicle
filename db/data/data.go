package data

import (
	"encoding/json"
	"log"
	"os"
)

func GetDBData() []VehicleJSON {
	file, err := os.ReadFile("db/data/init/VehicleInfo.json")
	if err != nil {
		log.Fatalf("Error while read data. Error: %s", err.Error())
	}
	var data []VehicleJSON
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatalf("Unmarshal data error. Error: %s", err.Error())
	}
	return data
}
