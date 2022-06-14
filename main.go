package main

import (
	"fmt"

	"github.com/stephen-mahon/discord-go-bot/bot"
	"github.com/stephen-mahon/discord-go-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.Start()

	<-make(chan struct{})
	return
}
