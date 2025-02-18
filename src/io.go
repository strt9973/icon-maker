package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

var IMAGE_SIZES = []int{64, 180, 192, 512}

func loadImage(path string) image.Image {
    file, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    img, _, err := image.Decode(file)
    if err != nil {
        log.Fatal(err)
    }
    return img
}

func saveImage(path string, img image.Image) {
	for _, size := range IMAGE_SIZES {
		file, err := os.Create(fmt.Sprintf("%s_%d.png", path, size))
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		image := editImage(img, size)
		err = png.Encode(file, image)
		if err != nil {
			log.Fatal(err)
		}
	}
    fmt.Println("Save image to", path)

}