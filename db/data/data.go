package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func GetDBData() []VehicleJSON {
	file, err := ioutil.ReadFile("db/data/init/VehicleInfo.json")
	if err != nil {
		log.Fatalf("Error while read data. Error: %s", err.Error())
	}
	var data []VehicleJSON
	json.Unmarshal(file, &data)
	return data
}
