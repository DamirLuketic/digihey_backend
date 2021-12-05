package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *APIHandler) GetVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result, err := h.db.GetVehicles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println("handlers.etVehicles Error: ", err.Error())
		return
	}
	vehiclesJSON := parseVehiclesDataToJSON(result)
	responseBody, _ := json.Marshal(vehiclesJSON)
	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("handlers.GetVehicles response write Error: ", err.Error())
	}
}

func (h *APIHandler) CreateVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	bodyBytes, err := ioutil.ReadAll(r.Body)
	reqBody := VehicleJSON{}
	err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println("handlers.CreatePermission: ", err.Error())
		return
	}
	if reqBody.Make == "" || reqBody.Model == "" || reqBody.Oid == "" || reqBody.Year == 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("handlers.CreatePermission: missing data in request body")
		return
	}
	vehicle := parseCreateVehicleEntry(reqBody.Make, reqBody.Model, reqBody.Oid, reqBody.Year)
	result, err := h.db.CreateVehicle(vehicle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println("CreateVehicles Error: ", err.Error())
		return
	}
	vehicleJSON := parseVehicleDataToJSON(result)
	responseBody, _ := json.Marshal(vehicleJSON)
	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("GetVehicles response write Error: ", err.Error())
	}
}

func (h *APIHandler) DeleteVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	queryParams := r.URL.Query()
	vehicleOID := queryParams.Get("oid")
	if vehicleOID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("handlers.DeleteVehicles: missing data in request")
		return
	}
	err := h.db.DeleteVehicle(vehicleOID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println("GetVehicles Error: ", err.Error())
		return
	}
}
