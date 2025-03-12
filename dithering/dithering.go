package dithering

import (
	"image"
	"image/color"
	"math"
)

// ErrorDiffusionDithering applies an error diffusion dithering effect to an image.
//
// Supported algorithms:
//   - floyd-steinberg
//   - false-floyd-steinberg
//   - jarvis-judice-ninke
//   - stucki
//   - atkinson
//   - sierra
//   - two-row-seirra
//   - sierra-lite
//
// Parameters:
//   - img: The input image
//   - algorithm: The name of the dithering algorithm to use
//   - level: The number of quantization levels (1 - 10)
//
// Returns:
//   - A new image.Image with the dithering effect applied
func ErrorDifusionDithering(img image.Image, algorithm string, level int) image.Image {
	quantize := func(value uint8, levels int) (uint8, int) {
		scale := 255.0 / float64(levels-1)
		newValue := uint8(math.Round(float64(value)*float64(levels-1)/255.0) * scale)
		return newValue, int(value) - int(newValue)
	}

	bounds := img.Bounds()
	image := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			image.Set(x, y, img.At(x, y))
		}
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			oldPix := image.RGBAAt(x, y)

			newR, errR := quantize(oldPix.R, level)
			newG, errG := quantize(oldPix.G, level)
			newB, errB := quantize(oldPix.B, level)
			newA, errA := quantize(oldPix.A, level)

			image.SetRGBA(x, y, color.RGBA{uint8(newR), uint8(newG), uint8(newB), uint8(newA)})

			switch algorithm {
			case "floyd-steinberg":
				makeDither(image, x+1, y, errR, errG, errB, errA, 7.0/16)
				makeDither(image, x-1, y+1, errR, errG, errB, errA, 3.0/16)
				makeDither(image, x, y+1, errR, errG, errB, errA, 5.0/16)
				makeDither(image, x+1, y+1, errR, errG, errB, errA, 1.0/16)

			case "false-floyd-steinberg":
				makeDither(image, x+1, y, errR, errG, errB, errA, 3.0/8)
				makeDither(image, x, y+1, errR, errG, errB, errA, 3.0/8)
				makeDither(image, x+1, y+1, errR, errG, errB, errA, 2.0/8)

			case "jarvis-judice-ninke":
				makeDither(image, x+1, y, errR, errG, errB, errA, 7.0/48)
				makeDither(image, x+2, y, errR, errG, errB, errA, 5.0/48)
				makeDither(image, x-2, y+1, errR, errG, errB, errA, 3.0/48)
				makeDither(image, x-1, y+1, errR, errG, errB, errA, 5.0/48)
				makeDither(image, x, y+1, errR, errG, errB, errA, 7.0/48)
				makeDither(image, x+1, y+1, errR, errG, errB, errA, 5.0/48)
				makeDither(image, x+2, y+1, errR, errG, errB, errA, 3.0/48)
				makeDither(image, x-2, y+2, errR, errG, errB, errA, 1.0/48)
				makeDither(image, x-1, y+2, errR, errG, errB, errA, 3.0/48)
				makeDither(image, x, y+2, errR, errG, errB, errA, 5.0/48)
				makeDither(image, x+1, y+2, errR, errG, errB, errA, 3.0/48)
				makeDither(image, x+2, y+2, errR, errG, errB, errA, 1.0/48)

			case "stucki":
				makeDither(image, x+1, y, errR, errG, errB, errA, 8.0/42)
				makeDither(image, x+2, y, errR, errG, errB, errA, 4.0/42)
				makeDither(image, x-2, y+1, errR, errG, errB, errA, 2.0/42)
				makeDither(image, x-2, y+1, errR, errG, errB, errA, 4.0/42)
				makeDither(image, x, y+1, errR, errG, errB, errA, 8.0/42)
				makeDither(image, x+1, y+1, errR, errG, errB, errA, 4.0/42)
				makeDither(image, x+2, y+1, errR, errG, errB, errA, 2.0/42)
				makeDither(image, x-2, y+2, errR, errG, errB, errA, 1.0/42)
				makeDither(image, x-1, y+2, errR, errG, errB, errA, 2.0/42)
				makeDither(image, x, y+2, errR, errG, errB, errA, 4.0/42)
				makeDither(image, x+1, y+2, errR, errG, errB, errA, 2.0/42)
				makeDither(image, x+2, y+2, errR, errG, errB, errA, 1.0/42)

			case "atkinson":
				makeDither(image, x+1, y, errR, errG, errB, errA, 1.0/8)
				makeDither(image, x+2, y, errR, errG, errB, errA, 1.0/8)
				makeDither(image, x-1, y+1, errR, errG, errB, errA, 1.0/8)
				makeDither(image, x, y+1, errR, errG, errB, errA, 1.0/8)
				makeDither(image, x+1, y+1, errR, errG, errB, errA, 1.0/8)
				makeDither(image, x, y+2, errR, errG, errB, errA, 1.0/8)

			case "Sierra":
				makeDither(image, x+1, y, errR, errG, errB, errA, 5.0/32)
				makeDither(image, x+2, y, errR, errG, errB, errA, 3.0/32)
				makeDither(image, x+2, y+1, errR, errG, errB, errA, 2.0/32)
				makeDither(image, x+1, y+1, errR, errG, errB, errA, 4.0/32)
				makeDither(image, x, y+1, errR, errG, errB, errA, 5.0/32)
				makeDither(image, x+2, y+1, errR, errG, errB, errA, 2.0/32)
				makeDither(image, x-1, y+2, errR, errG, errB, errA, 2.0/32)
				makeDither(image, x, y+2, errR, errG, errB, errA, 3.0/32)
				makeDither(image, x+1, y+2, errR, errG, errB, errA, 2.0/32)

			case "two-row-seirra":
				makeDither(image, x+1, y, errR, errG, errB, errA, 4.0/16)
				makeDither(image, x+2, y, errR, errG, errB, errA, 3.0/16)
				makeDither(image, x-2, y+1, errR, errG, errB, errA, 2.0/16)
				makeDither(image, x-1, y+1, errR, errG, errB, errA, 4.0/16)
				makeDither(image, x, y+1, errR, errG, errB, errA, 3.0/16)
				makeDither(image, x+1, y+1, errR, errG, errB, errA, 2.0/16)
				makeDither(image, x+2, y+1, errR, errG, errB, errA, 1.0/16)

			case "sierra-lite":
				makeDither(image, x+1, y, errR, errG, errB, errA, 2.0/4)
				makeDither(image, x-1, y+1, errR, errG, errB, errA, 1.0/4)
				makeDither(image, x, y+1, errR, errG, errB, errA, 1.0/4)
			}
		}
	}

	return image
}

func makeDither(img *image.RGBA, x, y int, r, g, b, a int, factor float64) {
	clamp := func(value int) uint8 {
		if value < 0 {
			return 0
		} else if value > 255 {
			return 255
		}
		return uint8(value)
	}

	bounds := img.Bounds()
	if x >= bounds.Min.X && x < bounds.Max.X && y >= bounds.Min.Y && y < bounds.Max.Y {
		pixel := img.RGBAAt(x, y)

		newPixel := color.RGBA{
			R: clamp(int(pixel.R) + int(float64(r)*factor)),
			G: clamp(int(pixel.G) + int(float64(g)*factor)),
			B: clamp(int(pixel.B) + int(float64(b)*factor)),
			A: clamp(int(pixel.A) + int(float64(a)*factor)),
		}
		img.SetRGBA(x, y, newPixel)
	}
}

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

	return matrix

}
