package threshold

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// ThresholdRGB processes an image pixel by pixel. For each pixel, it determines
// which color channel (Red, Green, or Blue) has the highest value. It then creates
// a new pixel where only the dominant color channel's value is retained, setting
// the other two color channels to zero. The alpha channel is preserved from the
// original pixel. If no single channel is dominant (e.g., multiple channels have
// the same highest value, or all are zero), the resulting pixel's RGB values are
// all set to zero.
//
// Parameters:
//   - img: The input image (image.Image) to be thresholded.
//
// Returns:
//   - image.Image
func ThresholdRGB(img image.Image) image.Image {

	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {

			r, g, b, a := img.At(x, y).RGBA()

			var rr, gg, bb uint32
			if r >= g && r >= b {
				rr = r
			} else if g > r && g >= b {
				gg = g
			} else if b > r && b > g {
				bb = b
			}

			newImage.Set(x, y, color.RGBA64{uint16(rr), uint16(gg), uint16(bb), uint16(a)})
		}
	}

	return newImage

}

// MultiThresholdRGB applies a multi-level thresholding effect to an image based on
// the luminance of each pixel. It divides pixels into five ranges based on their
// luminance (calculated using 16-bit precision) and assigns colors accordingly.
//
// Pixels with luminance >= 53083 (approx. 81% of 65535) are set to white.
// Pixels with luminance >= 39976 (approx. 61%) are colored based on the c1 parameter.
// Pixels with luminance >= 26869 (approx. 41%) are colored based on the c2 parameter.
// Pixels with luminance >= 13762 (approx. 21%) are colored based on the c3 parameter.
// Pixels with luminance < 13762 are set to black.
//
// The c1, c2, and c3 parameters specify the color channel to preserve for their
// respective luminance ranges. Valid values are "red", "green", or "blue". If a
// parameter is not one of these values, the original pixel color is used for that range.
// For "red", only the red channel is kept, green and blue are set to 0.
// For "green", only the green channel is kept, red and blue are set to 0.
// For "blue", only the blue channel is kept, red and green are set to 0.
// The alpha channel is always preserved from the original pixel.
//
// Parameters:
//   - img: The input image (image.Image) to be thresholded.
//   - c1:  The color ("red", "green", "blue", or other for original) for the luminance range [39976, 53083).
//   - c2:  The color ("red", "green", "blue", or other for original) for the luminance range [26869, 39976).
//   - c3:  The color ("red", "green", "blue", or other for original) for the luminance range [13762, 26869).
//
// Returns:
//   - image.Image
func MultiThresholdRGB(img image.Image, c1, c2, c3 string) image.Image {

	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	newColor := func(r, g, b, a uint32, c string) color.Color {

		var rr, gg, bb uint32

		switch c {
		case "red":
			rr = r
			gg = 0
			bb = 0
		case "green":
			rr = 0
			gg = g
			bb = 0
		case "blue":
			rr = 0
			gg = 0
			bb = b
		default:
			rr = r
			gg = g
			bb = b
		}

		return color.RGBA64{uint16(rr), uint16(gg), uint16(bb), uint16(a)}

	}

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {

			r, g, b, a := img.At(x, y).RGBA()
			pixel := utils.Luminance16bit(r, g, b)

			if pixel >= 53083 { // ~81%
				rr := uint32(65535)
				gg := uint32(65535)
				bb := uint32(65535)
				newImage.Set(x, y, color.RGBA64{uint16(rr), uint16(gg), uint16(bb), uint16(a)})

			} else if pixel >= 39976 { // ~61%

				newImage.Set(x, y, newColor(r, g, b, a, c1))
			} else if pixel >= 26869 { // ~41%

				newImage.Set(x, y, newColor(r, g, b, a, c2))
			} else if pixel >= 13762 { // ~21%

				newImage.Set(x, y, newColor(r, g, b, a, c3))
			} else { // < 21%
				rr := uint32(0)
				gg := uint32(0)
				bb := uint32(0)

				newImage.Set(x, y, color.RGBA64{uint16(rr), uint16(gg), uint16(bb), uint16(a)})

			}
		}
	}

	return newImage

}
