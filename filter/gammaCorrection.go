package filter

import (
	"image"
	"image/color"
	"math"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// GammaCorrection applies gamma correction to an image.
//
// Parameters:
//   - img: The input image
//   - gamma: Gamma correction factor (> 1)
//
// Returns:
//   - image.Image
func GammaCorrection(img image.Image, gamma float64) image.Image {

	gamma = float64(utils.ClampGeneric(int(gamma), 1, 20))
	bounds := img.Bounds()
	newImage := image.NewRGBA64(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			stretch := func(value uint32) uint16 {
				normalized := float64(value) / 65535.0
				corrected := math.Pow(normalized, gamma)
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
