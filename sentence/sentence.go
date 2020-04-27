package sentence

const (
	BadRequest = "付けられるステータスは32文字以下である必要があります。"
	Forbidden = "Botに権限がないので変更できません。落ちぶれましょう。"
	Hello = "Hello, world!"
	Notify = "名前を元に戻しました"
	Usage = `
試作段階のBotです。現在実装されているコマンドは、
%hello: 挨拶をします。
%name: 発言者の名前を表示します。
%id: 発言者のIDを表示します。
%status [Emoji]: 表示名を絵文字にします。
%reset: 名前をユーザーネームに戻します。
	`
	Wrong = "コードが間違っている可能性があります。開発者に問い合わせてください。"
)
