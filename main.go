package main

import (
	"events-scrapper/scrappers"
	"fmt"
)

func main() {
	events := scrappers.Scrape()
	fmt.Println("The Events", events)
}
