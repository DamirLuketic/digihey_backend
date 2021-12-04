package handlers

import (
	"crypto/subtle"
	"errors"
	"github.com/DamirLuketic/digihey_backend/config"
	"github.com/DamirLuketic/digihey_backend/db"
	"log"
	"net/http"
)

type APIHandler struct {
	db          db.DataStore
	APIUsername string
	APIPassword string
}

func NewApiHandler(db db.DataStore, c *config.ServerConfig) APIHandler {
	return APIHandler{
		db:          db,
		APIUsername: c.APIUser,
		APIPassword: c.APIPassword,
	}
}

func (h *APIHandler) authorized(r *http.Request) bool {
	username, password, isOk := r.BasicAuth()
	if isOk {
		return subtle.ConstantTimeCompare([]byte(h.APIUsername), []byte(username)) == 1 &&
			subtle.ConstantTimeCompare([]byte(h.APIPassword), []byte(password)) == 1
	}
	return false
}

func (h *APIHandler) validation(method string, w http.ResponseWriter, r *http.Request) error {
	if !h.authorized(r) {
		w.WriteHeader(http.StatusUnauthorized)
		errorString := "Unauthorized"
		w.Write([]byte(errorString))
		log.Println(errorString)
		return errors.New(errorString)
	}
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		errorString := "Method not allowed"
		w.Write([]byte(errorString))
		log.Println(errorString)
		return errors.New(errorString)
	}
	return nil
}
