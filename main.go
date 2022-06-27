package main

import (
	"log"
	"sync"
)

func main() {
	log.Println("Starting...")

	js, err := JetStreamInit()
	if err != nil {
		log.Println(err)
		return
	}

	// Let's assume that publisher and consumer are services running on different servers.
	// So run publisher and consumer asynchronously to see how it works
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		publishReviews(js)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		consumeReviews(js)
	}()

	wg.Wait()

	log.Println("Exit...")
}
