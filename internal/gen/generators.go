package gen

import (
	"image/color"
	"math/cmplx"
)

type SetGenerator interface {
	GenerateSet(width, height int, xMin, yMin, xMax, yMax float64) [][]uint8
}

type MandelbrotGenerator struct {
	Iterations int
	Contrast   int
}

func (m *MandelbrotGenerator) GenerateSet(width, height int, xMin, yMin, xMax, yMax float64) [][]color.RGBA {
	set := make([][]color.RGBA, height)
	for y := 0; y < height; y++ {
		set[y] = make([]color.RGBA, width)
		for x := 0; x < width; x++ {
			c := complex(
				xMin+(xMax-xMin)*float64(x)/float64(width),
				yMin+(yMax-yMin)*float64(y)/float64(height),
			)
			z := complex(0, 0)
			var i int
			for i = 0; i < m.Iterations; i++ {
				z = z*z + c
				if real(z)*real(z)+imag(z)*imag(z) > 4 {
					break
				}
			}
			set[y][x] = getColor(i, m.Iterations)
		}
	}
	return set
}

type JuliaGenerator struct {
	Iterations int
	Contrast   int
	C          complex128
}

func (j *JuliaGenerator) GenerateSet(width, height int, xMin, yMin, xMax, yMax float64) [][]color.RGBA {
	set := make([][]color.RGBA, height)
	for y := 0; y < height; y++ {
		set[y] = make([]color.RGBA, width)
		for x := 0; x < width; x++ {
			z := complex(
				xMin+(xMax-xMin)*float64(x)/float64(width),
				yMin+(yMax-yMin)*float64(y)/float64(height),
			)
			var i int
			for i = 0; i < j.Iterations; i++ {
				z = z*z + j.C
				if real(z)*real(z)+imag(z)*imag(z) > 4 {
					break
				}
			}
			set[y][x] = getColor(i, j.Iterations)
		}
	}
	return set
}

type NewtonGenerator struct {
	Iterations int
	Contrast   int
	Function   func(complex128) complex128
	Derivative func(complex128) complex128
}

func (n *NewtonGenerator) GenerateSet(width, height int, xMin, yMin, xMax, yMax float64) [][]color.RGBA {
	set := make([][]color.RGBA, height)
	for y := 0; y < height; y++ {
		set[y] = make([]color.RGBA, width)
		for x := 0; x < width; x++ {
			z := complex(
				xMin+(xMax-xMin)*float64(x)/float64(width),
				yMin+(yMax-yMin)*float64(y)/float64(height),
			)
			var i int
			for i = 0; i < n.Iterations; i++ {
				z = z - n.Function(z)/n.Derivative(z)
				if cmplx.Abs(n.Function(z)) < 1e-6 {
					break
				}
			}
			set[y][x] = getColor(i, n.Iterations)
		}
	}
	return set
}
