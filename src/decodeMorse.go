package src

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

var alphabetTable = map[string]string{
	"a":             ".-",
	"b":             "-...",
	"c":             "-.-.",
	"d":             "-..",
	"e":             ".",
	"f":             "..-.",
	"g":             "--.",
	"h":             "....",
	"i":             "..",
	"j":             ".---",
	"k":             "-.-",
	"l":             ".-..",
	"m":             "--",
	"n":             "-.",
	"o":             "---",
	"p":             ".--.",
	"q":             "--.-",
	"r":             ".-.",
	"s":             "...",
	"t":             "-",
	"u":             "..-",
	"v":             "...-",
	"w":             ".--",
	"x":             "-..-",
	"y":             "-.--",
	"z":             "--..",
	"0":             "-----",
	"1":             ".----",
	"2":             "..---",
	"3":             "...--",
	"4":             "....-",
	"5":             ".....",
	"6":             "-....",
	"7":             "--...",
	"8":             "---..",
	"9":             "----.",
	".":             ".-.-.-",
	",":             "--..--",
	"?":             "..--..",
	"'":             ".----.",
	"!":             "-.-.--",
	"/":             "-..-.",
	"&":             ".-...",
	":":             "---...",
	";":             "-.-.-.",
	"=":             "-...-",
	"+":             ".-.-.",
	"-":             "-....-",
	"_":             "..--.-",
	"\"":            ".-..-.",
	"$":             "...-..-",
	"@":             ".--.-.",
	" ***amend*** ": "........",
	" ":             "*",
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
	reg := regexp.MustCompile(`\s+`)
	sentence = reg.ReplaceAllString(sentence, " ")
	log.Printf(sentence)
	for _, part := range strings.Split(sentence, " ") {

		alphabet, found := searchTable(part)

		if !found {
			Err = errors.New(fmt.Sprintf("Not found such code: %s", part))
			return
		}

		response += alphabet
	}
	return
}

func DecodeMorse(messageContent string) (decodeResult string, Err error) {
	sentence := strings.TrimSpace(messageContent)
	decodeResult, Err = decode(sentence)
	if Err != nil {
		log.Println("Failed to decode:", Err)
		return
	}
	return
}

func morseCodeOperation(mode string, content string) (answerSentence string, Err error) {
	switch mode {
	case "decode":
		answerSentence, Err = DecodeMorse(content)
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
