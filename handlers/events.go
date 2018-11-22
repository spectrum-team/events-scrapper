package handlers

import (
	"encoding/json"
	"errors"
	"events-scrapper/models"
	"fmt"
	"io/ioutil"
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

func (e *EventHandler) QueryEvents(w http.ResponseWriter, r *http.Request) {

	var query interface{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(400)
		return
	}

	err = json.Unmarshal(body, &query)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(400)
		return
	}

	fmt.Println("The payload: ", query)

	events := make([]*models.Event, 0)
	err = e.Db.C("event").Find(query).All(&events)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(events)
}
