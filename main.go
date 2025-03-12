package imgeffects

import (
	"image"

	"github.com/BrunoPoiano/imgeffects/blur"
	"github.com/BrunoPoiano/imgeffects/dithering"
	"github.com/BrunoPoiano/imgeffects/flip"
	"github.com/BrunoPoiano/imgeffects/hsl"
	"github.com/BrunoPoiano/imgeffects/kuwahara"
)

const Version = "0.1.0"

// GaussianBlur applies a Gaussian blur filter to an image.
//
// Parameters:
//   - img: The input image
//   - size: Kernel size from 0 to 20
//
// Returns:
//   - A new image.Image with the blur effect applied
func GaussianBlur(img image.Image, size int) image.Image {
	return blur.GaussianBlur(img, size)
}

// ErrorDiffusionDithering applies an error diffusion dithering effect to an image.
//
// Supported algorithms:
//   - floyd-steinberg
//   - false-floyd-steinberg
//   - jarvis-judice-ninke
//   - stucki
//   - atkinson
//   - sierra
//   - two-row-seirra
//   - sierra-lite
//
// Parameters:
//   - img: The input image
//   - algorithm: The name of the dithering algorithm to use
//   - level: The number of quantization levels (1 - 10)
//
// Returns:
//   - A new image.Image with the dithering effect applied
func ErrorDiffusionDithering(img image.Image, algorithm string, level int) image.Image {
	return dithering.ErrorDifusionDithering(img, algorithm, level)
}

// OrderedDithering applies an ordered dithering effect to an image.
//
// Parameters:
//   - img: The input image
//   - level: The number of quantization levels (1 - 10)
//   - size: The size of the dithering matrix (must be multiple of 2)
//
// Returns:
//   - A new image.Image with the dithering effect applied
func OrderedDithering(img image.Image, level, size int) image.Image {
	return dithering.OrderedDithering(img, level, size)
}

// FlipHorizontal flips the given image horizontally.
//
// Parameters:
//   - img: The input image
//
// Returns:
//   - A new image.Image with the horizontal flip applied
func FlipHorizontal(img image.Image) image.Image {
	return flip.FlipHorizontal(img)
}

// FlipVertical flips the given image vertically.
//
// Parameters:
//   - img: The input image
//
// Returns:
//   - A new image.Image with the vertical flip applied
func FlipVertical(img image.Image) image.Image {
	return flip.FlipVertical(img)
}

// AdjustHue changes the hue of an image.
//
// Parameters:
//   - img: The input image
//   - change: Hue shift in degrees (0-360)
//
// Returns:
//   - A new image.Image with the hue adjustment applied
func AdjustHue(img image.Image, change int) image.Image {
	return hsl.Hue(img, change)
}

// AdjustSaturation changes the saturation of an image.
//
// Parameters:
//   - img: The input image
//   - change: Saturation adjustment (-1.0 to 1.0)
//
// Returns:
//   - A new image.Image with the saturation adjustment applied
func AdjustSaturation(img image.Image, change float64) image.Image {
	return hsl.Saturation(img, change)
}

// AdjustLuminance changes the luminance of an image.
//
// Parameters:
//   - img: The input image
//   - change: Luminance adjustment (-1.0 to 1.0)
//
// Returns:
//   - A new image.Image with the luminance adjustment applied
func AdjustLuminance(img image.Image, change float64) image.Image {
	return hsl.Luminance(img, change)
}

// KuwaharaFilter applies the Kuwahara filter to an image.
//
// Parameters:
//   - img: The input image
//   - size: Filter size from 1 to 20
//
// Returns:
//   - A new image.Image with the Kuwahara filter applied
func KuwaharaFilter(img image.Image, size int) image.Image {
	return kuwahara.KuwaharaFilter(img, size)
}
