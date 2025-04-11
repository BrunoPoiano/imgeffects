package noise

import (
	"image"
	"image/color"
	"sync"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// BlendingNoiseToImage applies a noise effect to an image by blending the original image
// with a generated noise pattern.
//
// Parameters:
//   - img: The input image to which noise will be applied
//   - alpha: The blending factor, ranging from 0 to 1. Higher values preserve more of
//     the original image, while lower values add more noise
//   - noiseType: The type of noise to generate and apply:
//   - "gray": Grayscale noise (same intensity across R,G,B channels)
//   - "color": RGB color noise with random values in each channel
//   - any other value: Default noise pattern
//
// Returns:
//   - image.Image
//
// The function uses parallel processing to efficiently apply the noise effect.
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
