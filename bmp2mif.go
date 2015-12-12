package main

import (
	"fmt"
	"golang.org/x/image/bmp"
	"image/color"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./bmp2mif bmp-file")
		os.Exit(-1)
	}

	// open the bmp input file
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// decode the bmp file to raw pixel
	img, err := bmp.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// print the mif file to stdout
	for y := 0; y < img.Bounds().Max.Y; y++ {
		for x := 0; x < img.Bounds().Max.X; x++ {
			color := color.NRGBAModel.Convert(img.At(x,y)).(color.NRGBA)
			fmt.Printf("%08b%08b%08b\n", color.R, color.G, color.B)
		}
	}

	// close the bmp file
	file.Close()
}
