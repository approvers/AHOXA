package main

import (
	"change-status-go/secret"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	discordBrain, err := discordgo.New()
	discordBrain.Token = secret.Token
	if err != nil {
		fmt.Println("Error logging in")
		fmt.Println(err)
	}

	discordBrain.AddHandler(messageCreate)

	err = discordBrain.Open()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Bot起動完了、命令待機中")
	sc := make(chan os.Signal,1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<- sc
	return
}
