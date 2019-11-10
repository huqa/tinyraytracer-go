package vector

import "math"

// Vector3 defines a struct holding the X, Y, Z values of a vector
// as float64
type Vector3 struct {
	X float64
	Y float64
	Z float64
}

// NewVector3 constructs a new vector3
func NewVector3(x, y, z float64) Vector3 {
	return Vector3{
		x,
		y,
		z,
	}
}

// Vector2 defines a struct holding the X, Y values of a vector
type Vector2 struct {
	X float64
	Y float64
}

// NewVector2 constructs a new vector2
func NewVector2(x, y float64) Vector2 {
	return Vector2{
		x,
		y,
	}
}

// DotProduct returns the dot product of vectors u and v
func (u Vector3) DotProduct(v Vector3) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

// Magnitude returns the length of vector u
func (u Vector3) Magnitude() float64 {
	return math.Sqrt(u.DotProduct(u))
}

// CrossProduct returns the cross product of vectors u and v
func (u Vector3) CrossProduct(v Vector3) Vector3 {
	return Vector3{
		u.Y*v.Z - u.Z*v.Y,
		u.Z*v.X - u.X*v.Z,
		u.X*v.Y - u.Y*v.X,
	}
}

// ScalarMultiply multiplies vector u with scalar
func (u Vector3) ScalarMultiply(scalar float64) Vector3 {
	return Vector3{
		u.X * scalar,
		u.Y * scalar,
		u.Z * scalar,
	}
}

// Normalize normalizes vector u
func (u Vector3) Normalize() Vector3 {
	length := u.Magnitude()
	if length > 0 {
		return u.ScalarMultiply(1.0 / length)
	}
	return u
}

// Add adds two vectors u, v together
func (u Vector3) Add(v Vector3) Vector3 {
	return Vector3{
		u.X + v.X,
		u.Y + v.Y,
		u.Z + v.Z,
	}
}

// Subtract subtracts vector v from vector u
func (u Vector3) Subtract(v Vector3) Vector3 {
	return Vector3{
		u.X - v.X,
		u.Y - v.Y,
		u.Z - v.Z,
	}
}

// ConvertToRGB returns a triplet of unsigned 8-bit integers from vector u corresponding to RBG values
func (u Vector3) ConvertToRGB() (uint8, uint8, uint8) {
	return uint8(255 * math.Max(0.0, math.Min(1.0, u.X))),
		uint8(255 * math.Max(0.0, math.Min(1.0, u.Y))),
		uint8(255 * math.Max(0.0, math.Min(1.0, u.Z)))
}

// Copy copies values from vector v to vector u
func (u *Vector3) Copy(v Vector3) {
	u.X = v.X
	u.Y = v.Y
	u.Z = v.Z
}
