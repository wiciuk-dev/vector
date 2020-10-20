package main

import "math"

// Vector represents a mathematical n-dimensional vector
type Vector []float64

// Equal compares a and b elementwise under a precision eps.
func (a Vector) Equal(b Vector, eps float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if math.Abs(a[i]-b[i]) > eps {
			return false
		}
	}
	return true
}

// Add sums a and b elementwise.
// Add panics if dimensions are different.
func (a Vector) Add(b Vector) Vector {
	if len(a) != len(b) {
		panic("different dimensions")
	}
	c := make(Vector, len(a))
	for i := range a {
		c[i] = a[i] + b[i]
	}
	return c
}

// Negate changes the sign of all elements of a.
func (a Vector) Negate() Vector {
	b := make(Vector, len(a))
	for i := range a {
		b[i] = -a[i]
	}
	return b
}

// Sub subtracts a and b elementwise.
// Sub panics if dimensions are different.
func (a Vector) Sub(b Vector) Vector {
	return a.Add(b.Negate())
}

// HadamardProduct multiplies a and b elementwise.
// HadamardProduct panics if dimensions are different.
func (a Vector) HadamardProduct(b Vector) Vector {
	if len(a) != len(b) {
		panic("different dimensions")
	}
	c := make(Vector, len(a))
	for i := range a {
		c[i] = a[i] * b[i]
	}
	return c
}

// Sum returns the sum of elements of a.
func (a Vector) Sum() float64 {
	var sum float64
	for i := range a {
		sum += a[i]
	}
	return sum
}

// Dot performs the dot product of a and b.
// Dot panics if dimensions are different.
func (a Vector) Dot(b Vector) float64 {
	return a.HadamardProduct(b).Sum()
}

// Lenght returns the length of a.
func (a Vector) Lenght() float64 {
	return math.Sqrt(a.Dot(a))
}

// Angle computes the angle between a and b.
// Angle panics if a or b are zero.
func (a Vector) Angle(b Vector) float64 {
	m, n := a.Lenght(), b.Lenght()
	if m == 0 || n == 0 {
		panic("vector is zero")
	}
	return math.Acos(a.Dot(b) / (m * n))
}

// Cross computes the cross product of a and b.
// Cross panics if a and b are not 3D.
func (a Vector) Cross(b Vector) Vector {
	if len(a) != 3 || len(b) != 3 {
		panic("vectors are not 3D")
	}
	c := make(Vector, 3)
	c[0] = a[1]*b[2] - a[2]*b[1]
	c[1] = a[2]*b[0] - a[0]*b[2]
	c[2] = a[0]*b[1] - a[1]*b[0]
	return c
}
