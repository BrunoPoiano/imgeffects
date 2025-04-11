package dithering

import (
	"image"
	"image/color"
	"math"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// ErrorDiffusionDithering applies an error diffusion dithering effect to an image.
// This technique distributes the quantization error of a pixel to neighboring
// pixels according to different distribution patterns defined by various algorithms.
//
// Supported algorithms:
//   - floyd-steinberg: Classic algorithm with good balance of quality and performance
//   - false-floyd-steinberg: Simplified version with fewer error terms
//   - jarvis-judice-ninke: Higher quality with wider error distribution (12 neighboring pixels)
//   - stucki: Modified Jarvis algorithm with improved weights
//   - atkinson: Partial error distribution that preserves detail (only distributes 3/4 of error)
//   - sierra: Good quality with moderate computational cost
//   - two-row-seirra: Two-row variant of Sierra algorithm with reduced complexity
//   - sierra-lite: Simplified one-row Sierra variant for faster processing
//   - none: No error diffusion applied (simple quantization)
//
// Parameters:
//   - img: The input image to be processed
//   - algorithm: The name of the dithering algorithm to use (case-sensitive)
//   - level: The number of quantization levels per channel (1-10), where:
//     Lower values (1-3) produce more posterized results with high contrast.
//     Higher values (8-10) produce more subtle dithering with greater color depth.
//
// Returns:
//   - image.Image: A new RGBA image with the dithering effect applied
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
