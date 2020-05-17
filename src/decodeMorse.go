package src

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

const prefixMorseDecode = "DM"
const prefixMorseEncode = "EM"

func decode(sentence string, wordSep string) (string, error) {
	var response string
	for _, s := range strings.Split(sentence, " ") {
		if s ==wordSep {
			response += " "
		}
		for key, value := range alphabetTable {
			if s != value {
				return "[]", fmt.Errorf("Not such a code: %s", s)
			}
			response += key

		}
	}
	return sentence, nil
}

func DecodeMorse(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}
	if !strings.HasPrefix(message.Content, prefixMorseDecode) {
		return
	}
	sentence := message.Content[2:]
	content, Err := decode(sentence, " ")
	if Err != nil {
		log.Println("Failed to decode:", Err)
		return
	}
	_, Err = session.ChannelMessageSend(message.ChannelID, content)
	if Err != nil {
		log.Println("Error at ChannelMessageSend:", Err)
		return
	}
}

var alphabetTable = map[string]string{
	"a"     : ".-",
	"b"     : "-...",
	"c"     : "-.-.",
	"d"     : "-..",
	"e"     : ".",
	"f"     : "..-.",
	"g"     : "--.",
	"h"     : "....",
	"i"     : "..",
	"j"     : ".---",
	"k"     : "-.-",
	"l"     : ".-..",
	"m"     : "--",
	"n"     : "-.",
	"o"     : "---",
	"p"     : ".--.",
	"q"     : "--.-",
	"r"     : ".-.",
	"s"     : "...",
	"t"     : "-",
	"u"     : "..-",
	"v"     : "...-",
	"w"     : "-..-",
	"x"     : "-..-",
	"z"     : "--..",
	"."     : ".--.-.",
	","     : "--..--",
	":"     : "---...",
	"?"     : "..--..",
	"'"     : ".----.",
	"-"     : "-....-",
	"("     : "-.--.",
	")"     : "-.--.-",
	"/"     : "-..-.",
	"="     : "-...-.",
	"+"     : ".-.-.-",
	"\""    : ".-..-.",
	"*"     : "-..-",
	"@"     : ".--.-.",
	"amend" : "........",
}
