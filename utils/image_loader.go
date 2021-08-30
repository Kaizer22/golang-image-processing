package utils

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

type ImageLoader struct {
	File string
}

func (ip *ImageLoader) Open(filePath string) (image.Image, error) {
	_, filename := filepath.Split(filePath)
	filetype := strings.Split(filename, ".")[1]

	switch filetype {
	case "jpeg":
		return ip.OpenJPEG(filePath)
	case "png":
		return ip.OpenPNG(filePath)
	case "gif":
		return ip.OpenGIF(filePath)
	default:
		errText := fmt.Sprintf("cannot open file %s: unsupported file type", filePath)
		return nil, errors.New(errText)
	}
}

func (ip *ImageLoader) OpenPNG(path string) (image.Image, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (ip *ImageLoader) OpenJPEG(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := jpeg.Decode(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (ip *ImageLoader) OpenGIF(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := gif.Decode(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}
