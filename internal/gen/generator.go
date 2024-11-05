package gen

import "math/cmplx"

type SetGenerator interface {
	GenerateSet(width, height int, xMin, yMin, xMax, yMax float64) [][]uint8
}

type MandelbrotGenerator struct {
	Iterations int
	Contrast   int
}

func (m *MandelbrotGenerator) GenerateSet(width, height int, xMin, yMin, xMax, yMax float64) [][]uint8 {
	set := make([][]uint8, height)
	for y := 0; y < height; y++ {
		set[y] = make([]uint8, width)
		for x := 0; x < width; x++ {
			c := complex(
				xMin+(xMax-xMin)*float64(x)/float64(width),
				yMin+(yMax-yMin)*float64(y)/float64(height),
			)
			z := complex(0, 0)
			for i := 0; i < m.Iterations; i++ {
				z = z*z + c
				if real(z)*real(z)+imag(z)*imag(z) > 4 {
					set[y][x] = uint8(i * m.Contrast)
					break
				}
			}
		}
	}
	return set
}

type JuliaGenerator struct {
	Iterations int
	Contrast   int
	C          complex128
}

func (j *JuliaGenerator) GenerateSet(width, height int, xMin, yMin, xMax, yMax float64) [][]uint8 {
	set := make([][]uint8, height)
	for y := 0; y < height; y++ {
		set[y] = make([]uint8, width)
		for x := 0; x < width; x++ {
			z := complex(
				xMin+(xMax-xMin)*float64(x)/float64(width),
				yMin+(yMax-yMin)*float64(y)/float64(height),
			)
			c := j.C
			for i := 0; i < j.Iterations; i++ {
				z = z*z + c
				if real(z)*real(z)+imag(z)*imag(z) > 4 {
					set[y][x] = uint8(i * j.Contrast)
					break
				}
			}
		}
	}
	return set
}

type NewtonGenerator struct {
	Iterations int
	Contrast   int
}

func (n *NewtonGenerator) GenerateSet(width, height int, xMin, yMin, xMax, yMax float64) [][]uint8 {
	set := make([][]uint8, height)
	for y := 0; y < height; y++ {
		set[y] = make([]uint8, width)
		for x := 0; x < width; x++ {
			z := complex(
				xMin+(xMax-xMin)*float64(x)/float64(width),
				yMin+(yMax-yMin)*float64(y)/float64(height),
			)
			for i := 0; i < n.Iterations; i++ {
				z = z - (z*z*z-1)/(5*z*z)
				if cmplx.Abs(z*z*z-1) < 1e-6 {
					set[y][x] = uint8(i * n.Contrast)
					break
				}
			}
		}
	}
	return set
}
