package handlers

import (
	"encoding/json"
	"errors"
	"events-scrapper/models"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

type EventHandler struct {
	Db *mgo.Database
}

func NewEventHandler(db *mgo.Database) (*EventHandler, error) {

	if db == nil {
		return nil, errors.New("Missing database connection...")
	}

	return &EventHandler{
		Db: db,
	}, nil
}

func (e *EventHandler) GetAllEvents(w http.ResponseWriter, r *http.Request) {

	events := make([]*models.Event, 0)

	err := e.Db.C("event").Find(nil).All(&events)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(events)
}
