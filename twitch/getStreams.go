package twitch

import (
	"encoding/json"
	"fmt"
	"gregorylaceda/livechangegamenotify/config"
	"gregorylaceda/livechangegamenotify/domain"
	"io"
	"net/http"
)

func GetStreams(cfg *config.Config, token string) (streamsRespomse domain.StreamResponse, err error) {

	url := fmt.Sprintf("https://api.twitch.tv/helix/streams?%s", cfg.TWITCH_CHANNELS_TO_MONITOR)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return domain.StreamResponse{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Client-Id", cfg.TWITCH_CLIENT_ID)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return domain.StreamResponse{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return domain.StreamResponse{}, err
	}

	err = json.Unmarshal(body, &streamsRespomse)
	if err != nil {
		return domain.StreamResponse{}, err
	}

	return streamsRespomse, nil
}
