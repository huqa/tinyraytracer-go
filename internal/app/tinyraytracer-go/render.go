package tinyraytracer

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/huqa/tinyraytracer-go/internal/pkg/object"
	"github.com/huqa/tinyraytracer-go/internal/pkg/vector"
)

// Width defines the width of the resulting image in pixels
const Width = 1024

// Height defines the height of the resulting image in pixels
const Height = 768

// Fov defines the field of view angle for the camera
const Fov = math.Pi / 2.0

// framebuffer is a one dimensional array of Vectors
var framebuffer []vector.Vector

// Render renders an image and saves it to disk
func Render(sphere *object.Sphere) {
	// init framebuffer
	framebuffer = make([]vector.Vector, Width*Height)

	fmt.Println("Filling framebuffer")
	origin := vector.NewVector(0, 0, 0)
	// fill framebuffer
	for j := 0; j < Height; j++ {
		for i := 0; i < Width; i++ {
			x := (2*(float64(i)+0.5)/float64(Width) - 1) * math.Tan(Fov/2.0) * Width / float64(Height)
			y := -(2*(float64(j)+0.5)/float64(Height) - 1) * math.Tan(Fov/2.0)
			direction := vector.NewVector(x, y, -1).Normalize()
			framebuffer[i+j*Width] = CastRay(&origin, &direction, sphere)
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

// CastRay casts a ray and checks if the ray intersects with our sphere
func CastRay(origin *vector.Vector, direction *vector.Vector, sphere *object.Sphere) vector.Vector {
	sphereDistance := math.MaxFloat64
	if !sphere.RayIntersects(*origin, *direction, sphereDistance) {
		return vector.NewVector(0.2, 0.7, 0.8) // background color
	}
	return vector.NewVector(0.4, 0.4, 0.3)
}
