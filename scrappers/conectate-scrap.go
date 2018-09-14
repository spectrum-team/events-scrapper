package scrappers

import (
	"events-scrapper/models"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/goodsign/monday"
)

// Scrape goes into the conectate webpage and gather the events listed there
func Scrape() []models.Event {
	eventMap := make(map[int]models.Event)

	log.Println("This is a test log")
	doc, err := goquery.NewDocument("http://www.conectate.com.do/eventos/")
	if err != nil {
		log.Fatal(err)
	}

	// Name of the event
	doc.Find("td.post-col-main div.post-main span.entry-title").Each(func(index int, item *goquery.Selection) {
		eventMap[index] = models.Event{Name: item.Text()}
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
		dateStr := strings.Replace(item.Text(), ",", "", -1)
		date, _ := monday.ParseInLocation("02 Jan 2006", dateStr, time.UTC, monday.LocaleEsES)
		currentEvent.Date = date
		eventMap[index] = currentEvent
	})

	// Place of event
	doc.Find("td.post-col-location a").Each(func(index int, item *goquery.Selection) {
		currentEvent := eventMap[index]
		currentEvent.Place = item.Text()
		eventMap[index] = currentEvent
	})

	events := make([]models.Event, len(eventMap))

	for e, event := range eventMap {
		events[e] = event
	}

	return events
}
