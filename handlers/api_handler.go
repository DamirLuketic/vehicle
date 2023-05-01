package handlers

import (
	"crypto/subtle"
	"errors"
	"github.com/DamirLuketic/vehicle/config"
	"github.com/DamirLuketic/vehicle/db"
	"log"
	"net/http"
)

type APIHandler interface {
	GetVehicles(w http.ResponseWriter, r *http.Request)
	CreateVehicles(w http.ResponseWriter, r *http.Request)
	DeleteVehicles(w http.ResponseWriter, r *http.Request)
}

type APIHandlerImpl struct {
	db          db.DataStore
	APIUsername string
	APIPassword string
}

func NewApiHandler(db db.DataStore, c *config.ServerConfig) APIHandler {
	return &APIHandlerImpl{
		db:          db,
		APIUsername: c.APIUser,
		APIPassword: c.APIPassword,
	}
}

func (h *APIHandlerImpl) authorized(r *http.Request) bool {
	username, password, isOk := r.BasicAuth()
	if isOk {
		return subtle.ConstantTimeCompare([]byte(h.APIUsername), []byte(username)) == 1 &&
			subtle.ConstantTimeCompare([]byte(h.APIPassword), []byte(password)) == 1
	}
	return false
}

func (h *APIHandlerImpl) validation(method string, w http.ResponseWriter, r *http.Request) error {
	if !h.authorized(r) {
		w.WriteHeader(http.StatusUnauthorized)
		errorString := "unauthorized"
		w.Write([]byte(errorString))
		log.Println(errorString)
		return errors.New(errorString)
	}
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		errorString := "method not allowed"
		w.Write([]byte(errorString))
		log.Println(errorString)
		return errors.New(errorString)
	}
	return nil
}
