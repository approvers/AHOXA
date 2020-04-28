package src

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var (
	prefix = "%"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//Ping -> Pong
	if m.Author.ID == s.State.User.ID {
		return
	}
	if !strings.HasPrefix(m.Content, prefix) {
		return
	}

	var cmd = strings.Split(m.Content, " ")[0][len(prefix):]
	//temp := strings.Split(m.Content, " ")[1]
	Err, _ := s.ChannelMessageSend(m.ChannelID, fetchMessage(cmd))

	if Err != nil {
		fmt.Println("something wrong:", Err)
		return
	}

}
