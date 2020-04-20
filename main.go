package main

import (
	"change-status-go/command"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

var (
	StopBot = make(chan bool)
)
func main() {
	var discord, err = discordgo.New()
	secretEnv := os.Getenv("SECRET")
	discord.Token = secretEnv
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
