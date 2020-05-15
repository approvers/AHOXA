package src

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"strconv"
	"strings"
)

func ParseColorCode(colorCode string) (Result color.RGBA, Err error) {
	red, Err := strconv.ParseInt(colorCode[0:2],16,32)
	if Err != nil {
		log.Println(fmt.Sprintf("Red --- strconv.ParseInt: %s", Err))
		return
	}
	green, Err := strconv.ParseInt(colorCode[2:4],16,32)
	if Err != nil {
		log.Println(fmt.Sprintf("Green --- strconv.ParseInt: %s", Err))
		return
	}
	blue, Err := strconv.ParseInt(colorCode[4:6],16,32)
	if Err != nil {
		log.Println(fmt.Sprintf("Blue --- strconv.ParseInt: %s", Err))
		return
	}
	log.Printf("\"%s\" parsed as %d, %d, %d,", colorCode, red, green, blue)
	Result.R = byte(red)
	Result.G = byte(green)
	Result.B = byte(blue)
	return

}

func GenerateImage(session *discordgo.Session, message *discordgo.MessageCreate) {
	const (
		statrX = 0
		startY = 0
		width  = 40
		height = 30
	)
	content := strings.TrimSpace(message.Content)
	if !strings.HasPrefix(content, "#") {
		return
	}

	if v, Err := strconv.ParseInt(content[1:], 16, 32); Err != nil || v < 0 {
		log.Println("strconv: invalid value; not Hex")
		return
	}

	colorCode := content[len("#"):]
	if len(colorCode) != 6 {
		log.Println("generateImage: len(colorCode) must be just 6")
		_, Err := session.ChannelMessageSend(message.ChannelID, "不正な値です。形式は16進のカラーコードである必要があります。")
		if Err != nil {
			log.Println(Err)
		}
		return
	}



	img := image.NewRGBA(image.Rect(statrX, startY, width, height))

	colorInfo, Err := ParseColorCode(colorCode)
	for x := img.Rect.Min.Y; x < img.Rect.Max.Y; x++ {
		for y := img.Rect.Min.X; y < img.Rect.Max.X; y++ {
			img.Set(x, y, colorInfo)
		}
	}

	File, _ := os.Create("sample.jpeg")
	defer File.Close()

	Err = jpeg.Encode(File, img, &jpeg.Options{Quality: 60})
	if Err != nil {
		text := fmt.Sprintf("Error at encoding jpeg: %s",Err)
		log.Println(text)
	}

	log.Println("generatedImage: process ended")

	file, Err := os.Open("sample.jpeg")
	if Err != nil {
		log.Println(Err)
	}

	_, Err = session.ChannelFileSend(message.ChannelID, "sample.jpeg", file)
	if Err != nil {
		log.Println(Err)
	}
}
