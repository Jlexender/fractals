package main

import (
	"mandelbrotSet/internal/export"
	"mandelbrotSet/internal/gen"
)

type GenerationRequest struct {
	Width, Height          int
	XMin, YMin, XMax, YMax float64
	Iterations, Contrast   int
	Function, Derivative   func(complex128) complex128
	C                      complex128
}

func generateNewton(req GenerationRequest) {
	generator := gen.NewtonGenerator{
		Iterations: req.Iterations,
		Contrast:   req.Contrast,
		Function:   req.Function,
		Derivative: req.Derivative,
	}

	set := generator.GenerateSet(req.Width, req.Height, req.XMin, req.YMin, req.XMax, req.YMax)
	export.SaveImage(set, "output/newton.png")
}

func generateJulia(req GenerationRequest) {
	generator := gen.JuliaGenerator{
		Iterations: req.Iterations,
		Contrast:   req.Contrast,
		C:          req.C,
	}

	set := generator.GenerateSet(req.Width, req.Height, req.XMin, req.YMin, req.XMax, req.YMax)
	export.SaveImage(set, "output/julia.png")
}

func generateMandelbrot(req GenerationRequest) {
	generator := gen.MandelbrotGenerator{
		Iterations: req.Iterations,
		Contrast:   req.Contrast,
	}

	set := generator.GenerateSet(req.Width, req.Height, req.XMin, req.YMin, req.XMax, req.YMax)
	export.SaveImage(set, "output/mandelbrot.png")
}

func main() {
	req := GenerationRequest{
		Width: 800, Height: 800,
		XMin: -2, YMin: -2, XMax: 2, YMax: 2,
		Iterations: 100, Contrast: 100,
		Function: func(z complex128) complex128 {
			return z*z*z - 1
		},
		Derivative: func(z complex128) complex128 {
			return 3 * z * z
		},
		C: complex(-0.4, 0.6),
	}

	generateNewton(req)
}
