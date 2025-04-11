package rgb

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// AdjustLevels modifies the RGB color intensity of an image by applying scaling factors.
//
// Each color channel (red, green, blue) can be adjusted independently with percentage values.
// The function automatically clamps input values to ensure they fall within the valid range
// of 1-100, where 100 represents the original color intensity and lower values reduce intensity.
// Alpha channel values remain unchanged.
//
// Parameters:
//   - img: The source image to be processed
//   - red: Red channel intensity percentage (1-100)
//   - green: Green channel intensity percentage (1-100)
//   - blue: Blue channel intensity percentage (1-100)
//
// Returns:
//   - A new image.Image with adjusted RGB levels
func AdjustLevels(img image.Image, red, green, blue int) image.Image {

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
