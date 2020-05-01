package src

import (
	"github.com/bwmarrin/discordgo"
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

	cmd := strings.Split(m.Content, " ")[0][:len(prefix)]
	Err, _ := s.ChannelMessageSend(m.ChannelID, fetchMessage(cmd))
	nilCheck(Err)
}


