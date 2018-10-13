package actions

import (
	"events-scrapper/models"
	"log"
	"os"
	"strings"

	mgo "gopkg.in/mgo.v2"
)

func getDBSession() (*mgo.Session, error) {
	conn := os.Getenv("CONN_STRING")
	session, err := mgo.Dial(conn)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func UpdateEventCollection(events []models.Event) error {
	db, err := getDBSession()
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer db.Close()

	for _, e := range events {
		err = db.DB("quehaysd").C("event").Insert(e)
		if err != nil {
			if !strings.Contains(err.Error(), "name_1_type_1_place_1 dup key") {
				return err
			}
		}
	}

	return nil
}
