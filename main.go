package main

import (
	"fmt"
	"gregorylaceda/livechangegamenotify/config"
	"gregorylaceda/livechangegamenotify/twitch"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	token, err := twitch.GetToken(cfg)
	if err != nil {
		log.Fatal(err)
	}

	actualGame := make(map[string]string)

	sessionDisc, err := discordgo.New("Bot " + cfg.DISCORD_BOT_TOKEN)
	if err != nil {
		log.Fatal(err)
	}

	sessionDisc.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sessionDisc.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sessionDisc.Close()

	for {
		streamsResponse, err := twitch.GetStreams(cfg, token)
		if err != nil {
			//TODO: valid expired token to renew
			log.Fatal(err)
		}
		if len(streamsResponse.Data) == 0 {
			log.Println("No streams in live right now.")
			return
		}

		for _, stream := range streamsResponse.Data {

			if stream.GameName != actualGame[stream.UserLogin] {
				sessionDisc.ChannelMessageSend(cfg.DISCORD_BOT_CHANNEL_ID, fmt.Sprintf("%s is streaming %s\n", stream.UserLogin, stream.GameName))
				actualGame[stream.UserLogin] = stream.GameName
			}
			fmt.Println("finished - "+time.Now().Format(time.RFC1123), stream.UserLogin)
		}

		time.Sleep(10 * time.Minute)
	}

}
