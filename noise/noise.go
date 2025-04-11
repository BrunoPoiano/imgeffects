package noise

import (
	"image"
	"image/color"
	"math/rand/v2"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// NoiseGeneratorColor generates a random color noise image where each pixel has random RGB values.
// The generated image uses the RGBA64 color model with full opacity (alpha = 65535).
//
// Parameters:
//   - width: The width of the image in pixels
//   - height: The height of the image in pixels
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

// NoiseGenerator generates a black and white (binary) noise image.
// Each pixel is determined by comparing the luminance of random RGB values
// against a 50% threshold, resulting in either black (0) or white (255) pixels.
//
// Parameters:
//   - width: The width of the image in pixels
//   - height: The height of the image in pixels
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

// NoiseGeneratorGrayScale generates a grayscale noise image with continuous tones.
// Each pixel's value is determined by calculating the luminance of random RGB values,
// creating a full range of gray shades from black to white.
//
// Parameters:
//   - width: The width of the image in pixels
//   - height: The height of the image in pixels
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
