// Defines the main executable for tinyraytracer-go
//
// This is a golang implementation of tinyraytracer made in c++ (or at least it tries to be)
// https://github.com/ssloy/tinyraytracer/wiki/Part-1:-understandable-raytracing
//
// Lots of parts have been changed from the original implementation
// Author of this pos: huqa (ville.m.riikonen@gmail.com)
package main

import (
	"fmt"

	"github.com/huqa/tinyraytracer-go/internal/pkg/object"
	"github.com/huqa/tinyraytracer-go/internal/pkg/vector"

	"github.com/huqa/tinyraytracer-go/internal/app/tinyraytracer-go"
)

func main() {
	fmt.Println("tinyraytracer-go v0.08")

	ivory := object.NewMaterial(vector.NewVector4(0.6, 0.3, 0.1, 0.0), vector.NewVector3(0.4, 0.4, 0.3), 50.0, 1.0)
	glass := object.NewMaterial(vector.NewVector4(0.0, 0.5, 0.1, 0.8), vector.NewVector3(0.6, 0.7, 0.8), 125.0, 1.5)
	redRubber := object.NewMaterial(vector.NewVector4(0.9, 0.1, 0.0, 0.0), vector.NewVector3(0.3, 0.1, 0.1), 10.0, 1.0)
	mirror := object.NewMaterial(vector.NewVector4(0.0, 10.0, 0.8, 0.0), vector.NewVector3(1.0, 1.0, 1.0), 1425.0, 1.0)

	/*sphere := object.Sphere{
		Center: vector.NewVector(-3, 0, -16),
		Radius: 2,
	}*/

	spheres := make([]object.Sphere, 0)
	spheres = append(spheres, object.NewSphere(vector.NewVector3(-3, 0, -16), 2, ivory))
	spheres = append(spheres, object.NewSphere(vector.NewVector3(-1.0, -1.5, -12), 2, glass))
	spheres = append(spheres, object.NewSphere(vector.NewVector3(3, -0.5, -18), 3, redRubber))
	spheres = append(spheres, object.NewSphere(vector.NewVector3(7, 5, -18), 4, mirror))

	lights := make([]object.Light, 0)
	lights = append(lights, object.NewLight(vector.NewVector3(-20, 20, 20), 1.5))
	lights = append(lights, object.NewLight(vector.NewVector3(20, -20, -20), 1.8))
	lights = append(lights, object.NewLight(vector.NewVector3(30, 20, 30), 1.7))

	fmt.Println("spheres ", spheres)
	fmt.Println("lights ", lights)
	tinyraytracer.Render(spheres, lights)
}
