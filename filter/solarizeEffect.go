package filter

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// SolarizeEffect is the effect of tone reversal observed in cases of extreme overexposure of the photographic film in the camera.
//
// Parameters:
//   - img: The input image
//   - level: Filter size from 1 to 100
//
// Returns:
//   - image.Image
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
