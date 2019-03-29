package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"math/rand"
	"time"

	distorts "github.com/john-pettigrew/image_distort/distorts"
)

func main() {
	var inputPath string
	var outputPath string

	flag.StringVar(&inputPath, "in", "", "Input file path (required)")
	flag.StringVar(&outputPath, "out", "", "Output file path (required)")

	flag.Parse()

	if inputPath == "" {
		errorAndExit("inputPath is required")
	}
	if outputPath == "" {
		errorAndExit("Output file is required\n")
	}

	inputImg, err := readImage(inputPath)
	if err != nil {
		fmt.Println(err)
		errorAndExit("Error reading image")
	}
    newImg := inputImg
	newImg = distorts.ChromaticAberation(newImg)


  rand.Seed(time.Now().UTC().UnixNano())
  shiftNum := rand.Intn(4)
    for i := 0; i < shiftNum; i++{
      newImg = distorts.PixelShift(newImg)
    }

	//Create output file
	output, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(err)
	}

	//Save output file
	options := jpeg.Options{Quality: 100}
	err = jpeg.Encode(output, newImg, &options)
	if err != nil {
		log.Fatal(err)
	}
}

func readImage(file string) (image.Image, error) {
	imageFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	image, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, err
	}

	return image, nil
}

func errorAndExit(errStr string) {
	fmt.Println(errStr)
	flag.Usage()
	os.Exit(1)
}
