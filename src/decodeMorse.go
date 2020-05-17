package src

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"regexp"
	"strings"
)

const (
	prefixMorseDecode = "%decode"
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

func searchTable(morseInput string) (string, bool) {
	for alphabet, morse := range alphabetTable {
		if morseInput == morse {
			return alphabet, true
		}
	}

	return "", false
}

func decode(sentence string) (response string, Err error) {
	reg := regexp.MustCompile(`[ \t]+`)
	sentence = reg.ReplaceAllString(sentence, " ")
	log.Printf(sentence)
	for _, part := range strings.Split(sentence, " ") {

		alphabet, found := searchTable(part)

		if !found {
			return "", errors.New(fmt.Sprintf("Not found such code: %s", part))
		}

		response += alphabet
	}
	return
}

func DecodeMorse(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot {
		return
	}
	if !strings.HasPrefix(message.Content, prefixMorseDecode) {
		return
	}
	sentence := message.Content[len(prefixMorseDecode)+1:]
	decodeResult, Err := decode(sentence)
	if Err != nil {
		log.Println("Failed to decode:", Err)
		return
	}
	_, Err = session.ChannelMessageSend(message.ChannelID, decodeResult)
	if Err != nil {
		log.Println("Error at ChannelMessageSend:", Err)
	}
}
