package command

import (
	"change-status-go/sentence"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

const (
	errorBadRequest = `HTTP 400 Bad Request, {"nick": ["Must be 32 or fewer in length."]}`
	errorForbidden = `HTTP 403 Forbidden, {"message": "Missing Permissions", "code": 50013}`
	prefix = "%"
)

func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, Err := s.State.Channel(m.ChannelID)
	if Err != nil {
		log.Println("Error getting channel: ", Err)
		return
	}
	if !strings.HasPrefix(m.Content, prefix) {
		return
	}
	commandName := strings.Split(m.Content, " ")[0][len(prefix):]
	switch commandName {
		case status:
			operator := strings.Split(m.Content, " ")[1]
			Emoji := strings.Split(m.Content, " ")[2]

			if operator == "add" {
				Err = s.GuildMemberNickname(m.GuildID, m.Author.ID, m.Member.Nick + Emoji)
			} else if operator == "rm" {
				if !strings.Contains(m.Member.Nick, Emoji) {
					return
				}
				Err = s.GuildMemberNickname(m.GuildID, m.Author.ID, m.Member.Nick[:(len(m.Member.Nick) - len(Emoji))])
			}
			if Err != nil {
				if Err.Error() == errorForbidden {
					sendMessage(s, c, sentence.Forbidden)
					return
				}
				if Err.Error() == errorBadRequest {
					sendMessage(s, c, sentence.BadRequest)
					return
				}
				fmt.Println(Err)
				sendMessage(s, c, sentence.Wrong)
				return
			}
			sendMessage(s, c, fmt.Sprintf("%sですね。%s、行ってらっしゃい。", Emoji, m.Author.Username + Emoji))
		case hello:
			sendMessage(s, c, sentence.Hello)
		case usage:
			sendMessage(s, c, sentence.Usage)
		case callName:
			if m.Member.Nick != "" {
				sendMessage(s, c, m.Member.Nick)
			} else {
				sendMessage(s, c, m.Author.Username)
			}
		case uid:
			sendMessage(s, c, m.Author.ID)
	}
}


func sendMessage(s *discordgo.Session, c *discordgo.Channel, msg string) {
	_, Err := s.ChannelMessageSend(c.ID, msg)

	log.Println(">>> " + msg)
	if Err != nil {
		log.Println("Error sending message: ", Err)
	}
}
