package main

import (
	"flag"
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
	flag.Usage = func() {
		help := `Usage: %s [options]

Description:
  正方形画像を角丸にして、アイコン用の複数サイズに変換します

Options:
`
		fmt.Fprintf(flag.CommandLine.Output(), help, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	var (
		srcP = flag.String("i", "source", "ソースディレクトリ。正方形の画像(png/jpg)を入れる")
		dstP = flag.String("o", "output", "出力ディレクトリ。加工後の画像が生成される")
	)
	flag.Parse()
	src, dst := *srcP, *dstP

	imagePathList := getImagePathList(src)

	for _, filename := range imagePathList {
		img := loadImage(filepath.Join(src, filename))
		stem := strings.TrimSuffix(filename, filepath.Ext(filename))
		saveImage(filepath.Join(dst, stem), img)
	}
	fmt.Println("Done")
}
