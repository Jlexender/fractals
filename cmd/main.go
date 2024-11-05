package main

import (
	"log"
	"mandelbrotSet/internal/export"
	"mandelbrotSet/internal/gen"
)

func main() {
	const (
		width, height          = 16384, 16384
		xMin, yMin, xMax, yMax = -2, -2, +2, +2
		iterations             = 1000
		contrast               = 15
	)

	generator := &gen.MandelbrotGenerator{
		Iterations: iterations,
		Contrast:   contrast,
	}

	set := generator.GenerateSet(width, height, xMin, yMin, xMax, yMax)

	err := export.Image(set, "output.png")
	if err != nil {
		log.Fatalf("failed to export image: %v", err)
	}
}
