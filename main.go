package main

import (
	"fmt"
	"main/basic_operators"
	"main/transformations"
	"main/utils"
	"time"
)

func main() {
	imageLoader := utils.ImageLoader{}
	imageSaver := utils.ImageSaver{OutputDir: "test_results/"}
	testImage1, err := imageLoader.Open("test_images/Lenna_2048.png")

	fmt.Println("start test" + time.Now().String())
	interpolationResult, err := transformations.NearestNeighborScaling(testImage1, 0.03, 0.03)
	if err != nil {
	}
	upscaleResult, err := transformations.NearestNeighbourToResolution(interpolationResult, 2048, 2048)
	iTest, err := basic_operators.SubtractImages(testImage1, upscaleResult)
	fmt.Println("finish test" + time.Now().String())

	err = imageSaver.SavePNG(iTest, "subtract_result")
	//clr := color.RGBA{A: 255, R: 100, G: 150, B: 212}
	img, err := transformations.RandomAddition(testImage1, 200, 200)
	if err != nil {
	}
	err = imageSaver.SavePNG(img, "addition_result")
}
