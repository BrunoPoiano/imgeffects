package filter

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// SolarizeEffect applies a solarization effect to an image, which simulates the effect of tone reversal
// observed in photographic film when it's extremely overexposed during development.
//
// The effect inverts colors of pixels whose luminance exceeds a certain threshold determined by the level parameter.
// Higher level values will affect more pixels in the image, creating a more pronounced solarization effect.
//
// Parameters:
//   - img: The input image to apply the solarization effect to
//   - level: Intensity of the effect, ranging from 1 (minimal) to 100 (maximum)
//     This controls the luminance threshold above which colors will be inverted
//
// Returns:
//   - image.Image: A new image with the solarization effect applied
func SolarizeEffect(img image.Image, level int) image.Image {
	level = utils.ClampGeneric(level, 1, 100)
	bounds := img.Bounds()
	newImage := image.NewRGBA64(bounds)

	treshold_level := (65535.0 * level) / 100

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {

			r, g, b, a := img.At(x, y).RGBA()
			pixel := utils.Luminance16bit(r, g, b)

			if pixel > float64(treshold_level) {
				newImage.Set(x, y, color.RGBA64{
					uint16(65535.0 - r),
					uint16(65535.0 - g),
					uint16(65535.0 - b),
					uint16(a),
				})
			} else {
				newImage.Set(x, y, color.RGBA64{
					uint16(r),
					uint16(g),
					uint16(b),
					uint16(a),
				})
			}

		}
	}

	return newImage

}
