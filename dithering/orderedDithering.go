package dithering

import (
	"image"
	"image/color"
	"math"
)

// OrderedDithering applies an ordered dithering effect to an image.
//
// Parameters:
//   - img: The input image
//   - level: The number of quantization levels (1 - 10)
//   - size: The size of the dithering matrix (must be multiple of 2)
//
// Returns:
//   - A new image.Image with the dithering effect applied
func OrderedDithering(img image.Image, level, size int) image.Image {
	bounds := img.Bounds()
	newImage := image.NewRGBA64(bounds)

	threshold := thresholdMatrix(size)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			tx := x % size
			ty := y % size

			th := threshold[tx][ty]

			newImage.SetRGBA64(x, y, color.RGBA64{
				uint16(orderedDither(uint64(r), level, th)),
				uint16(orderedDither(uint64(g), level, th)),
				uint16(orderedDither(uint64(b), level, th)),
				uint16(orderedDither(uint64(a), level, th)),
			})
		}
	}
	return newImage
}

func orderedDither(value uint64, level int, theshold float64) uint64 {
	normalizedValue := float64(value) / 65535.0
	adjustedValue := normalizedValue + (theshold-0.5)/float64(level)
	quantizeValue := math.Floor(adjustedValue*float64(level-1)+0.5) / float64(level-1)

	return uint64(math.Min(math.Max(quantizeValue*65535.0, 0.0), 65535.0))
}

func thresholdMatrix(size int) [][]float64 {

	if size < 2 {
		size = 2
	} else if size%2 != 0 {
		size++
	}

	matrix := make([][]float64, size)
	for i := range matrix {
		matrix[i] = make([]float64, size)
	}

	if size == 2 {
		matrix[0][0] = 0.0
		matrix[0][1] = 0.5
		matrix[1][0] = 0.75
		matrix[1][1] = 0.25

		return matrix
	}

	halfSize := size / 2
	smallerMatrix := thresholdMatrix(halfSize)

	for x := 0; x < halfSize; x++ {
		for y := 0; y < halfSize; y++ {
			value := smallerMatrix[x][y]
			matrix[x][y] = value / 4.0
			matrix[x][y+halfSize] = value/4 + 0.5
			matrix[x+halfSize][y] = value/4 + 0.75
			matrix[x+halfSize][y+halfSize] = value/4 + 0.25
		}
	}

	println(matrix)
	return matrix
}
