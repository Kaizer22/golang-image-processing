package transformations

import (
	"image"
	"main/utils"
)

type Interpolator interface {
}

func ProportionalNearestNeighbor(sourceImage image.Image, scale float64) (interpolatedImage *image.RGBA, err error) {
	return NearestNeighborScaling(sourceImage, scale, scale)
}

func NearestNeighborScaling(sourceImage image.Image, scaleX float64, scaleY float64) (interpolatedImage *image.RGBA, err error) {
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

func NearestNeighbourToResolution(sourceImage image.Image, targetWidth int32, targetHeight int32) (interpolatedImage *image.RGBA, err error) {
	width, height := utils.GetSize(sourceImage)
	scaleX := float64(targetWidth) / float64(width)
	scaleY := float64(targetHeight) / float64(height)
	return NearestNeighborScaling(sourceImage, scaleX, scaleY)
}

func BilinearInterpolation() (interpolatedImage *image.RGBA, err error) {
	//TODO
	return
}

func BicubicInterpolation() (interpolatedImage *image.RGBA, err error) {
	//TODO
	return
}
