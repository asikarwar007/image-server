package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// Create a new router
	http.HandleFunc("/", serveImage)

	// Start the server
	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveImage(w http.ResponseWriter, r *http.Request) {
	// Get the image name from the URL parameter
	imageName := r.URL.Path[len("/"):]

	// Open the original image file
	filePath := fmt.Sprintf("./%s", imageName)

	file, err := http.Dir("").Open(filePath)

	if err != nil {
		http.Error(w, "Image not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Decode the original image
	originalImage, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Error decoding image", http.StatusInternalServerError)
		return
	}

	// Parse query parameters for width, height, and quality
	widthStr := r.URL.Query().Get("w")
	heightStr := r.URL.Query().Get("h")
	qualityStr := r.URL.Query().Get("q")

	// Default values if query parameters are not provided or invalid
	width := 0
	height := 0
	quality := 80

	// Parse width parameter
	if widthStr != "" {
		width, err = strconv.Atoi(widthStr)
		if err != nil {
			http.Error(w, "Invalid width parameter", http.StatusBadRequest)
			return
		}
	}

	// Parse height parameter
	if heightStr != "" {
		height, err = strconv.Atoi(heightStr)
		if err != nil {
			http.Error(w, "Invalid height parameter", http.StatusBadRequest)
			return
		}
	}

	// Parse quality parameter
	if qualityStr != "" {
		quality, err = strconv.Atoi(qualityStr)
		if err != nil {
			http.Error(w, "Invalid quality parameter", http.StatusBadRequest)
			return
		}
	}

	// Resize the image if width or height is provided
	if width > 0 || height > 0 {
		resizedImage := resizeImage(originalImage, width, height)

		// Serve the resized image
		w.Header().Set("Content-Type", "image/jpeg")
		jpeg.Encode(w, resizedImage, &jpeg.Options{Quality: quality})
	} else {
		// Serve the original image
		w.Header().Set("Content-Type", "image/jpeg")
		jpeg.Encode(w, originalImage, &jpeg.Options{Quality: quality})
	}
}

func resizeImage(img image.Image, width, height int) image.Image {
	if width <= 0 && height <= 0 {
		return img
	}

	if width <= 0 {
		width = img.Bounds().Dx() * height / img.Bounds().Dy()
	}

	if height <= 0 {
		height = img.Bounds().Dy() * width / img.Bounds().Dx()
	}

	resizedImage := image.NewRGBA(image.Rect(0, 0, width, height))
	originalBounds := img.Bounds()
	resizedBounds := resizedImage.Bounds()
	for y := resizedBounds.Min.Y; y < resizedBounds.Max.Y; y++ {
		for x := resizedBounds.Min.X; x < resizedBounds.Max.X; x++ {
			srcX := x * originalBounds.Dx() / resizedBounds.Dx()
			srcY := y * originalBounds.Dy() / resizedBounds.Dy()
			resizedImage.Set(x, y, img.At(srcX, srcY))
		}
	}

	return resizedImage
}
