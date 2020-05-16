package src

import "errors"

func FetchMessage (cmd string)(message string, Err error){
	switch cmd {
	case "help":
		return HelpMessage, nil
	case "ping":
		return "Pong!", nil
	default:
		return DefaultMessage, errors.New("FetchMessage: No such word")
	}
}
