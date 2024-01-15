package main

import (
	"image"
	"image/png"
	"log"
	"os"

	"github.com/disintegration/imaging"
	compression "github.com/nurlantulemisov/imagecompression"
)

func main() {
	src, err := imaging.Open("./original/profile.jpg")
	if err != nil {
		log.Fatal("Failed to open the image: %v", err)
	}

	// src = imaging.CropAnchor(src, 300, 300, imaging.Center)

	err = imaging.Save(src, "./compressed/compressed-profile.jpg")
	if err != nil {
		log.Fatal("Failed to save the image : %v", err)
	}

	//create a blurred version of the image
	img1 := imaging.Blur(src, 25)
	err = imaging.Save(img1, "./compressed/blurred-profile.jpg")
	if err != nil {
		log.Fatal("Failed to save the image : %v", err)
	}

	//create a grayscale version of the image
	img2 := imaging.Grayscale(src)
	err = imaging.Save(img2, "./compressed/grayscale-profile.jpg")
	if err != nil {
		log.Fatal("Failed to save the image : %v", err)
	}

	compress()
}

func compress() {
	// Open the image file
	file, err := os.Open("./original/profile.jpg")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("Error decoding the image:", err)
	}

	// Compress the image with a compression level
	compressing, err := compression.New(70)
	if err != nil {
		log.Fatal("Error creating compression:", err)
	}

	compressingImage := compressing.Compress(img)

	// Create the compressed file
	compressedFile, err := os.Create("./compressed/c70.jpg")
	if err != nil {
		log.Fatal("Error creating the compressed file:", err)
	}
	defer compressedFile.Close()

	// Encode and save the compressed image
	err = png.Encode(compressedFile, compressingImage)
	if err != nil {
		log.Fatal("Error encoding the compressed image:", err)
	}
}
