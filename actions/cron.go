package actions

import (
	"events-scrapper/scrappers"
	"fmt"
	"time"

	"github.com/robfig/cron"
)

// Job is a struct used for handling
// scheduled actions using CRON
type Job struct {
	Cron *cron.Cron
}

// NewJob is a function used to initialized
// a Job struct with a Cron that can be reused
func NewJob() *Job {
	return &Job{
		Cron: cron.New(),
	}
}

// GetConcertData uses the Scrape function in the scrappers package to update
// the events data in the DB. This is for concert event data
func (j *Job) GetConcertData(duration string) {
	started := time.Now()
	fmt.Println("*** [*] CRON job 'GetConcertData' started ***")
	fmt.Printf("*** [*] CRON job 'GetConcertData' start time: %v ***\n", started)

	spec := fmt.Sprintf("@every %s", duration)
	j.Cron.AddFunc(spec, func() {
		events := scrappers.Scrape()

		err := UpdateEventCollection(events)
		if err != nil {
			ended := time.Now()
			fmt.Println("*** [*] CRON job 'GetConcertData' finished unexpectedly ***")
			fmt.Printf("*** [*] CRON job 'GetConcertData' Errors: [%v] ***\n", err)
			fmt.Printf("*** [*] CRON job 'GetConcertData' end time: %v ***\n", ended)
			fmt.Printf("*** [*] CRON job 'GetConcertData' time elapsed: %v ***\n", ended.Sub(started))
		}

		ended := time.Now()
		fmt.Println("*** [*] CRON job 'GetConcertData' finished succesfully ***")
		fmt.Printf("*** [*] CRON job 'GetConcertData' end time: %v ***\n", ended)
		fmt.Printf("*** [*] CRON job 'GetConcertData' time elapsed: %v ***\n", ended.Sub(started))
	})

	j.Cron.Start()
}
