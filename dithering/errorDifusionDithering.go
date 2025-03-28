package dithering

import (
	"image"
	"image/color"
	"math"

	"github.com/BrunoPoiano/imgeffects/utils"
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
//   - none
//
// Parameters:
//   - img: The input image
//   - algorithm: The name of the dithering algorithm to use
//   - level: The number of quantization levels (1 - 10)
//
// Returns:
//   - image.Image
func ErrorDifusionDithering(img image.Image, algorithm string, level int) image.Image {

	level = utils.ClampGeneric(level, 1, 10)

	quantize := func(value uint8, levels int) (uint8, int) {
		scale := 255.0 / float64(levels-1)
		newValue := utils.Clamp8bit(int(math.Round(float64(value)*float64(levels-1)/255.0) * scale))

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

			case "sierra":
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

			case "none":
			}
		}
	}

	return image
}

func makeDither(img *image.RGBA, x, y int, r, g, b, a int, factor float64) {
	bounds := img.Bounds()
	if x >= bounds.Min.X && x < bounds.Max.X && y >= bounds.Min.Y && y < bounds.Max.Y {
		pixel := img.RGBAAt(x, y)

		newPixel := color.RGBA{
			R: utils.Clamp8bit(int(pixel.R) + int(float64(r)*factor)),
			G: utils.Clamp8bit(int(pixel.G) + int(float64(g)*factor)),
			B: utils.Clamp8bit(int(pixel.B) + int(float64(b)*factor)),
			A: utils.Clamp8bit(int(pixel.A) + int(float64(a)*factor)),
		}
		img.SetRGBA(x, y, newPixel)
	}
}
