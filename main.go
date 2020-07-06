package main

import (
	command "change-status-go/src"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discordBrain, err := discordgo.New()

	discordToken := loadToken()
	if discordToken == "" {
		log.Println("Error: no discord token exists.")
		return
	}
	discordBrain.Token = discordToken
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
	}

	discordBrain.AddHandler(command.MessageCreate)

	err = discordBrain.Open()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Bot起動完了、命令待機中")
	discordBrain.AddHandlerOnce(command.BootNotify)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	return
}

func loadToken() string {
	token := os.Getenv("DISCORD_TOKEN")
	return token
}
