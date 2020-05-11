package src

import (
"fmt"
"image"
color "image/color"
"image/jpeg"
"log"
"os"
"strconv"
)

var (
	x      = 0
	y      = 0
	width  = 400
	height = 300
	red    uint8
	blue   uint8
	green  uint8
)

func intToHexRGB(code uint32) (uint8, uint8, uint8) {
	var rgb [3]uint8
	for i := 0;i < 3;i++ {
		rgb[i] = uint8(code >> (3 - i) * 8 & 0xFF)
	}

	return rgb[0], rgb[1], rgb[2]
}

func generateImage(colorCode string) {
	if len(colorCode) != 6 {
		os.Exit(0)
	}
	num, _ := strconv.ParseInt(colorCode,16 ,32)

	img := image.NewRGBA(image.Rect(x, y, width, height))

	red, green, blue := intToHexRGB(uint32(num))
	for i := img.Rect.Min.Y; i < img.Rect.Max.Y; i++ {
		for j := img.Rect.Min.X; j < img.Rect.Max.X; j++ {
			img.Set(j, i, color.RGBA{R: red, G: green, B: blue})
		}
	}

	file, _ := os.Create("sample.jpeg")
	defer file.Close()

	err := jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
	if err != nil {
		log.Println(err)
	}

	log.Println("end")
}

func GenImage(code string) {
	fmt.Println(code[1:])
	generateImage(code[len("#"):])
}

