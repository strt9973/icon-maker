package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
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


func saveSingleImage(filename string, img image.Image, size int, margin bool) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    var processedImg image.Image
    if margin {
        processedImg = editImage(img, size, true)
    } else {
        processedImg = editImage(img, size, false)
    }
    if err := png.Encode(file, processedImg); err != nil {
        return err
    }
    return nil
}


func saveImages(path string, img image.Image) {
	_ = os.MkdirAll(filepath.Dir(path), os.ModePerm)
	for _, size := range IMAGE_SIZES {
		normalFilename := fmt.Sprintf("%s_%d.png", path, size)
        if err := saveSingleImage(normalFilename, img, size, false); err != nil {
            log.Fatal(err)
        }
        marginFilename := fmt.Sprintf("%s_%d_with_margin.png", path, size)
        if err := saveSingleImage(marginFilename, img, size, true); err != nil {
            log.Fatal(err)
        }
	}
    fmt.Println("Save image to", path)

}