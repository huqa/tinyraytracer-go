package tinyraytracer

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/huqa/tinyraytracer-go/internal/pkg/vector"
)

// Width defines the width of the resulting image in pixels
const Width = 1024

// Height defines the height of the resulting image in pixels
const Height = 768

// framebuffer is a one dimensional array of Vectors
var framebuffer []vector.Vector

// Render renders an image and saves it to disk
func Render() {
	// init framebuffer
	framebuffer = make([]vector.Vector, Width*Height)

	fmt.Println("Filling framebuffer")
	// fill framebuffer
	for j := 0; j < Height; j++ {
		for i := 0; i < Width; i++ {
			framebuffer[i+j*Width] = vector.NewVector(
				float64(j)/float64(Height),
				float64(i)/float64(Width),
				0,
			)
		}
	}

	// save framebuffer to file
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))
	fmt.Println("Saving framebuffer to file")

	for j := 0; j < Height; j++ {
		for i := 0; i < Width; i++ {
			v := framebuffer[i+j*Width]
			r, g, b := v.ConvertToRGB()
			c := color.RGBA{r, g, b, 255}
			img.Set(i, j, c)
		}
	}

	output, err := os.Create("test.png")
	if err != nil {
		panic("Oh noes I can't create the image file")
	}
	defer output.Close()

	png.Encode(output, img)
}
