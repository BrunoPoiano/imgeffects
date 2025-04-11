package threshold

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// MultiThreshold applies multiple thresholds to an image, creating a quantized grayscale effect.
//
// The function converts the image to grayscale and then limits the pixel values to a specific
// set of thresholds, creating distinct visual bands.
//
// Parameters:
//   - img: The input image to be processed
//   - quantity: The number of thresholds to apply, must be between 2-100 (will be clamped)
//
// Returns:
//   - image.Image
func MultiThreshold(img image.Image, quantity int) image.Image {
	quantity = utils.ClampGeneric(quantity, 2, 100)

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

// MultiThresholdColor applies multiple thresholds to an image while preserving color information.
//
// Unlike MultiThreshold which produces a grayscale image, this function quantizes each RGB color
// channel separately, resulting in a color-banded effect with distinct color plateaus.
//
// Parameters:
//   - img: The input image to be processed
//   - quantity: The number of thresholds to apply, must be between 2-10 (will be clamped)
//
// Returns:
//   - image.Image
func MultiThresholdColor(img image.Image, quantity int) image.Image {
	quantity = utils.ClampGeneric(quantity, 2, 10)

	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)
	levels := 100 / quantity

	var thresholds []int

	for i := 101; i > levels; i -= levels {

		calc := int((65535 * i) / 100)
		if calc > 65535 {
			calc = 65535
		}

		thresholds = append(thresholds, calc)
	}
	thresholds = append(thresholds, 0)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {

			r, g, b, a := img.At(x, y).RGBA()
			var rr, gg, bb uint16

			for _, t := range thresholds {
				if r >= uint32(t) {
					rr = uint16(t)
					break
				}
			}
			for _, t := range thresholds {
				if g >= uint32(t) {
					gg = uint16(t)
					break
				}
			}

			for _, t := range thresholds {
				if b >= uint32(t) {
					bb = uint16(t)
					break
				}
			}

			newImage.Set(x, y, color.RGBA64{rr, gg, bb, uint16(a)})
		}
	}

	return newImage

}
