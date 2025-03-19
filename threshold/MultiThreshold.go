package threshold

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// MultiThreshold applies a multiples thresholds to an image.
//
// Parameters:
//   - img: The input image
//   - quantity: The number of thresholds to apply: 2-100
//
// Returns:
//   - image.Image
func MultiThreshold(img image.Image, quantity int) image.Image {

	bounds := img.Bounds()
	newImage := image.NewGray(bounds)
	levels := 100 / quantity

	var thresholds []int

	for i := 101; i > levels; i -= levels {

		calc := int((255 * i) / 100)
		if calc > 255 {
			calc = 255
		}

		thresholds = append(thresholds, calc)
	}
	thresholds = append(thresholds, 0)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {

			r, g, b, _ := img.At(x, y).RGBA()
			pixel := utils.Luminance8bit(r, g, b)

			for _, t := range thresholds {
				if pixel >= float64(t) {
					pixel = float64(t)
					break
				}
			}

			newImage.Set(x, y, color.Gray{uint8(pixel)})
		}
	}

	return newImage

}
