package object

import (
	"math"

	"github.com/huqa/tinyraytracer-go/internal/pkg/vector"
)

// Sphere defines a struct for a sphere object
type Sphere struct {
	Center   vector.Vector3
	Radius   float64
	Material Material
}

// NewSphere constructs a new sphere
func NewSphere(center vector.Vector3, radius float64, material Material) Sphere {
	return Sphere{
		Center:   center,
		Radius:   radius,
		Material: material,
	}
}

// RayIntersects checks if a given ray (originating from origin in direction of direction vector)
// intersects with our sphere
func (s Sphere) RayIntersects(origin vector.Vector3, direction vector.Vector3, t0 float64) (float64, bool) {
	// solve for tc
	// find vector from ray origin to sphere center
	L := s.Center.Subtract(origin)
	tc := L.DotProduct(direction)

	if tc < 0.0 {
		return t0, false
	}
	d2 := L.DotProduct(L) - (tc * tc)
	r2 := s.Radius * s.Radius
	if d2 > r2 {
		return t0, false
	}

	// solve t1c
	t1c := math.Sqrt(r2 - d2)

	// solve intersection points
	t0 = tc - t1c
	t1 := tc + t1c
	if t0 < 0.0 {
		t0 = t1
	}
	if t0 < 0.0 {
		return t0, false
	}
	return t0, true
}
