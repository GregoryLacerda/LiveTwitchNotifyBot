package domain

import (
	"time"

	"golang.org/x/oauth2/clientcredentials"
	"golang.org/x/oauth2/twitch"
)

var (
	clientID     = "oqlos1q8kicuq7a4ytr4aviogb2ef7"
	clientSecret = "etjq5ck3ue1wye6vnzfoa9ogcudhl6"
	Oauth2Config = &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     twitch.Endpoint.TokenURL,
	}
)

var (
	BotToken     = "MTI5MDUxODc1NDU2ODA0ODc1Mg.G8g1uN.NhzQbcJ_PIFbE3SzUcWWYcdf3DHdeOmM7wdM9I"
	BotChannelId = "1290523488712392751"
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
