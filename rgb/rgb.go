package rgb

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// adjustLevels manually adjust the red, green, and blue levels of an image.
//
// Parameters:
//   - img: The input image
//   - red: The desired red level (0-100)
//   - green: The desired green level (0-100)
//   - blue: The desired blue level (0-100)
//
// Returns:
//   - image.Image
func adjustLevels(img image.Image, red, green, blue int) image.Image {

	red = utils.ClampGeneric(red, 1, 100)
	green = utils.ClampGeneric(green, 1, 100)
	blue = utils.ClampGeneric(blue, 1, 100)

	bounds := img.Bounds()
	newImage := image.NewRGBA64(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			rr := (r * uint32(red)) / 100
			gg := (g * uint32(green)) / 100
			bb := (b * uint32(blue)) / 100

			newImage.Set(x, y, color.RGBA64{
				uint16(rr),
				uint16(gg),
				uint16(bb),
				uint16(a),
			})
		}
	}

	return newImage

}
