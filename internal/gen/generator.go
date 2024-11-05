package gen

import (
	"math/cmplx"
)

type SetGenerator interface {
	GenerateSet(width, height int, xMin, yMin, xMax, yMax float64) [][]uint8
}

type MandelbrotGenerator struct {
	Iterations int
	Contrast   int
}

func (g *MandelbrotGenerator) mandelbrot(c complex128) uint8 {
	var z complex128
	for n := uint8(0); n < uint8(g.Iterations); n++ {
		z = z*z + c
		if cmplx.Abs(z) > 2 {
			return 255 - uint8(g.Contrast)*n
		}
	}
	return 0
}

func (g *MandelbrotGenerator) GenerateSet(width, height int, xMin, yMin, xMax, yMax float64) [][]uint8 {
	set := make([][]uint8, height)
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(yMax-yMin) + yMin
		set[py] = make([]uint8, width)
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xMax-xMin) + xMin
			z := complex(x, y)
			set[py][px] = g.mandelbrot(z)
		}
	}
	return set
}
