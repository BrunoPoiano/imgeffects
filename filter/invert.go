package filter

import (
	"image"
	"image/color"
)

// Invert creates a negative image by inverting all color channels of the input image.
// This function processes each pixel by subtracting its RGB values from the maximum
// possible value (65535 for 16-bit color depth), creating a photographic negative effect.
// The alpha channel remains unchanged to preserve the original transparency.
//
// Parameters:
//   - img: The source image to be inverted (can be any image.Image implementation)
//
// Returns:
//   - image.Image
func Invert(img image.Image) image.Image {

	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {

			r, g, b, a := img.At(x, y).RGBA()

			rr := uint16(65535 - r)
			gg := uint16(65535 - g)
			bb := uint16(65535 - b)

			newImage.Set(x, y, color.RGBA64{rr, gg, bb, uint16(a)})
		}
	}

	return newImage

}
