package export

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func SaveImage(set [][]color.RGBA, filename string) error {
	height := len(set)
	width := len(set[0])
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, set[y][x])
		}
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}
