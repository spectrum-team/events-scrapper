package main

import (
	"events-scrapper/actions"
	"events-scrapper/scrappers"
	"log"
)

func main() {
	events := scrappers.Scrape()
	err := actions.UpdateEventCollection(events)
	if err != nil {
		log.Fatal(err)
	}
}
