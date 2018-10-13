package main

import (
	"events-scrapper/actions"
	"events-scrapper/handlers"
	"log"
	"net/http"
	"os"

	gorillah "github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	mgo "gopkg.in/mgo.v2"
)

// "mongodb://testquehay:testquehay@ds117540.mlab.com:17540/quehaysd"
func getDBSession(conn string) (*mgo.Session, error) {
	session, err := mgo.Dial(conn)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func main() {

	// CRON Jobs
	job := actions.NewJob()
	job.GetConcertData("8h")

	conn := os.Getenv("CONN_STRING")

	db, err := getDBSession(conn)
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
