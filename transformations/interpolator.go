package transformations

import (
	"errors"
	"image"
	"main/utils"
)

func NearestNeighbor(sourceImage image.Image, scale float64) (interpolatedImage *image.RGBA, err error) {
	if scale < 1 {
		return nil, errors.New("method does not support downscaling")
	}
	width, height := utils.GetSize(sourceImage)
	targetWidth := int(float64(width) * scale)
	targetHeight := int(float64(height) * scale)

	interpolatedImage = image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))
	for i := 0; i < targetHeight; i++ {
		for j := 0; j < targetWidth; j++ {
			interpolatedImage.Set(j, i, sourceImage.At(int(float64(j)/scale), int(float64(i)/scale)))
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
