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
const Fov = math.Pi / 3.0

// framebuffer is a one dimensional array of Vectors
var framebuffer []vector.Vector3

// Render renders an image and saves it to disk
func Render(spheres []object.Sphere, lights []object.Light) {
	// init framebuffer
	framebuffer = make([]vector.Vector3, Width*Height)

	fmt.Println("Filling framebuffer")
	origin := vector.NewVector3(0, 0, 0)
	// fill framebuffer
	for j := 0; j < Height; j++ {
		for i := 0; i < Width; i++ {
			//x := (2*(float64(i)+0.5)/float64(Width) - 1) * math.Tan(Fov/2.0) * Width / float64(Height)
			//y := -(2*(float64(j)+0.5)/float64(Height) - 1) * math.Tan(Fov/2.0)
			//direction := vector.NewVector3(x, y, -1).Normalize()
			x := (float64(i) + 0.5) - Width/2.0
			y := -(float64(j) + 0.5) + Height/2.0
			z := -float64(Height) / (2.0 * math.Tan(Fov/2.0))
			direction := vector.NewVector3(x, y, z).Normalize()
			framebuffer[i+j*Width] = CastRay(&origin, &direction, spheres, lights, 0)
		}
	}

	// save framebuffer to file
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))
	fmt.Println("Saving framebuffer to file")

	var max float64
	for j := 0; j < Height; j++ {
		for i := 0; i < Width; i++ {
			v := framebuffer[i+j*Width]
			max = math.Max(v.X, math.Max(v.Y, v.Z))
			if max > 1 {
				v = v.ScalarMultiply((1.0 / max))
			}
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
func CastRay(
	origin *vector.Vector3,
	direction *vector.Vector3,
	spheres []object.Sphere,
	lights []object.Light,
	depth uint,
) vector.Vector3 {

	point := &vector.Vector3{}
	N := &vector.Vector3{}
	mat := &object.Material{}

	if depth > 4 || !SceneIntersect(origin, direction, spheres, point, N, mat) {
		return vector.NewVector3(0.2, 0.7, 0.8) // background color
	}

	// offset
	p := 1E-3

	// reflections
	reflectionDirection := Reflect(*direction, *N)
	var reflectionOrigin vector.Vector3
	if reflectionDirection.DotProduct(*N) < 0 {
		reflectionOrigin = point.Subtract(N.ScalarMultiply(p))
	} else {
		reflectionOrigin = point.Add(N.ScalarMultiply(p))
	}
	reflectionColor := CastRay(&reflectionOrigin, &reflectionDirection, spheres, lights, depth+1)

	// refractions
	refractionDirection := Refract(*direction, *N, mat.RefractiveIndex, 1.0).Normalize()
	var refractionOrigin vector.Vector3
	if refractionDirection.DotProduct(*N) < 0 {
		refractionOrigin = point.Subtract(N.ScalarMultiply(p))
	} else {
		refractionOrigin = point.Add(N.ScalarMultiply(p))
	}
	refractionColor := CastRay(&refractionOrigin, &refractionDirection, spheres, lights, depth+1)

	// lights and shadows
	var diffuseLightIntensity float64
	var specularLightIntensity float64
	var lightDistance float64
	var shadowOrigin vector.Vector3
	for _, light := range lights {
		lr := light.Position.Subtract(*point)
		lightDirection := lr.Normalize()
		lightDistance = lr.Magnitude()
		if lightDirection.DotProduct(*N) < 0 {
			shadowOrigin = point.Subtract(N.ScalarMultiply(p))
		} else {
			shadowOrigin = point.Add(N.ScalarMultiply(p))
		}
		shadowPoint := vector.Vector3{}
		shadowNormal := vector.Vector3{}
		tempMaterial := object.Material{}

		if SceneIntersect(&shadowOrigin, &lightDirection, spheres, &shadowPoint, &shadowNormal, &tempMaterial) &&
			shadowPoint.Subtract(shadowOrigin).Magnitude() < lightDistance {
			continue
		}

		diffuseLightIntensity += light.Intensity * math.Max(0.0, lightDirection.DotProduct(*N))
		specularLightIntensity += math.Pow(
			math.Max(0.0, Reflect(lightDirection, *N).DotProduct(*direction)),
			mat.SpecularExponent,
		) * light.Intensity
	}

	l := mat.DiffuseColor.ScalarMultiply(diffuseLightIntensity).ScalarMultiply(mat.Albedo.X)
	sp := specularLightIntensity * mat.Albedo.Y
	l1 := vector.NewVector3(1.0, 1.0, 1.0).ScalarMultiply(sp)
	re := reflectionColor.ScalarMultiply(mat.Albedo.Z)
	rf := refractionColor.ScalarMultiply(mat.Albedo.W)
	return l.Add(l1).Add(re).Add(rf)
}

// SceneIntersect checks if a ray intersects with objects in the scene and
// determines what material that ray casts on to
func SceneIntersect(
	origin *vector.Vector3,
	direction *vector.Vector3,
	spheres []object.Sphere,
	hit *vector.Vector3,
	N *vector.Vector3,
	mat *object.Material,
) bool {
	spheresDistance := math.MaxFloat64
	for _, sphere := range spheres {
		var distI float64
		t0, intersects := sphere.RayIntersects(*origin, *direction, distI)
		if intersects && t0 < spheresDistance {
			spheresDistance = t0
			k := origin.Add(direction.ScalarMultiply(t0))
			hit.Copy(k)
			n := k.Subtract(sphere.Center).Normalize()
			N.Copy(n)
			mat.DiffuseColor = sphere.Material.DiffuseColor
			mat.Albedo = sphere.Material.Albedo
			mat.SpecularExponent = sphere.Material.SpecularExponent
			mat.RefractiveIndex = sphere.Material.RefractiveIndex
		}
	}
	return spheresDistance < 1000.0
}

// Reflect computes the illumination for a point
// using the Phong reflection model
func Reflect(I vector.Vector3, N vector.Vector3) vector.Vector3 {
	return I.Subtract(N.ScalarMultiply(((I.DotProduct(N)) * 2.0)))
}

// Refract using Snell's law
func Refract(I vector.Vector3, N vector.Vector3, etaT float64, etaI float64) vector.Vector3 {
	cosi := -math.Max(-1.0, math.Min(1.0, I.DotProduct(N)))
	if cosi < 0 {
		// if the ray comes from the inside the object, swap the air and the media
		return Refract(I, N.Negate(), etaI, etaT)
	}
	eta := etaI / etaT
	k := 1 - eta*eta*(1-cosi*cosi)
	if k < 0 {
		return vector.NewVector3(1, 0, 0)
	}
	return I.ScalarMultiply(eta).Add(N.ScalarMultiply((eta*cosi - math.Sqrt(k))))
}
