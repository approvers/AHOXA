package src

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func BootNotify(s *discordgo.Session, m *discordgo.Ready) {
	Err, _ := s.ChannelMessageSend("699941274484080660", "BootBot! <@!622077711309078529>")
	if Err == nil {
		fmt.Println("Bot did not Booted.")
		return
	}
}
