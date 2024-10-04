package twitch

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"
)

func GetToken(oauth2Config *clientcredentials.Config) (string, error) {
	token, err := oauth2Config.Token(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return token.AccessToken, nil
}
