package src

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"regexp"
	"strings"
)

const (
	prefixMorseDecode = "DM"
	prefixMorseEncode = "EM"
	space             = " "
)

var alphabetTable = map[string]string{
	"a":     ".-",
	"b":     "-...",
	"c":     "-.-.",
	"d":     "-..",
	"e":     ".",
	"f":     "..-.",
	"g":     "--.",
	"h":     "....",
	"i":     "..",
	"j":     ".---",
	"k":     "-.-",
	"l":     ".-..",
	"m":     "--",
	"n":     "-.",
	"o":     "---",
	"p":     ".--.",
	"q":     "--.-",
	"r":     ".-.",
	"s":     "...",
	"t":     "-",
	"u":     "..-",
	"v":     "...-",
	"w":     "-..-",
	"x":     "-..-",
	"z":     "--..",
	".":     ".--.-.",
	",":     "--..--",
	":":     "---...",
	"?":     "..--..",
	"'":     ".----.",
	"-":     "-....-",
	"(":     "-.--.",
	")":     "-.--.-",
	"/":     "-..-.",
	"=":     "-...-.",
	"+":     ".-.-.-",
	"\"":    ".-..-.",
	"*":     "-..-",
	"@":     ".--.-.",
	"amend": "........",
}

func decode(sentence string) (string, error) {
	var response string
	reg := regexp.MustCompile(`[ \t+]`)
	sentence = reg.ReplaceAllString(sentence, " ")
	for _, part := range strings.Split(sentence, " ") {
		if part == space {
			response += " "
			continue
		}
		for alphabet, morse := range alphabetTable {
			if part == morse {
				response += alphabet
				break
			}
			return "[]", fmt.Errorf("Not such a code: %s", part)
		}
	}
	return response, nil
}

func DecodeMorse(session *discordgo.Session, message *discordgo.MessageCreate) {
	if (message.Author.ID == session.State.User.ID) || message.Author.Bot {
		return
	}
	if !strings.HasPrefix(message.Content, prefixMorseDecode) {
		return
	}
	sentence := message.Content[2:]
	decodeResult, Err := decode(sentence)
	if Err != nil {
		log.Println("Failed to decode:", Err)
		return
	}
	_, Err = session.ChannelMessageSend(message.ChannelID, decodeResult)
	if Err != nil {
		log.Println("Error at ChannelMessageSend:", Err)
		return
	}
}
