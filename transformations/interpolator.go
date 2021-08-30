package transformations

import (
	"errors"
	"image"
	"main/utils"
)

func ProportionalNearestNeighbor(sourceImage image.Image, scale float64) (interpolatedImage *image.RGBA, err error) {
	return NearestNeighbor(sourceImage, scale, scale)
}

func NearestNeighbor(sourceImage image.Image, scaleX float64, scaleY float64) (interpolatedImage *image.RGBA, err error) {
	if scaleX < 1 || scaleY < 1 {
		return nil, errors.New("method does not support downscaling")
	}
	width, height := utils.GetSize(sourceImage)
	targetWidth := int(float64(width) * scaleX)
	targetHeight := int(float64(height) * scaleY)

	interpolatedImage = image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))
	for i := 0; i < targetHeight; i++ {
		for j := 0; j < targetWidth; j++ {
			interpolatedImage.Set(j, i, sourceImage.At(int(float64(j)/scaleX), int(float64(i)/scaleY)))
		}
	}
	return
}

func BilinearInterpolation() (interpolatedImage *image.RGBA, err error) {
	//TODO
	return
}

func BicubicInterpolation() (interpolatedImage *image.RGBA, err error) {
	//TODO
	return
}
