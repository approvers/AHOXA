package src

func fetchMessage (cmd string) string{
	switch cmd {
	case "status":
		return updateMessage
	case "help":
		return helpMessage
	case "ping":
		return "Pong!"
	default:
		return "該当するコマンドがありません`%help`を参照してください。"
	}
}
