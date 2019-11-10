package object

import (
	"github.com/huqa/tinyraytracer-go/internal/pkg/vector"
)

// Material defines a type of material to be used in an object
type Material struct {
	DiffuseColor vector.Vector
}

// NewMaterial constructs a new material
func NewMaterial(diffuseColor vector.Vector) Material {
	return Material{
		DiffuseColor: diffuseColor,
	}
}

// NewEmptyMaterial constructs a new material with a default black diffuse color
func NewEmptyMaterial() Material {
	return Material{
		DiffuseColor: vector.NewVector(0, 0, 0),
	}
}
