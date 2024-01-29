package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func main() {
	// Open the image file
	file, err := os.Open("./images/original.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Create folder if it does not exist
	err = os.MkdirAll("images", os.ModePerm)
	if err != nil {
		fmt.Println("Error creating images folder:", err)
		return
	}

	// Display input image
	displayImage("Input Image", img)

	// Save input image to a new file
	saveImage("images/input_image.jpg", img)

	// Create negative image
	negativeImg := createNegativeImage(img)

	// Display negative image
	displayImage("Negative Image", negativeImg)

	// Save negative image to a new file
	saveImage("images/negative_image.jpg", negativeImg)
}

// Function to create negative image
func createNegativeImage(img image.Image) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	negativeImg := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, a := img.At(x, y).RGBA()

			// Calculate negative values
			negR := 255 - int(r>>8)
			negG := 255 - int(g>>8)
			negB := 255 - int(b>>8)

			// Set negative color to the new image
			negativeImg.Set(x, y, color.RGBA{uint8(negR), uint8(negG), uint8(negB), uint8(a >> 8)})
		}
	}

	return negativeImg
}

// Function to display image
func displayImage(title string, img image.Image) {
	fmt.Println("Displaying", title)
	// You need to implement your own method to display images as per your environment
}

// Function to save image to file
func saveImage(filename string, img image.Image) {
	outFile, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, img, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Saved image to", filename)
}
