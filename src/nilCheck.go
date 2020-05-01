package src

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func nilCheck(Err *discordgo.Message) {
	if Err != nil {
		log.Println("Error: ", Err)
		return
	}
}
