package src

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"strconv"
	"strings"
)

func ParseColorCode(colorCode string) (Result color.RGBA, Err error) {
	red, Err := strconv.ParseInt(colorCode[0:2], 16, 32)
	if Err != nil {
		log.Println(fmt.Sprintf("Red --- strconv.ParseInt: %s", Err))
		return
	}
	green, Err := strconv.ParseInt(colorCode[2:4], 16, 32)
	if Err != nil {
		log.Println(fmt.Sprintf("Green --- strconv.ParseInt: %s", Err))
		return
	}
	blue, Err := strconv.ParseInt(colorCode[4:6], 16, 32)
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

func genImage(colorInfo color.RGBA) *image.RGBA {
	const (
		statrX = 0
		startY = 0
		width  = 40
		height = 30
	)

	img := image.NewRGBA(image.Rect(statrX, startY, width, height))

	for x := img.Rect.Min.Y; x < img.Rect.Max.Y; x++ {
		for y := img.Rect.Min.X; y < img.Rect.Max.X; y++ {
			img.Set(x, y, colorInfo)
		}
	}
	return img
}

func GenerateImage(session *discordgo.Session, message *discordgo.MessageCreate) {

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

	colorInfo, Err := ParseColorCode(colorCode)
	if Err != nil {
		log.Println(Err)
		return
	}

	colorImage := genImage(colorInfo)
	var (
		buffer     bytes.Buffer
		fileWriter = bufio.NewWriter(&buffer)
		fileReader = bufio.NewReader(&buffer)
	)

	Err = jpeg.Encode(fileWriter, colorImage, &jpeg.Options{Quality: 60})
	if Err != nil {
		text := fmt.Sprintf("Error at encoding jpeg: %s", Err)
		log.Println(text)
		return
	}
	Err = fileWriter.Flush()
	if Err != nil {
		text := fmt.Sprintf("Error at io.Writer flush: %s", Err)
		log.Println(text)
		return
	}

	log.Println("generatedImage: process ended")

	_, Err = session.ChannelFileSend(message.ChannelID, "sample.jpeg", fileReader)
	if Err != nil {
		log.Println(Err)
		return
	}
}
