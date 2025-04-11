package dithering

import (
	"image"
	"image/color"
	"math"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// OrderedDithering applies an ordered dithering effect to an image using a Bayer matrix pattern.
// Ordered dithering creates a retro visual effect by systematically applying threshold patterns
// to reduce color depth while maintaining visual structure.
//
// Parameters:
//   - img: The input image to be processed
//   - level: The number of quantization levels (1 - 20); higher values result in more color bands
//     and less pronounced dithering effect. Values outside this range will be clamped.
//   - size: The size of the dithering matrix (must be a power of 2, e.g., 2, 4, 8);
//     larger matrices create more complex dithering patterns. Non-power-of-2 values
//     will be adjusted to the next even number.
//
// Returns:
//   - A new image.Image with the ordered dithering effect applied, in RGBA64 format
//
// Note: The alpha channel is also dithered by default. The function automatically
// handles bounds checking and matrix size adjustments.
func OrderedDithering(img image.Image, level, size int) image.Image {
	bounds := img.Bounds()
	newImage := image.NewRGBA64(bounds)
	level = utils.ClampGeneric(level, 1, 20)
	threshold := thresholdMatrix(size)

	if size%2 != 0 {
		size++
	}

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

	return matrix
}
