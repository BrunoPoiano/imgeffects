package threshold

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// GlobalThreshold applies a global binary threshold to an image, converting it to a black and white image.
//
// The function calculates the luminance of each pixel and compares it to the threshold level.
// Pixels with luminance above the threshold become white (255), while pixels below become black (0).
//
// Parameters:
//   - img: The input image to be thresholded
//   - level: The threshold level, ranging from 1 to 100 (percentage of maximum pixel value)
//     Lower values create more white pixels, higher values create more black pixels
//
// Returns:
//   - image.Image: A new grayscale image with binary (black and white) pixels
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

// GlobalThresholdColor applies a global binary threshold to an image, preserving color channels.
//
// The function evaluates each color channel (R,G,B) independently against the threshold.
// For each channel, values above the threshold become maximum (65535), while values below become zero (0).
// The alpha channel remains unchanged, preserving transparency information.
//
// Parameters:
//   - img: The input image to be thresholded
//   - level: The threshold level, ranging from 1 to 100 (percentage of maximum pixel value)
//     Lower values create more bright colors, higher values create more dark areas
//
// Returns:
//   - image.Image
func GlobalThresholdColor(img image.Image, level int) image.Image {
	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)
	level = utils.ClampGeneric(level, 1, 100)
	treshold_level := (65535 * level) / 100

	colorFun := func(c uint32) uint16 {
		if float64(c) > float64(treshold_level) {
			c = 65535
		} else {
			c = 0
		}

		return uint16(c)
	}

	for y := 1; y < bounds.Max.Y-1; y++ {
		for x := 1; x < bounds.Max.X-1; x++ {

			r, g, b, a := img.At(x, y).RGBA()

			newImage.Set(x, y, color.RGBA64{colorFun(r), colorFun(g), colorFun(b), uint16(a)})
		}
	}

	return newImage
}
