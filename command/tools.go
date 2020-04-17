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

		case strings.HasPrefix(m.Content, fmt.Sprintf("%s", callName)):
			sendMessage(s, c, m.Member.Nick)

		case strings.HasPrefix(m.Content, fmt.Sprintf(uid)):
			sendMessage(s, c, m.Author.ID)

		case strings.HasPrefix(m.Content, fmt.Sprintf("%s", unko)): {
			Err := s.GuildMemberNickname(m.GuildID, m.Author.ID,"💩")
			if fmt.Sprint(Err) == "403 Forbidden" {
				sendMessage(s, c, "権限がないので変更できません。落ちぶれましょう。")
			}
			if Err != nil {
				fmt.Println(Err)
				sendMessage(s, c, "あなたのコードが間違っています")
			}
			sendMessage(s, c, "今から私は" + m.Member.Nick + "です。")
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
