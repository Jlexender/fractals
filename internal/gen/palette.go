package gen

import "image/color"

func getColor(iteration, maxIterations int) color.RGBA {
	if iteration == maxIterations {
		return color.RGBA{A: 255}
	}
	t := float64(iteration) / float64(maxIterations) * 2
	r := uint8(9 * (1 - t) * t * t * t * 255)
	g := uint8(15 * (1 - t) * (1 - t) * t * t * 255)
	b := uint8(8.5 * (1 - t) * (1 - t) * (1 - t) * t * 255)
	return color.RGBA{R: r, G: g, B: b, A: 255}
}
