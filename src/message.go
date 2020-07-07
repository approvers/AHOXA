package src

var (
	helpMessage = "```asciidoc\n" +
		`= Title =
	something
= Description =
	表示名に絵文字を付けることでステータスを表示するBotの予定
= Command =
	(Prefix: "%")
	help :: このBotの概要を説明します。
	ping :: Pong!と返します。かわいいですね。
	color :: 第二引数に以下の形式で入力:
		(prefix: "#")
		[HexRGB] :: 16進カラーコードの画像を返します。略記法は現在対応していません。
	morse :: モールス暗号に関する変換をします。第2引数で動作を指定します。
		decode :: 暗号文　→　平文
		[Tips]
			現在任意の長さの空白を暗号変換に際して圧縮しています。これにより本来長い空白を入れることによって
			平文での空白を表現していたところを、*を代替しています。
= Source =
	github.com/brokenManager/change-status-go
` + "```\n"
)
