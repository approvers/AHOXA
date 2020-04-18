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
		case status:
			Emoji := m.Content[8:]
			Err :=s.GuildMemberNickname(m.GuildID, m.Author.ID, m.Author.Username + Emoji)
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

		case reset:
			Err := s.GuildMemberNickname(m.GuildID, m.Author.ID, m.Author.Username)
			if Err != nil {
				fmt.Println(Err)
				sendMessage(s, c, sentence.Wrong)
				return
			}
			sendMessage(s, c, sentence.Notify + "、" + m.Author.Username)
	}
}


func sendMessage(s *discordgo.Session, c *discordgo.Channel, msg string) {
	_, Err := s.ChannelMessageSend(c.ID, msg)

	log.Println(">>> " + msg)
	if Err != nil {
		log.Println("Error sending message: ", Err)
	}
}
