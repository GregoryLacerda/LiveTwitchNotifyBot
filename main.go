package main

import (
	"fmt"
	"gregorylaceda/livechangegamenotify/domain"
	"gregorylaceda/livechangegamenotify/twitch"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {

	actualGame := make(map[string]string)

	sessionDisc, err := discordgo.New("Bot " + domain.BotToken)
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

		token, err := twitch.GetToken(domain.Oauth2Config)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Access token: %s\n", token)

		streamsResponse, err := twitch.GetStreams(token)
		if err != nil {
			log.Fatal(err)
		}
		if len(streamsResponse.Data) == 0 {
			log.Println("No streams are live right now.")
			return
		}

		for _, stream := range streamsResponse.Data {

			if stream.UserLogin == "yoda" && stream.GameName != "LeagueOfLegends" && stream.GameName != actualGame[stream.UserLogin] {
				sessionDisc.ChannelMessageSend(domain.BotChannelId, fmt.Sprintf("Yoda is streaming %s\n", stream.GameName))
				actualGame[stream.UserLogin] = stream.GameName
			}

			if stream.UserLogin == "gaules" && stream.GameName != "Counter-Strike: Global Offensive" && stream.GameName != actualGame[stream.UserLogin] {
				sessionDisc.ChannelMessageSend(domain.BotChannelId, fmt.Sprintf("Gaules is streaming %s\n", stream.GameName))
				actualGame[stream.UserLogin] = stream.GameName
			}
			println("end all verifications")
		}

		time.Sleep(10 * time.Minute)
	}

}
