package src

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"log"
	"strconv"
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

	for y := img.Rect.Min.Y; y < img.Rect.Max.Y; y++ {
		for x := img.Rect.Min.X; x < img.Rect.Max.X; x++ {
			img.Set(x, y, colorInfo)
		}
	}
	return img
}

func GenerateImage(colorCode string) (fileReader io.Reader, Err error) {
	var (
		buffer     bytes.Buffer
		fileWriter = bufio.NewWriter(&buffer)
	)

	if v, Err := strconv.ParseInt(colorCode[len("#"):], 16, 32); Err != nil || v < 0 {
		log.Println("strconv: invalid value; not Hex")
		return nil, Err
	}

	colorInfo, Err := ParseColorCode(colorCode)
	if Err != nil {
		log.Println(Err)
		return nil, Err
	}

	colorImage := genImage(colorInfo)

	fileReader = bufio.NewReader(&buffer)

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
	return
}
