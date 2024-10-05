package domain

import (
	"time"

	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

var (
	clientID     = "<YOU_CLIENT_ID"
	clientSecret = "<YOUR_Secret>"
	Oauth2Config = &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     twitch.Endpoint.TokenURL,
	}
)

var (
	BotToken     = "<YOUR_BOT_TOKEN>"
	BotChannelId = "<YOUR_CHANNEL_ID>"
)

type StreamResponse struct {
	Data []struct {
		ID           string    `json:"id"`
		UserID       string    `json:"user_id"`
		UserLogin    string    `json:"user_login"`
		UserName     string    `json:"user_name"`
		GameID       string    `json:"game_id"`
		GameName     string    `json:"game_name"`
		Type         string    `json:"type"`
		Title        string    `json:"title"`
		ViewerCount  int       `json:"viewer_count"`
		StartedAt    time.Time `json:"started_at"`
		Language     string    `json:"language"`
		ThumbnailURL string    `json:"thumbnail_url"`
		TagIds       []any     `json:"tag_ids"`
		Tags         []string  `json:"tags"`
		IsMature     bool      `json:"is_mature"`
	} `json:"data"`
}
