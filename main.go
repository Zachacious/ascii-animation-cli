package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"time"

	"github.com/qeesung/image2ascii/convert"
)

type Options struct {
	Ratio           float64 // convert ratio
	FixedWidth      int     // convert the image width fixed width
	FixedHeight     int     // convert the image width fixed height
	FitScreen       bool    // only work on terminal, fit the terminal height or width
	StretchedScreen bool    // only work on terminal, stretch the width and heigh to overspread the terminal screen
	Colored         bool    // only work on terminal, output ascii with color
	Reversed        bool    // if reverse the ascii pixels
}

type Converter interface {
	// convert a image object to ascii matrix
	Image2ASCIIMatrix(image image.Image, imageConvertOptions *Options) []string
	// convert a image object to ascii matrix and then join the matrix to a string
	Image2ASCIIString(image image.Image, options *Options) string
	// convert a image object by input a string to ascii matrix
	ImageFile2ASCIIMatrix(imageFilename string, option *Options) []string
	// convert a image object by input a string to ascii matrix then join the matrix to a string
	ImageFile2ASCIIString(imageFilename string, option *Options) string
}

func playAnimation(converter *convert.ImageConverter, options convert.Options, numframes int, framedelay int, repeat int, path string, numlines int) {
	// var converter Converter = *conv

	for ij := 0; ij < repeat; ij++ {
		for i := 0; i < numframes; i++ {
			fmt.Print(converter.ImageFile2ASCIIString(fmt.Sprintf("%s/%d.gif", path, i), &options))
			// delay
			time.Sleep(time.Duration(framedelay) * time.Millisecond)
			//clear last lines
			fmt.Printf("\033[%dA", numlines)
		}
	}
	// skip 40 lines
	fmt.Printf("\033[%dB", numlines)
}

func main() {

	pathFlagPtr := flag.String("path", "frames", "path to frames")
	numframesFlagPtr := flag.Int("frames", 10, "number of frames")
	framedelayFlagPtr := flag.Int("delay", 30, "delay between frames")
	repeatFlagPtr := flag.Int("repeat", 1, "number of times to repeat animation")
	widthFlagPtr := flag.Int("width", 70, "width of output")
	heightFlagPtr := flag.Int("height", 40, "height of output")

	flag.Parse()

	// Create convert options
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = *widthFlagPtr
	convertOptions.FixedHeight = *heightFlagPtr
	// convertOptions.Ratio = 0.5
	// convertOptions.FitScreen = true

	// Create the image converter
	converter := convert.NewImageConverter()

	playAnimation(converter, convertOptions, *numframesFlagPtr, *framedelayFlagPtr, *repeatFlagPtr, *pathFlagPtr, *heightFlagPtr)

}
