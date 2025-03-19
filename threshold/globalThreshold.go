package threshold

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// GlobalThreshold applies a global threshold to an image.
//
// Parameters:
//   - img: The input image
//   - level: The threshold level, ranging from 1 to 100
//
// Returns:
//   - image.Image
func GlobalThreshold(img image.Image, level int) image.Image {
	bounds := img.Bounds()
	newImage := image.NewGray(bounds)
	level = utils.ClampGeneric(level, 1, 100)
	treshold_level := (255 * level) / 100

	for y := 1; y < bounds.Max.Y-1; y++ {
		for x := 1; x < bounds.Max.X-1; x++ {

			r, g, b, _ := img.At(x, y).RGBA()
			pixel := utils.Luminance8bit(r, g, b)

			if pixel > float64(treshold_level) {
				pixel = 255
			} else {
				pixel = 0
			}

			newImage.Set(x, y, color.Gray{uint8(pixel)})
		}
	}

	return newImage

}
