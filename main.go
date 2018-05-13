package main

import (
	"events-scrapper/actions"
	"events-scrapper/handlers"
	"events-scrapper/scrappers"
	"log"
	"net/http"
	"os"

	gorillah "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	mgo "gopkg.in/mgo.v2"
)

func getDBSession() (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://testquehay:testquehay@ds117540.mlab.com:17540/quehaysd")
	if err != nil {
		return nil, err
	}

	return session, nil
}

func main() {
	events := scrappers.Scrape()
	err := actions.UpdateEventCollection(events)
	if err != nil {
		log.Fatal(err)
	}

	db, err := getDBSession()
	if err != nil {
		log.Fatal(err)
	}

	e, err := handlers.NewEventHandler(db.DB("quehaysd"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/events", e.GetAllEvents).Methods("GET")

	listen := os.Getenv("PORT")

	if listen == "" {
		listen = "9000"
	}

	if err := http.ListenAndServe(":"+listen, gorillah.CombinedLoggingHandler(os.Stdout, router)); err != nil {
		log.Fatal(err)
	}
}
