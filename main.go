package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func init() {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

func main() {
	img, err := openImage("img.jpg")
	if err != nil {
		fmt.Println("img.jpg file not found!")
		os.Exit(1)
	}
	fmt.Printf("imgfile: %T\n", img)
	size := img.Bounds().Size()
	fmt.Printf("size: %d\n", size)
	fmt.Printf("size type: %T\n", size)
	fmt.Printf("Width X: %d\n", size.X)
	fmt.Printf("Height Y: %d\n", size.Y)
	imgRect := img.Bounds()
	fmt.Printf("Rect type: %T\n", imgRect)
	fmt.Println(imgRect)
	fmt.Println(imgRect.Bounds().Min.X)
	tpixel := img.At(0, 0)
	fmt.Printf("tpixel type: %T\n", tpixel)
	fmt.Println(tpixel.RGBA())

	// loop over image, let's do something real...
	for y := imgRect.Min.Y; y < imgRect.Max.Y; y++ {
		for x := imgRect.Min.X; x < imgRect.Max.X; x++ {
			pixelValue := rgbaToPixel(img.At(x, y).RGBA())
			pixelEven := pixelIsEven(pixelValue.R, pixelValue.G, pixelValue.B)
			if pixelEven {
				fmt.Printf("[%d,%d:%t]:[White]\n", x, y, pixelEven)
			} else {
				fmt.Printf("[%d,%d:%t]:[Black]\n", x, y, pixelEven)
			}

		}
	}
}

func openImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()
	fi, _ := f.Stat()

	fmt.Println(fi.Name())

	img, format, err := image.Decode(f)
	if err != nil {
		fmt.Println("Decoding error:", err.Error())
		return nil, err
	}
	if format != "jpeg" {
		fmt.Println("image format is not jpeg")
		return nil, errors.New("")
	}
	return img, nil
}

func pixelIsEven(r int, g int, b int) bool {
	pixelSum := r + g + b
	//fmt.Println(pixelSum)
	//fmt.Println(pixelSum % 2)
	if pixelSum%2 == 0 {
		return true
	} else {
		return false
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
