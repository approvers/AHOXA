package main

import (
	command "change-status-go/src"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discordBrain, err := discordgo.New()
	if err != nil {
		panic(err)
	}

	discordToken := loadToken()
	if discordToken == "" {
		panic("no discord token exists.")
	}
	discordBrain.Token = discordToken

	discordBrain.AddHandler(command.MessageCreate)

	err = discordBrain.Open()
	if err != nil {
		panic(err)
	}
	defer discordBrain.Close()

	fmt.Println("Bot起動完了、命令待機中")
	discordBrain.AddHandlerOnce(command.BootNotify)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	return
}

func loadToken() string {
	return os.Getenv("DISCORD_TOKEN")
}
