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
	testImage2, err := imageLoader.Open("test_images/Lenna_2048.png")

	fmt.Println("start interpolation" + time.Now().String())
	interpolationResult, err := transformations.NearestNeighbor(testImage2, 1)
	if err != nil {

	}
	fmt.Println("finish interpolation" + time.Now().String())
	iTest, err := basic_operators.SubtractImages(testImage1, interpolationResult)
	err = imageSaver.SavePNG(iTest, "addition_result")
	if err != nil {

	}
}
