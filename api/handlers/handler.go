package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *APIHandler) GetVehicles(w http.ResponseWriter, r *http.Request) {
	err := h.validation("GET", w, r)
	if err != nil {
		return
	}
	result, err := h.db.GetVehicles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println("handlers.etVehicles Error: ", err.Error())
		return
	}
	responseBody, _ := json.Marshal(result)
	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("handlers.GetVehicles response write Error: ", err.Error())
	}
}

func (h *APIHandler) CreateVehicles(w http.ResponseWriter, r *http.Request) {
	err := h.validation("POST", w, r)
	if err != nil {
		return
	}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	reqBody := CreateVehicleRequest{}
	err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("handlers.CreatePermission: ", err)
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
		log.Println("GetVehicles Error: ", err.Error())
		return
	}
	responseBody, _ := json.Marshal(result)
	_, err = w.Write(responseBody)
	if err != nil {
		log.Println("GetVehicles response write Error: ", err.Error())
	}
}

func (h *APIHandler) DeleteVehicles(w http.ResponseWriter, r *http.Request) {
	err := h.validation("DELETE", w, r)
	if err != nil {
		return
	}
	queryParams := r.URL.Query()
	vehicleOID := queryParams.Get("oid")
	if vehicleOID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("handlers.CreatePermission: missing data in request")
		return
	}
	err = h.db.DeleteVehicle(vehicleOID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println("GetVehicles Error: ", err.Error())
		return
	}
}
