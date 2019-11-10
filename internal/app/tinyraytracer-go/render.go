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
func Render(spheres []object.Sphere) {
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
			framebuffer[i+j*Width] = CastRay(&origin, &direction, spheres)
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

// CastRay casts a ray and checks if the ray intersects with objects in our scene
func CastRay(origin *vector.Vector, direction *vector.Vector, spheres []object.Sphere) vector.Vector {
	point := &vector.Vector{}
	N := &vector.Vector{}
	mat := &object.Material{}

	if !SceneIntersect(origin, direction, spheres, point, N, mat) {
		return vector.NewVector(0.2, 0.7, 0.8) // background color
	}
	return mat.DiffuseColor
}

// SceneIntersect checks if a ray intersects with objects in the scene and
// determines what material that ray casts on to
func SceneIntersect(
	origin *vector.Vector,
	direction *vector.Vector,
	spheres []object.Sphere,
	hit *vector.Vector,
	N *vector.Vector,
	mat *object.Material) bool {
	spheresDistance := math.MaxFloat64
	for _, sphere := range spheres {
		var distI float64
		t0, intersects := sphere.RayIntersects(*origin, *direction, distI)
		if intersects && t0 < spheresDistance {
			spheresDistance = t0
			k := origin.Add(direction.ScalarMultiply(distI))
			hit = &k
			n := hit.Subtract(sphere.Center).Normalize()
			N = &n
			mat.DiffuseColor = sphere.Material.DiffuseColor
		}
	}
	return spheresDistance < 1000.0
}
