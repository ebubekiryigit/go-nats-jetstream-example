package main

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"go-nats-jetstream-example/config"
	"go-nats-jetstream-example/models"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func publishReviews(js nats.JetStreamContext) {
	reviews, err := getReviews()
	if err != nil {
		log.Println(err)
		return
	}

	for _, oneReview := range reviews {

		// create random message intervals to slow down
		r := rand.Intn(1500)
		time.Sleep(time.Duration(r) * time.Millisecond)

		reviewString, err := json.Marshal(oneReview)
		if err != nil {
			log.Println(err)
			continue
		}

		// publish to REVIEWS.rateGiven subject
		_, err = js.Publish(config.SubjectNameReviewCreated, reviewString)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("Publisher  =>  Message: %s\n", oneReview.Text)
		}
	}
}

func getReviews() ([]models.Review, error) {
	rawReviews, _ := ioutil.ReadFile("./reviews.json")
	var reviewsObj []models.Review
	err := json.Unmarshal(rawReviews, &reviewsObj)

	return reviewsObj, err
}
