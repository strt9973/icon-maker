package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getImagePathList(dir string) []string {
    names := []string{}
    file, err := os.Open(dir)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    files, err := file.Readdir(-1)
    if err != nil {
        log.Fatal(err)
    }
    for _, f := range files {
        if !f.IsDir() {
            ext := strings.ToLower(filepath.Ext(f.Name()))
            if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
                names = append(names, f.Name())
            }
        }
    }

    return names
}

func main() {
    imagePathList := getImagePathList("source")

    for _, path := range imagePathList {
        img := loadImage("source/" + path)
        ext := filepath.Ext(path)
        saveIcons("output/" + strings.TrimSuffix(path, ext), img)
    }
    fmt.Println("Done")
}
