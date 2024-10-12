package twitch

import (
	"context"
	"gregorylaceda/livechangegamenotify/config"
	"log"

	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

func GetToken(cfg *config.Config) (string, error) {
	oauth2Config := clientcredentials.Config{
		ClientID:     cfg.TWITCH_CLIENT_ID,
		ClientSecret: cfg.TWITCH_CLIENT_SECRET,
		TokenURL:     twitch.Endpoint.TokenURL,
	}

	token, err := oauth2Config.Token(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return token.AccessToken, nil
}
