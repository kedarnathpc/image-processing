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

	// Split RGB channels
	redChannel, greenChannel, blueChannel := splitRGBChannels(img)

	// Display RGB split channel images
	displayImage("Red Channel", redChannel)
	displayImage("Green Channel", greenChannel)
	displayImage("Blue Channel", blueChannel)

	// Save RGB split channel images
	saveImage("images/red_channel.jpg", redChannel)
	saveImage("images/green_channel.jpg", greenChannel)
	saveImage("images/blue_channel.jpg", blueChannel)

	// Enhanced color image using arithmetic operations
	enhancedImg := enhanceColor(img)

	// Display enhanced color image
	displayImage("Enhanced Color Image", enhancedImg)

	// Save enhanced color image to a new file
	saveImage("images/enhanced_color_image.jpg", enhancedImg)
}

// Function to split RGB channels
func splitRGBChannels(img image.Image) (red, green, blue *image.RGBA) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	red = image.NewRGBA(image.Rect(0, 0, width, height))
	green = image.NewRGBA(image.Rect(0, 0, width, height))
	blue = image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			red.Set(x, y, color.RGBA{uint8(r >> 8), 0, 0, 255})
			green.Set(x, y, color.RGBA{0, uint8(g >> 8), 0, 255})
			blue.Set(x, y, color.RGBA{0, 0, uint8(b >> 8), 255})
		}
	}

	return red, green, blue
}

// Function to enhance color using arithmetic operations
func enhanceColor(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	enhancedImg := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, a := img.At(x, y).RGBA()

			// Enhance the color (example: increasing red component)
			r = r * 2 // You can adjust this operation based on your enhancement requirement
			// Similar operations for green and blue channels can be applied here

			enhancedImg.Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
		}
	}

	return enhancedImg
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
