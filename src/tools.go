package src

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

const prefix = "%"

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	cmd := strings.Split(m.Content, " ")[0][len(prefix):]
	Err, _ := s.ChannelMessageSend(m.ChannelID, fetchMessage(cmd))
	if Err != nil {
		log.Println("failed send message: ", Err)
	}
}

func BootNotify(s *discordgo.Session, m *discordgo.Ready) {
	// BootNotify is sending message when this bot is booted.
	_, Err := s.ChannelMessageSend("699941274484080660", "BootBot! <@!622077711309078529>")
	if Err != nil {
		log.Println("Boot failed: ", Err)
	}
}
