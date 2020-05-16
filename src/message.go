package src

import "fmt"

var (
	HelpMessage = "```asciidoc\n" +
`= Title =
	something
= Description =
	表示名に絵文字を付けることでステータスを表示するBotの予定
= Command =
	(Prefix: "%")
	help :: このBotの概要を説明します。
	ping :: Pong!と返します。かわいいですね。
	(prefix: "#")
	[HexRGB] :: 16進カラーコードの画像を返します。略記法は現在対応していません。
= Source =
	github.com/brokenManager/change-status-go
` + "```\n"
	DefaultMessage = fmt.Sprintf("該当するコマンドがありません。`%shelp`を参照してください。",prefix)
)
