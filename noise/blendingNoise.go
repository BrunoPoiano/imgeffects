package noise

import (
	"image"
	"image/color"
	"sync"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// BlendingNoiseToImage applies a noise effect to an image.
//
// Parameters:
//   - img: The input image
//   - alpha: The blending factor, ranging from 0 to 1
//   - noiseType: The type of noise to apply: [gray, color, default]
//
// Returns:
//   - image.Image
func BlendingNoiseToImage(img image.Image, alpha float64, noiseType string) image.Image {

	bounds := img.Bounds()
	var noiseImage image.Image

	switch noiseType {
	case "gray":
		noiseImage = NoiseGeneratorGrayScale(bounds.Max.X, bounds.Max.Y)
	case "color":
		noiseImage = NoiseGeneratorColor(bounds.Max.X, bounds.Max.Y)
	default:
		noiseImage = NoiseGenerator(bounds.Max.X, bounds.Max.Y)
	}

	blendingFunc := func(start, end int, newImage *image.RGBA64, wg *sync.WaitGroup) {
		defer wg.Done()

		for y := start; y < end; y++ {
			for x := 0; x < bounds.Max.X; x++ {

				r1, g1, b1, a1 := img.At(x, y).RGBA()
				r2, g2, b2, a2 := noiseImage.At(x, y).RGBA()

				color := color.RGBA64{
					R: uint16((alpha * float64(r1)) + ((1 - alpha) * float64(r2))),
					G: uint16((alpha * float64(g1)) + ((1 - alpha) * float64(g2))),
					B: uint16((alpha * float64(b1)) + ((1 - alpha) * float64(b2))),
					A: uint16((alpha * float64(a1)) + ((1 - alpha) * float64(a2))),
				}

				newImage.SetRGBA64(x, y, color)
			}
		}
	}

	return utils.ParallelExecution(utils.ParallelExecutionStruct{Image: img, Function: blendingFunc})

}
