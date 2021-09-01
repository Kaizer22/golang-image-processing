package transformations

import (
	"image"
	"image/color"
	"main/utils"
	"math/rand"
	"time"
)

func ImageAddition(sourceImage image.Image, horizontalBordersSize uint32, verticalBordersSize uint32, color color.Color) (resultImage *image.RGBA, err error) {
	width, height := utils.GetSize(sourceImage)

	actualWidth := int(width + verticalBordersSize*2)
	actualHeight := int(height + horizontalBordersSize*2)

	secondVerticalBorderX := verticalBordersSize + width
	secondHorizontalBorderY := horizontalBordersSize + height

	resultImage = image.NewRGBA(image.Rect(0, 0, actualWidth, actualHeight))
	for i := 0; i < actualHeight; i++ {
		for j := 0; j < actualWidth; j++ {
			if uint32(i) < horizontalBordersSize || uint32(j) < verticalBordersSize {
				resultImage.Set(j, i, color)
			} else if uint32(i) > secondHorizontalBorderY || uint32(j) > secondVerticalBorderX {
				resultImage.Set(j, i, color)
			} else {
				sourceImagePixelX := int(uint32(j) - verticalBordersSize)
				sourceImagePixelY := int(uint32(i) - horizontalBordersSize)
				resultImage.Set(j, i,
					sourceImage.At(sourceImagePixelX, sourceImagePixelY))
			}
		}
	}
	return
}

func RandomAddition(sourceImage image.Image, horizontalBordersSize uint32, verticalBordersSize uint32) (resultImage *image.RGBA, err error) {
	width, height := utils.GetSize(sourceImage)
	rand.Seed(time.Now().UnixNano())

	actualWidth := int(width + verticalBordersSize*2)
	actualHeight := int(height + horizontalBordersSize*2)

	secondVerticalBorderX := verticalBordersSize + width
	secondHorizontalBorderY := horizontalBordersSize + height

	resultImage = image.NewRGBA(image.Rect(0, 0, actualWidth, actualHeight))
	for i := 0; i < actualHeight; i++ {
		for j := 0; j < actualWidth; j++ {
			if uint32(i) < horizontalBordersSize || uint32(j) < verticalBordersSize {
				resultImage.Set(j, i, sourceImage.At(rand.Intn(int(width)), rand.Intn(int(height))))
			} else if uint32(i) > secondHorizontalBorderY || uint32(j) > secondVerticalBorderX {
				resultImage.Set(j, i, sourceImage.At(rand.Intn(int(width)), rand.Intn(int(height))))
			} else {
				sourceImagePixelX := int(uint32(j) - verticalBordersSize)
				sourceImagePixelY := int(uint32(i) - horizontalBordersSize)
				resultImage.Set(j, i,
					sourceImage.At(sourceImagePixelX, sourceImagePixelY))
			}
		}
	}
	return
}

func BorderAddition(sourceImage *image.Image, horizontalBordersSize uint32, verticalBordersSize uint32) (resultImage *image.RGBA, err error) {
	//TODO
	return
}

func InterpolationAddition(sourceImage *image.Image, horizontalBordersSize uint32, verticalBordersSize uint32, backgroundBlur bool) (resultImage *image.RGBA, err error) {
	//TODO
	return
}
