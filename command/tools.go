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

	switch {
		case strings.HasPrefix(m.Content, fmt.Sprintf("%s", hello)):
			sendMessage(s, c, helloWorld)

		case strings.HasPrefix(m.Content, fmt.Sprintf("%s", usage)):
			sendMessage(s, c, usageSentence)

		case strings.HasPrefix(m.Content, fmt.Sprintf("%s", callName)): {
			if m.Member.Nick != "" {
				sendMessage(s, c, m.Member.Nick)
			} else {
				sendMessage(s, c, m.Author.Username)
			}

		}

		case strings.HasPrefix(m.Content, fmt.Sprintf(uid)):
			sendMessage(s, c, m.Author.ID)

		case strings.HasPrefix(m.Content, fmt.Sprintf("%s", status)): {
			if strings.Split(m.Content, " ")[0] == "!status" {
				Emoji := m.Content[len("!status"):]
				Err := s.GuildMemberNickname(m.GuildID, m.Author.ID,m.Author.Username + Emoji)
				if fmt.Sprint(Err) == `HTTP 403 Forbidden, {"message": "Missing Permissions", "code": 50013}` {
					sendMessage(s, c, forbidden)
				}
				if Err != nil {
					fmt.Println(Err)
					sendMessage(s, c, wrong)
				}
				sendMessage(s, c, Emoji + "ですね。" + m.Author.Username + "、行ってらっしゃい。")
			} else {
				sendMessage(s, c, advise)
			}

		}
		case strings.HasPrefix(m.Content, fmt.Sprintf("%s", reset)):{
			Err := s.GuildMemberNickname(m.GuildID, m.Author.ID, m.Author.Username)
			if Err != nil {
				fmt.Println(Err)
				sendMessage(s, c, wrong)
			}
			sendMessage(s, c, notify + "、"+ m.Author.Username)
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
