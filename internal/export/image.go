package export

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func Image(set [][]uint8, filename string) error {
	height := len(set)
	if height == 0 {
		return nil
	}
	width := len(set[0])

	img := image.NewGray(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.SetGray(x, y, color.Gray{Y: set[y][x]})
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return png.Encode(f, img)
}
