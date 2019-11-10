package vector

import "math"

// Vector defines a struct holding the X, Y, Z values of a vector
// as float64
type Vector struct {
	X float64
	Y float64
	Z float64
}

// NewVector constructs a new vector
func NewVector(x, y, z float64) Vector {
	return Vector{
		x,
		y,
		z,
	}
}

// DotProduct returns the dot product of vectors u and v
func (u Vector) DotProduct(v Vector) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

// Magnitude returns the length of vector u
func (u Vector) Magnitude() float64 {
	return math.Sqrt(u.DotProduct(u))
}

// CrossProduct returns the cross product of vectors u and v
func (u Vector) CrossProduct(v Vector) Vector {
	return Vector{
		u.Y*v.Z - u.Z*v.Y,
		u.Z*v.X - u.X*v.Z,
		u.X*v.Y - u.Y*v.X,
	}
}

// ScalarMultiply multiplies vector u with scalar
func (u Vector) ScalarMultiply(scalar float64) Vector {
	return Vector{
		u.X * scalar,
		u.Y * scalar,
		u.Z * scalar,
	}
}

// Normalize normalizes vector u
func (u Vector) Normalize() Vector {
	length := u.Magnitude()
	if length > 0 {
		return u.ScalarMultiply(1.0 / length)
	}
	return u
}

// Add adds two vectors u, v together
func (u Vector) Add(v Vector) Vector {
	return Vector{
		u.X + v.X,
		u.Y + v.Y,
		u.Z + v.Z,
	}
}

// Subtract subtracts vector v from vector u
func (u Vector) Subtract(v Vector) Vector {
	return Vector{
		u.X - v.X,
		u.Y - v.Y,
		u.Z - v.Z,
	}
}

// ConvertToRGB returns a triplet of unsigned 8-bit integers from vector u corresponding to RBG values
func (u Vector) ConvertToRGB() (uint8, uint8, uint8) {
	return uint8(255 * math.Max(0.0, math.Min(1.0, u.X))),
		uint8(255 * math.Max(0.0, math.Min(1.0, u.Y))),
		uint8(255 * math.Max(0.0, math.Min(1.0, u.Z)))
}

// Copy copies values from vector v to vector u
func (u *Vector) Copy(v Vector) {
	u.X = v.X
	u.Y = v.Y
	u.Z = v.Z
}
