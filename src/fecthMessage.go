package src

func FetchMessage (cmd string)(message string, Err error){
	switch cmd {
	case "help":
		return helpMessage, nil
	case "ping":
		return "Pong!", nil
	default:
		return DefaultMessage, nil
	}
	return
}
