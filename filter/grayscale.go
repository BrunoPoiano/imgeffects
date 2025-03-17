package filter

import (
	"image"
	"image/color"
)

// GrayScale16 converts the given image to grayscale using 16-bit color depth.
//
// Parameters:
//   - img: The input image
//
// Returns:
//   - image.Image
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

// GrayScale converts the given image to grayscale.
//
// Parameters:
//   - img: The input image
//
// Returns:
//   - image.Image
func GrayScale(img image.Image) image.Image {
	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			newImage.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}
	return newImage
