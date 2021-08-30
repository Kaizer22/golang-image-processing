package utils

import "image"

func GetSize(sourceImage image.Image) (width int64, height int64) {
	return int64(sourceImage.Bounds().Size().X), int64(sourceImage.Bounds().Size().Y)
}
