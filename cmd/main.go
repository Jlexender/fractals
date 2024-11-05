package main

import (
	"log"
	"mandelbrotSet/internal/export"
	"mandelbrotSet/internal/gen"
)

func main() {
	const (
		width, height          = 2048, 2048
		xMin, yMin, xMax, yMax = -2, -2, +2, +2
		iterations             = 100
		contrast               = 15
	)

	newton := &gen.NewtonGenerator{
		Iterations: iterations,
		Contrast:   contrast,
	}

	set := newton.GenerateSet(width, height, xMin, yMin, xMax, yMax)

	err := export.Image(set, "newton.png")
	if err != nil {
		log.Fatalf("failed to export image: %v", err)
	}
}
