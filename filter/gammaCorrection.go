package filter

import (
	"image"
	"image/color"
	"math"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// GammaCorrection applies a non-linear adjustment to image luminance by raising each color
// channel to the power of gamma (output = input^gamma).
//
// This transformation affects the brightness and contrast of an image non-uniformly:
//   - Gamma > 1: Darkens the image, with more effect on midtones than shadows
//   - Gamma < 0: Inverts the gamma effect (uses 1/-gamma), brightening the image
//   - Gamma = 0: Defaults to gamma = 1 (no change)
//
// The function clamps gamma values between -10 and 10 for predictable results.
// Alpha channel values remain unchanged during the transformation.
//
// Parameters:
//   - img: The source image to transform
//   - gamma: Gamma correction factor (will be clamped to range [-10, 10])
//
// Returns:
//   - image.Image
func GammaCorrection(img image.Image, gamma float64) image.Image {
	bounds := img.Bounds()
	newImage := image.NewRGBA64(bounds)

	gamma = float64(utils.ClampGeneric(int(gamma), -10, 10))
	if gamma == 0 {
		gamma = 1
	}
	effectiveGamma := gamma
	if gamma < 0 {
		effectiveGamma = 1.0 / -gamma
	}

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			stretch := func(value uint32) uint16 {
				normalized := float64(value) / 65535.0
				corrected := math.Pow(normalized, effectiveGamma)
				return uint16(corrected * 65535)
			}

			newImage.Set(x, y, color.RGBA64{
				stretch(r),
				stretch(g),
				stretch(b),
				uint16(a),
			})
		}
	}

	return newImage
}
