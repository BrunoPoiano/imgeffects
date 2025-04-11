package filter

import (
	"image"
	"image/color"
)

// GrayScale16 converts the given image to grayscale using 16-bit color depth.
// This provides higher precision color conversion than the 8-bit version,
// maintaining more detail in images with subtle tonal variations.
//
// Parameters:
//   - img: The input image to be converted to grayscale
//
// Returns:
//   - image.Image: A new grayscale image with 16-bit color depth
func GrayScale16(img image.Image) image.Image {
	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			newImage.Set(x, y, color.Gray16Model.Convert(img.At(x, y)))
		}
	}
	return newImage
}

// GrayScale converts the given image to grayscale using standard 8-bit color depth.
// This function transforms a color image to grayscale by applying the standard
// luminance conversion while preserving the original image dimensions.
//
// Parameters:
//   - img: The input image to be converted to grayscale
//
// Returns:
//   - image.Image: A new grayscale image with 8-bit color depth
func GrayScale(img image.Image) image.Image {
	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			newImage.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}
	return newImage
}
