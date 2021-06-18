package main

import (
	"fmt"
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

const Quality = 85

const Usage = "Usage: %s file1.png file2.png...\n" +
	"Converted JPEGs will be saved to file1.jpg file2.jpg...\n" +
	"Supported formats: %s\n"

const Formats = "bmp,gif,jpg,jpeg,jfif,png,tif,tiff,webp"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, Usage, os.Args[0], Formats)
		os.Exit(1)
	}

	for _, arg := range os.Args[1:] {
		err := convert(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: Error: %v\n", arg, err)
			continue
		}
		fmt.Printf("%s: Done\n", arg)
	}
}

func convert(path string) error {
	img, err := loadOrig(path)
	if err != nil {
		return err
	}

	conv := strings.TrimSuffix(path, filepath.Ext(path)) + ".jpg"

	return saveJPEG(conv, img)
}

func loadOrig(path string) (image.Image, error) {
	in, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer in.Close()

	img, _, err := image.Decode(in)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func saveJPEG(path string, img image.Image) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	opts := jpeg.Options{
		Quality: Quality,
	}
	return jpeg.Encode(out, img, &opts)
}
