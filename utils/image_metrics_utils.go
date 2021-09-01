package utils

import "image"

func GetSize(sourceImage image.Image) (width uint32, height uint32) {
	return uint32(sourceImage.Bounds().Size().X), uint32(sourceImage.Bounds().Size().Y)
}
