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
	fmt.Println("tinyraytracer-go v0.02")
	sphere := object.Sphere{
		Center: vector.NewVector(-3, 0, -16),
		Radius: 2,
	}
	tinyraytracer.Render(&sphere)
}
