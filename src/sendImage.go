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

func GenerateImage(session *discordgo.Session, message *discordgo.MessageCreate) {
	const (
		x      = 0
		y      = 0
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
		log.Println("generateImage: len(colorCode) != 6")
		_, Err := session.ChannelMessageSend(message.ChannelID, "不正な値です。形式は16進のカラーコードである必要があります。")
		if Err != nil {
			log.Println(Err)
		}
		return
	}

	img := image.NewRGBA(image.Rect(x, y, width, height))

	red, Err := strconv.ParseInt(colorCode[0:2],16,32)
	if Err != nil {
		log.Println(fmt.Sprintf("strconv.ParseInt: %s",Err))
	}
	green, Err := strconv.ParseInt(colorCode[2:4],16,32)
	if Err != nil {
		log.Println(fmt.Sprintf("strconv.ParseInt: %s",Err))
	}
	blue, Err := strconv.ParseInt(colorCode[4:6],16,32)
	if Err != nil {
		log.Println(fmt.Sprintf("strconv.ParseInt: %s",Err))
	}

	log.Printf("\"%s\" parsed as %d, %d, %d,", colorCode, red, green, blue)
	for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
		for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
			img.Set(j, i, color.RGBA{R: byte(red), G: byte(green), B: byte(blue)})
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
