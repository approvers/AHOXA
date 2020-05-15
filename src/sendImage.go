package src

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"strings"
)

func SendImage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	if!strings.HasPrefix(message.Content, "#") {
		return
	}
	GenImage(message.Content)

	file, Err := os.Open("sample.jpeg")
	if Err != nil {
		log.Println(Err)
	}

	_, Err = session.ChannelFileSend(message.ChannelID, "sample.jpeg", file)
	if Err != nil {
		log.Println(Err)
	}
}