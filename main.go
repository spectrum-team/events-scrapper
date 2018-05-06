package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

type Event struct {
	Name      string
	EventType string
	Date      string
	Place     string
}

func scrape() []Event {
	eventMap := make(map[int]Event)

	doc, err := goquery.NewDocument("http://www.conectate.com.do/eventos/")
	if err != nil {
		log.Fatal(err)
	}

	// Name of the event
	doc.Find("td.post-col-main div.post-main span.entry-title").Each(func(index int, item *goquery.Selection) {
		eventMap[index] = Event{Name: item.Text()}
	})

	// Type of event
	doc.Find("td.post-col-category").Each(func(index int, item *goquery.Selection) {
		currentEvent := eventMap[index]
		currentEvent.EventType = item.Text()
		eventMap[index] = currentEvent
	})

	// Date of event
	doc.Find("td.post-col-date span.entry-title").Each(func(index int, item *goquery.Selection) {
		currentEvent := eventMap[index]
		currentEvent.Date = item.Text()
		eventMap[index] = currentEvent
	})

	// Place of event
	doc.Find("td.post-col-location a").Each(func(index int, item *goquery.Selection) {
		currentEvent := eventMap[index]
		currentEvent.Place = item.Text()
		eventMap[index] = currentEvent
	})

	fmt.Println(len(eventMap))
	events := make([]Event, len(eventMap))

	for e, event := range eventMap {
		events[e] = event
	}

	return events
}

func main() {
	events := scrape()

	fmt.Println("The Events", events)

}
