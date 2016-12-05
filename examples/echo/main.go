package main

import (
	"log"
	"os"

	"strconv"

	"github.com/aoisensi/go-discordapp/discord"
)

func main() {
	if os.Getenv("DISCORD_BOT_TOKEN") == "" {
		log.Fatalln("DISCORD_BOT_TOKEN is empty.")
	}
	if os.Getenv("DISCORD_CHANNEL_ID") == "" {
		log.Fatalln("DISCORD_CHANNEL_ID is empty.")
	}
	id, err := strconv.ParseInt(os.Getenv("DISCORD_CHANNEL_ID"), 10, 63)
	cid := discord.Snowflake(id)
	if err != nil {
		log.Fatalln(err)
	}
	cli := discord.NewBotClient(nil, os.Getenv("DISCORD_BOT_TOKEN"))
	gw, err := discord.NewGateway()
	if err != nil {
		log.Fatalln(err)
	}
	var user discord.User
	gw.AddHandler(func(e *discord.EventReady) {
		user = e.User
	})
	gw.AddHandler(func(e *discord.EventMessageCreate) {
		if e.Author.ID == user.ID {
			return
		}
		if e.ChannelID != cid {
			return
		}
		_, err := cli.Channel(cid).CreateMessage(e.Content)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(e.Content)
	})
	log.Println("Start.")
	if err := gw.Start(os.Getenv("DISCORD_BOT_TOKEN")); err != nil {
		log.Fatalln(err)
	}
}
