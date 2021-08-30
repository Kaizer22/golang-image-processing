package main

import (
	"fmt"
	"main/transformations"
	"main/utils"
	"time"
)

func main() {
	imageLoader := utils.ImageLoader{}
	imageSaver := utils.ImageSaver{OutputDir: "test_results/"}
	testImage1, err := imageLoader.Open("test_images/Lenna_2048.png")

	fmt.Println("start interpolation" + time.Now().String())
	interpolationResult, err := transformations.NearestNeighbor(testImage1, 0.1, 0.01)
	if err != nil {

	}
	fmt.Println("finish interpolation" + time.Now().String())
	///iTest, err := basic_operators.SubtractImages(testImage1, interpolationResult)
	err = imageSaver.SavePNG(interpolationResult, "interpolation_result")
	if err != nil {

	}
}
