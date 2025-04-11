package contrast

import (
	"image"
	"image/color"
	"math"
)

// LogarithmicTransformation applies logarithmic transformation to an image,
// which can enhance details in darker regions while compressing higher intensity values.
//
// The transformation follows the formula: output = log(1 + variation*pixel) / log(1 + variation)
// where pixel values are normalized to the range [0,1].
//
// Parameters:
//   - img: The input image to be transformed
//   - variation: Controls the strength and direction of the transformation
//     Range: [-1,1]. Positive values enhance dark regions, negative values enhance
//     bright regions. Values 0 and -1 are automatically adjusted to 0.1 and -0.99 respectively to avoid mathematical issues.
//
// Returns:
//   - image.Image
func LogarithmicTransformation(img image.Image, variation float64) image.Image {

	bounds := img.Bounds()
	newImage := image.NewRGBA64(bounds)

	calc := func(value uint32) uint16 {
		switch variation {
		case 0:
			variation = 0.1
		case -1:
			variation = -0.99
		}
		v := float64(value) / 65535.0
		transformed := math.Log(1+variation*v) / math.Log(1+variation)
		return uint16(transformed * 65535.0)
	}

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			newImage.Set(x, y, color.RGBA64{
				calc(r),
				calc(g),
				calc(b),
				uint16(a),
			})
		}
	}

	return newImage
}
