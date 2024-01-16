package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	// Open the image file
	file, err := os.Open("./original/input.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// 1. Resize the image to 512x512 pixels
	resizedImg := resize.Resize(512, 512, img, resize.Lanczos3)

	// Save the resized image
	saveImage(resizedImg, "./compressed/resized.jpg")

	// 2. Convert the image to grayscale
	grayscaleImg := convertToGrayscale(resizedImg)

	// Save the grayscale image
	saveImage(grayscaleImg, "./compressed/grayscale.jpg")

	// 3. Display separate images for red, blue, and green components
	redImg := extractColorComponent(resizedImg, color.RGBA{255, 0, 0, 255})
	blueImg := extractColorComponent(resizedImg, color.RGBA{0, 0, 255, 255})
	greenImg := extractColorComponent(resizedImg, color.RGBA{0, 255, 0, 255})

	// Save the red, blue, and green images
	saveImage(redImg, "./compressed/red_component.jpg")
	saveImage(blueImg, "./compressed/blue_component.jpg")
	saveImage(greenImg, "./compressed/green_component.jpg")

	fmt.Println("Image processing completed.")
}

func saveImage(img image.Image, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = jpeg.Encode(file, img, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func convertToGrayscale(img image.Image) image.Image {
	bounds := img.Bounds()
	gray := image.NewGray(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			grayColor := color.GrayModel.Convert(img.At(x, y))
			gray.Set(x, y, grayColor)
		}
	}

	return gray
}

func extractColorComponent(img image.Image, targetColor color.RGBA) image.Image {
	bounds := img.Bounds()
	resultImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA)
			resultColor := color.NRGBA{0, 0, 0, originalColor.A}

			switch targetColor {
			case color.RGBA{255, 0, 0, 255}: // Red component
				resultColor.R = originalColor.R
			case color.RGBA{0, 0, 255, 255}: // Blue component
				resultColor.B = originalColor.B
			case color.RGBA{0, 255, 0, 255}: // Green component
				resultColor.G = originalColor.G
			}

			resultImg.Set(x, y, resultColor)
		}
	}

	return resultImg
}
