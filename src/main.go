package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func getImagePathList(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	filenames := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		filename := entry.Name()
		ext := filepath.Ext(filename)
		if slices.Contains([]string{".png", ".jpg", ".jpeg"}, strings.ToLower(ext)) {
			filenames = append(filenames, filename)
		}
	}
	return filenames
}

func main() {
    imagePathList := getImagePathList("source")

    for _, path := range imagePathList {
        img := loadImage("source/" + path)
        ext := filepath.Ext(path)
        saveImage("output/" + strings.TrimSuffix(path, ext), img)
    }
    fmt.Println("Done")
}
