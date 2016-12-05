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
	if err != nil {
		log.Fatalln(err)
	}
	cli := discord.NewBotClient(nil, os.Getenv("DISCORD_BOT_TOKEN"))
	_, err = cli.Channel(discord.Snowflake(id)).CreateMessage("Hello, world!")
	if err != nil {
		log.Fatalln(err)
	}
}
