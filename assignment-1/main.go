package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	// Open the original image file
	file, err := os.Open("./images/original.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Decode the original image
	originalImg, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Open the red, green, and blue channel images
	redFile, err := os.Open("./images/red_channel.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer redFile.Close()

	greenFile, err := os.Open("./images/green_channel.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer greenFile.Close()

	blueFile, err := os.Open("./images/blue_channel.jpg")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer blueFile.Close()

	// Decode the red, green, and blue channel images
	redChannel, _, err := image.Decode(redFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	greenChannel, _, err := image.Decode(greenFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	blueChannel, _, err := image.Decode(blueFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Combine images side by side
	combinedWidth := originalImg.Bounds().Dx() * 4
	combinedHeight := originalImg.Bounds().Dy()

	combinedImg := image.NewRGBA(image.Rect(0, 0, combinedWidth, combinedHeight))

	drawImage(combinedImg, originalImg, 0, 0)
	drawImage(combinedImg, redChannel, originalImg.Bounds().Dx(), 0)
	drawImage(combinedImg, greenChannel, originalImg.Bounds().Dx()*2, 0)
	drawImage(combinedImg, blueChannel, originalImg.Bounds().Dx()*3, 0)

	// Plot the combined image
	plotImage(combinedImg, "images/combined_image_plot.jpg")

	// Rotate the original image 180 degrees to the right
	rotatedImg := rotate180(originalImg)

	// Plot the rotated image
	plotImage(rotatedImg, "images/rotated_plot.jpg")

	// Enhance color of the original image
	enhancedImg := enhanceColor(originalImg)

	// Plot the enhanced color image
	plotImage(enhancedImg, "images/enhanced_color_image_plot.jpg")
}

// Function to draw an image onto another image at specified position
func drawImage(dst draw.Image, src image.Image, x, y int) {
	bounds := src.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			dst.Set(x+i, y+j, src.At(i, j))
		}
	}
}

// Function to plot image on a graph of pixels with visible axes
func plotImage(img image.Image, filename string) {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	// Create a new RGBA image with space for axes
	plot := image.NewRGBA(image.Rect(0, 0, width+20, height+20))

	// Fill the plot with white color
	draw.Draw(plot, plot.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Draw the image onto the plot
	draw.Draw(plot, image.Rect(10, 10, width+10, height+10), img, image.Point{}, draw.Src)

	// Draw the x-axis and y-axis
	for x := 0; x <= width; x++ {
		plot.Set(x+10, height+10, color.Black) // x-axis
	}
	for y := 0; y <= height; y++ {
		plot.Set(10, y+10, color.Black) // y-axis
	}

	// Save the plotted image
	saveImage(filename, plot)
}

// Function to rotate image 180 degrees
func rotate180(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	rotatedImg := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			rotatedX := width - x - 1
			rotatedY := height - y - 1
			rotatedImg.Set(rotatedX, rotatedY, img.At(x, y))
		}
	}

	return rotatedImg
}

// Function to enhance color using arithmetic operations
func enhanceColor(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	enhancedImg := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, a := img.At(x, y).RGBA()

			r = r * 2

			enhancedImg.Set(x, y, color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)})
		}
	}

	return enhancedImg
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
