package command

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Println("Error getting channel: ", err)
		return
	}
	commandName := strings.Split(m.Content, " ")[0][1:]
	switch commandName {
		case hello:
			sendMessage(s, c, helloWorld)
		case usage:
			sendMessage(s, c, usageSentence)
		case callName: {
			if m.Member.Nick != "" {
				sendMessage(s, c, m.Member.Nick)
			} else {
				sendMessage(s, c, m.Author.Username)
			}
		}
		case uid:
			sendMessage(s, c, m.Author.ID)
		case status: {
			Emoji := m.Content[len(status):]
			Err :=s.GuildMemberNickname(m.GuildID, m.Author.ID, m.Author.Username + Emoji)
			if fmt.Sprint(Err) == `HTTP 403 Forbidden, {"message": "Missing Permissions", "code": 50013}` {
				sendMessage(s, c, forbidden)
				return
			}
			if Err != nil {
				fmt.Println(Err)
				sendMessage(s, c, wrong)
				return
			}
			sendMessage(s, c, Emoji + "ですね。" + m.Author.Username + "、行ってらっしゃい。")

		}
		case reset: {
			Err := s.GuildNickname(m.guildID, m.Author.ID, m.Author.Username)
			if Err != nil {
				fmt.Println(Err)
				sendMessage(s, c, wrong)
				return
			}
			sendMessage(s, c, notify + "、" + m.Author.Username)
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
