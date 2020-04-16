package main

import (
	"change-status-go/command"
	"change-status-go/secret"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	StopBot = make(chan bool)
)
func main() {
	var discord, err = discordgo.New()
	discord.Token = secret.Token
	if err != nil {
		fmt.Println("Error logged in")
		fmt.Println(err)
	}

	discord.AddHandler(command.OnMessageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Listening...")
	<-StopBot
	return
}
