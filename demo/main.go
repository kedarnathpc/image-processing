package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func main() {
	// Open the JPEG image file
	imgFile, err := os.Open("./images/original.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer imgFile.Close()

	// Decode the JPEG image
	img, err := jpeg.Decode(imgFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Split image into RGB channels
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	red := image.NewRGBA(image.Rect(0, 0, width, height))
	green := image.NewRGBA(image.Rect(0, 0, width, height))
	blue := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			red.Set(x, y, color.RGBA{uint8(r >> 8), 0, 0, uint8(a >> 8)})
			green.Set(x, y, color.RGBA{0, uint8(g >> 8), 0, uint8(a >> 8)})
			blue.Set(x, y, color.RGBA{0, 0, uint8(b >> 8), uint8(a >> 8)})
		}
	}

	// Save each channel as an image
	saveImage("./images/red.jpg", red)
	saveImage("./images/green.jpg", green)
	saveImage("./images/blue.jpg", blue)
}

func saveImage(filename string, img image.Image) {
	// Create a new image file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Encode the image as JPEG and save to file
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
