package noise

import (
	"image"
	"image/color"
	"math/rand/v2"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// NoiseGeneratorColor generates a color noise image.
//
// Parameters:
//   - width: The width of the image
//   - height: The height of the image
//
// Returns:
//   - image.Image
func NoiseGeneratorColor(width, height int) image.Image {

	newImage := image.NewRGBA64(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			color := color.RGBA64{
				R: uint16(rand.Uint64N(65535)),
				G: uint16(rand.Uint64N(65535)),
				B: uint16(rand.Uint64N(65535)),
				A: 65535,
			}

			newImage.SetRGBA64(x, y, color)
		}
	}

	return newImage
}

// NoiseGenerator generates a black and white noise image.
//
// Parameters:
//   - width: The width of the image
//   - height: The height of the image
//
// Returns:
//   - image.Image
func NoiseGenerator(width, height int) image.Image {

	newImage := image.NewGray(image.Rect(0, 0, width, height))
	treshold_level := (255 * 50) / 100
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			r := uint32(rand.UintN(255))
			g := uint32(rand.UintN(255))
			b := uint32(rand.UintN(255))

			pixel := utils.Luminance16bit(r, g, b)
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

// NoiseGeneratorGrayScale generates a gray scale noise image.
//
// Parameters:
//   - width: The width of the image
//   - height: The height of the image
//
// Returns:
//   - image.Image
func NoiseGeneratorGrayScale(width, height int) image.Image {

	newImage := image.NewGray(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			r := uint32(rand.UintN(255))
			g := uint32(rand.UintN(255))
			b := uint32(rand.UintN(255))

			pixel := utils.Luminance16bit(r, g, b)

			newImage.Set(x, y, color.Gray{uint8(pixel)})
		}
	}

	return newImage
}
