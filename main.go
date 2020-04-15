package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var (
	Token = "Bot Njk5OTM0MTQ5NzI0NzMzNTIw.XpcOmQ.4laodM1AugWukek0aPpl-glBfEU"
	BotName = "699934149724733520"
	StopBot = make(chan bool)
	Hello = "!hello"
)
func main() {
	var discord, err = discordgo.New()
	discord.Token = Token
	if err != nil {
		fmt.Println("Error logged in")
		fmt.Println(err)
	}

	discord.AddHandler(onMessageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Listening...")
	<-StopBot
	return
}

func onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}

	if strings.HasPrefix(m.Content, fmt.Sprintf("%s", Hello)) {
		sendMessage(s, c, "Hello world!")
	}
}

func sendMessage(s *discordgo.Session, c *discordgo.Channel, msg string) {
	_, err := s.ChannelMessageSend(c.ID, msg)

	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
