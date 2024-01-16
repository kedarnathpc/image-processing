# Image Processing Assignment

This repository contains a simple Go program for image processing. The program performs various operations on an input image and generates separate images for each operation.

## Getting Started

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/kedarnathpc/image-processing
    
    cd image-processing
    ```

2. Ensure you have Go installed on your machine.

3. Install the required third-party package:

    ```bash
    go get -u github.com/nfnt/resize
    ```

4. Place your input image in the repository and update the filename in the `main.go` file.

5. Run the program:

    ```bash
    go run main.go
    ```

## Operations

### 1. Resize Image

The program resizes the input image to 512x512 pixels using the `github.com/nfnt/resize` package.

### 2. Convert to Grayscale

The resized image is converted to a grayscale image.

### 3. Separate Color Components

The program generates three images, each showing only one of the RGB color components. The images display the red, blue, and green components of the original image.

## Output

The processed images are saved in the repository with the following filenames:

- `resized.jpg`
- `grayscale.jpg`
- `red_component.jpg`
- `blue_component.jpg`
- `green_component.jpg`


Feel free to use and modify the code for your own projects!

