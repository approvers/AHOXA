package src

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"strings"
)

const (
	prefix      = "%"
	colorPrefix = "#"
)

type messageContext struct {
	s *discordgo.Session
	m *discordgo.MessageCreate
}

func (cxt *messageContext) messageSend(message string) (Err error) {
	_, Err = cxt.s.ChannelMessageSend(cxt.m.ChannelID, message)
	if Err != nil {
		log.Println("failed send message: ", Err)
		return
	}
	return
}

func (cxt *messageContext) fileSend(fileName string, data io.Reader) (Err error) {
	_, Err = cxt.s.ChannelFileSend(cxt.m.ChannelID, fileName, data)
	if Err != nil {
		log.Println("failed send file: ", Err)
		return
	}
	return
}

func morseCodeOperation(mode string, codeType string) (answerSentence string, Err error) {
	switch mode {
	case "decode":
		answerSentence, Err = DecodeMorse(codeType)
		return
	default:
		return "", fmt.Errorf("Error at morseCodeOperation: No such operation.")
	}
}

func morseAction(command string, codeSentence string, context messageContext) {
	contentText, Err := morseCodeOperation(command, codeSentence)
	if Err != nil {
		log.Println("failed decode morse: ", Err)
		return
	}
	Err = context.messageSend(contentText)
	if Err != nil {
		log.Println("failed send message: ", Err)
		return
	}
}

func colorAction(command string, context messageContext) {
	fileData, Err := GenerateImage(command)
	if Err != nil {
		log.Println("failed to genarateImage: ", Err)
		return
	}
	Err = context.fileSend("unkonow.jpeg", fileData)
	if Err != nil {
		log.Println("failed file send: ", Err)
		return
	}
}

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	Context := messageContext{
		session,
		message,
	}

	if message.Author.ID == session.State.User.ID {
		return
	}
	if !strings.HasPrefix(message.Content, prefix) {
		return
	}
	commandBody := strings.Split(message.Content, " ")
	switch commandBody[0][len(prefix):] {
	case "color":
		if len(commandBody) != 2 {
			Err := Context.messageSend("コマンドの形式が間違っています。`%help`を参照してください。")
			if Err != nil {
				log.Println("failed send message: ", Err)
				return
			}
			return
		}
		colorAction(commandBody[1], Context)
	case "morse":
		if len(commandBody) < 2 {
			Err := Context.messageSend("コマンドの形式が間違っています。`%help`を参照してください。")
			if Err != nil {
				log.Println("failed send message: ", Err)
				return
			}
			return
		}
		morseAction(commandBody[1], Context.m.Content[len("%morse decode"):], Context)
	default:
		contentText := fetchMessage(commandBody[0][len(prefix):])
		Err := Context.messageSend(contentText)
		if Err != nil {
			log.Println("failed send message: ", Err)
			return
		}
	}

}

func BootNotify(s *discordgo.Session, _ *discordgo.Ready) {
	// BootNotify is sending message when this bot is booted.
	_, Err := s.ChannelMessageSend("699941274484080660", "BootBot! <@!622077711309078529>")
	if Err != nil {
		log.Println("Boot failed: ", Err)
	}
}
