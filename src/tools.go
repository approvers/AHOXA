package src

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

var prefix = "%"

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	cmd := strings.Split(m.Content, " ")[0][len(prefix):]
	Err, _ := s.ChannelMessageSend(m.ChannelID, fetchMessage(cmd))
	nilCheck(Err)
}

func BootNotify(s *discordgo.Session, m *discordgo.Ready) {
	// BootNotify is sending message when this bot is booted.
	Err, _ := s.ChannelMessageSend("699941274484080660", "BootBot! <@!622077711309078529>")
	nilCheck(Err)
}

func nilCheck(Err *discordgo.Message) {
	if Err != nil {
		log.Println("Error: ", Err)
		return
	}
}
