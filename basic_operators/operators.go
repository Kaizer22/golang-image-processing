package basic_operators

import (
	"errors"
	"image"
	"image/color"
	"main/utils"
)

func ApplyPerPixelOperation(image1 image.Image,
	operation func(color1 color.Color) color.Color) (resultImage *image.RGBA, err error) {
	width1, height1 := utils.GetSize(image1)
	for i := 0; i < int(height1); i++ {
		for j := 0; j < int(width1); j++ {
			resultColor := operation(image1.At(j, i))
			resultImage.Set(j, i, resultColor)
		}
	}
	return
}

func ApplyCustomOperation(image1 image.Image, image2 image.Image,
	operation func(color1 color.Color, color2 color.Color) color.Color) (resultImage *image.RGBA, err error) {
	width1, height1 := utils.GetSize(image1)
	width2, height2 := utils.GetSize(image2)
	if width1 != width2 || height2 != height1 {
		return nil, errors.New("method does not support images with different sizes")
	}

	resultImage = image.NewRGBA(image1.Bounds())
	for i := 0; i < int(height1); i++ {
		for j := 0; j < int(width1); j++ {
			resultColor := operation(image1.At(j, i), image2.At(j, i))
			resultImage.Set(j, i, resultColor)
		}
	}
	return
}

func AddImages(image1 image.Image, image2 image.Image) (resultImage *image.RGBA, err error) {
	return ApplyCustomOperation(image1, image2, AddPixels)
}

func SubtractImages(image1 image.Image, image2 image.Image) (resultImage image.Image, err error) {
	return ApplyCustomOperation(image1, image2, SubtractPixels)
}

func MultiplyImages(image1 image.Image, image2 image.Image) (resultImage image.Image, err error) {
	return ApplyCustomOperation(image1, image2, MultiplyPixels)
}

func DivideImages(image1 image.Image, image2 image.Image) (resultImage image.Image, err error) {
	return ApplyCustomOperation(image1, image2, DividePixels)
}

func AddPixels(color1 color.Color, color2 color.Color) (resultColor color.Color) {
	r1, g1, b1, a1 := color1.RGBA()
	r2, g2, b2, a2 := color2.RGBA()

	var resCol color.RGBA
	if r1+r2 > 255 {
		resCol.R = 255
	} else {
		resCol.R = uint8(r1 + r2)
	}
	if g1+g2 > 255 {
		resCol.G = 255
	} else {
		resCol.G = uint8(g1 + g2)
	}
	if b1+b2 > 255 {
		resCol.B = 255
	} else {
		resCol.B = uint8(b1 + b2)
	}
	if a1+a2 > 255 {
		resCol.A = 255
	} else {
		resCol.A = uint8(a1 + a2)
	}
	return resCol
}
func SubtractPixels(color1 color.Color, color2 color.Color) (resultColor color.Color) {
	r1, g1, b1, a1 := color1.RGBA()
	r2, g2, b2, a2 := color2.RGBA()

	var resCol color.RGBA
	if r1 >= r2 {
		resCol.R = 0
	} else {
		resCol.R = uint8(r1 - r2)
	}
	if g1 >= g2 {
		resCol.G = 0
	} else {
		resCol.G = uint8(g1 - g2)
	}
	if b1 >= b2 {
		resCol.B = 0
	} else {
		resCol.B = uint8(b1 - b2)
	}
	//TODO Alpha channel unification
	if a1 >= a2 {
		resCol.A = 255
	} else {
		resCol.A = 255
	}
	return resCol
}
func MultiplyPixels(color1 color.Color, color2 color.Color) (resultColor color.Color) {
	r1, g1, b1, a1 := color1.RGBA()
	r2, g2, b2, a2 := color2.RGBA()

	var resCol color.RGBA
	if r1*r2 > 255 {
		resCol.R = 255
	} else {
		resCol.R = uint8(r1 * r2)
	}
	if g1*g2 > 255 {
		resCol.G = 255
	} else {
		resCol.G = uint8(g1 * g2)
	}
	if b1*b2 > 255 {
		resCol.B = 255
	} else {
		resCol.B = uint8(b1 * b2)
	}
	if a1*a2 > 255 {
		resCol.A = 255
	} else {
		resCol.A = uint8(a1 * a2)
	}
	return resCol
}
func DividePixels(color1 color.Color, color2 color.Color) (resultColor color.Color) {
	r1, g1, b1, a1 := color1.RGBA()
	r2, g2, b2, a2 := color2.RGBA()

	var resCol color.RGBA
	if r1/r2 > 255 {
		resCol.R = 255
	} else {
		resCol.R = uint8(r1 / r2)
	}
	if g1/g2 > 255 {
		resCol.G = 255
	} else {
		resCol.G = uint8(g1 / g2)
	}
	if b1/b2 > 255 {
		resCol.B = 255
	} else {
		resCol.B = uint8(b1 / b2)
	}
	if a1/a2 > 255 {
		resCol.A = 255
	} else {
		resCol.A = uint8(a1 / a2)
	}
	return resCol
}
