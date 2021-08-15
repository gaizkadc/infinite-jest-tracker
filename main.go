package main

import (
	"flag"
	"fmt"
	"github.com/gaizkadc/infinite-jest-tracker/twitter"
	"github.com/gaizkadc/infinite-jest-tracker/utils"
	"log"
)

func main() {
	var progress = flag.Int("progress", 0, "reading progress")
	var consumerKey = flag.String("consumer-key", "", "twitter consumer key")
	var consumerSecret = flag.String("consumer-secret", "", "twitter consumer secret")
	var accessToken = flag.String("access-token", "", "twitter access token")
	var accessTokenSecret = flag.String("access-token-secret", "", "twitter access token secret")
	flag.Parse()

	log.Printf("starting infinite-jest-tracker\n")
	log.Printf("progress: %d\n", *progress)

	twitterCredentials := twitter.GetCredentials(*consumerKey, *consumerSecret, *accessToken, *accessTokenSecret)
	if twitterCredentials == nil {
		log.Printf("could not retrieve twitter credentials\n")
	}

	twClient, tErr := twitter.GetClient(twitterCredentials)
	if tErr != nil {
		log.Printf("error creating twitter client: %s\n", tErr.Error())
	}

	tweetCaption := utils.CreateCaption(*progress)
	if *progress == 100 {
		tweetCaption += fmt.Sprint("\n\nPor. Puto. Fin.")
	}
	log.Println(tweetCaption)

	twClient.Statuses.Update(tweetCaption, nil)
}


