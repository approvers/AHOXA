package src

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

const prefix = "%"

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	cmd := strings.Split(m.Content, " ")[0][len(prefix):]
	msg, Err := FetchMessage(cmd)
	nilCheck(Err)
	_, Err = s.ChannelMessageSend(m.ChannelID, msg)
	nilCheck(Err)
}

func BootNotify(s *discordgo.Session, m *discordgo.Ready) {
	// BootNotify is sending message when this bot is booted.
	_, Err := s.ChannelMessageSend("699941274484080660", "BootBot! <@!622077711309078529>")
	nilCheck(Err)
}

func nilCheck(Err error) {
	if Err != nil {
		log.Println("Error: ", Err)
		return
	}
}
