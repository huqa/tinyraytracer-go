package object

import (
	"github.com/huqa/tinyraytracer-go/internal/pkg/vector"
)

// Light defines a struct for a light source
type Light struct {
	Position  vector.Vector3
	Intensity float64
}

// NewLight constructs a new light source with given position and intensity
func NewLight(position vector.Vector3, intensity float64) Light {
	return Light{
		Position:  position,
		Intensity: intensity,
	}
}
