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

const formats = "bmp,gif,jpg,jpeg,jfif,png,tif,tiff,webp"
const quality = 75

const usage = "Usage: %s path/to/file.png\n" +
	"Converted jpeg will be saved to path/to/file.jpg\n" +
	"Supported formats: %s\n"

func main() {
	if len(os.Args) < 2 {
		exit(1, usage, os.Args[0], formats)
	}

	fileName := os.Args[1]
	jpegName := strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ".jpg"

	img, err := loadOrig(fileName)
	if err != nil {
		exit(1, "%s: Cannot load file: %v\n", fileName, err)
	}

	if err := saveJPEG(jpegName, img); err != nil {
		exit(1, "%s: Cannot save file %s: %v\n", fileName, jpegName, err)
	}

	exit(0, "%s: File converted. Created %s", fileName, jpegName)
}

func exit(status int, format string, a ...interface{}) {
	fmt.Printf(format, a...)
	os.Exit(status)
}

func loadOrig(name string) (image.Image, error) {
	in, err := os.Open(name)
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

func saveJPEG(name string, img image.Image) error {
	out, err := os.Create(name)
	if err != nil {
		return err
	}
	defer out.Close()

	opts := jpeg.Options{
		Quality: quality,
	}
	return jpeg.Encode(out, img, &opts)
}
