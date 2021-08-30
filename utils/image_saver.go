package utils

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

type ImageSaver struct {
	OutputDir string
}

func (is *ImageSaver) SavePNG(imageToSave image.Image, filename string) (err error) {
	f, err := os.Create(is.OutputDir + filename + ".png")
	if err != nil {
		return err
	}
	defer f.Close()

	if err := png.Encode(f, imageToSave); err != nil {
		return err
	}
	return nil
}

func (is *ImageSaver) SaveJPEG(imageToSave image.Image, filename string, quality int) (err error) {
	f, err := os.Create(is.OutputDir + filename + ".jpeg")
	if err != nil {
		return err
	}
	defer f.Close()

	if err := jpeg.Encode(f, imageToSave, &jpeg.Options{Quality: quality}); err != nil {
		return err
	}
	return nil
}
