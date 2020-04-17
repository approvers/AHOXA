package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"change-status-go/secret"
	"strings"
)

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}

	switch {
		case strings.HasPrefix(m.Content, fmt.Sprintf("%s", hello)):
			sendMessage(s, c, helloWorld)

		case strings.HasPrefix(m.Content, fmt.Sprintf("%s", usage)):
			sendMessage(s, c, usageSentence)

		case strings.HasPrefix(m.Content, fmt.Sprintf("%s", callName)):
			sendMessage(s, c, m.Member.Nick)

		case strings.HasPrefix(m.Content, fmt.Sprintf(uid)):
			sendMessage(s, c, m.Author.ID)

		case strings.HasPrefix(m.Content, fmt.Sprintf("%s", unko)): {
			s.GuildMemberNickname(secret.GuildID,m.Message.Member.User.ID,":poop:")
		}
	}
}


func sendMessage(s *discordgo.Session, c *discordgo.Channel, msg string) {
	_, err := s.ChannelMessageSend(c.ID, msg)

	log.Println(">>> " + msg)
	if err != nil {
		log.Println("Error sending message: ", err)
	}
}
