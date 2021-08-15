package twitter

import (
	tw "github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"log"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func GetClient(creds *Credentials) (*tw.Client, error) {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := tw.NewClient(httpClient)

	verifyParams := &tw.AccountVerifyParams{
		SkipStatus:   tw.Bool(true),
		IncludeEmail: tw.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("Username: %s\n", user.Name)
	return client, nil
}

func GetCredentials(consumerKey string, consumerSecret string, accessToken string, accessTokenSecret string) *Credentials {
	return &Credentials{
		ConsumerKey:       consumerKey,
		ConsumerSecret:    consumerSecret,
		AccessToken:       accessToken,
		AccessTokenSecret: accessTokenSecret,
	}
}
