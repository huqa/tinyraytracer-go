package object

import (
	"github.com/huqa/tinyraytracer-go/internal/pkg/vector"
)

// Material defines a type of material to be used in an object
type Material struct {
	Albedo           vector.Vector3
	DiffuseColor     vector.Vector3
	SpecularExponent float64
}

// NewMaterial constructs a new material
func NewMaterial(albedo vector.Vector3, diffuseColor vector.Vector3, specularExponent float64) Material {
	return Material{
		Albedo:           albedo,
		DiffuseColor:     diffuseColor,
		SpecularExponent: specularExponent,
	}
}
