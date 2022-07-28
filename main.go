package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func main() {
	imgfile, err := os.Open("./img.jpg")

	if err != nil {
		fmt.Println("img.jpg file not found!")
		os.Exit(1)
	}

	defer imgfile.Close()

	imgCfg, _, err := image.DecodeConfig(imgfile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	width := imgCfg.Width
	height := imgCfg.Height

	fmt.Println("Width : ", width)
	fmt.Println("Height : ", height)

	imgfile.Seek(0, 0)

	// get the image
	img, _, err := image.Decode(imgfile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Pixel map:")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rgbaPixel := rgbaToPixel(img.At(x, y).RGBA())

			fmt.Printf("[X : %d Y : %v] R : %v, G : %v, B : %v, A : %v  \n", x, y, rgbaPixel.R, rgbaPixel.G, rgbaPixel.B, rgbaPixel.A)
		}
	}
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

// Pixel struct example
type Pixel struct {
	R int
	G int
	B int
	A int
}
