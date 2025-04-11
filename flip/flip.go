package flip

import (
	"image"
)

// FlipHorizontal flips the given image horizontally (mirror image along the vertical axis).
// Each pixel in the result image is positioned at a horizontally mirrored location
// from its position in the original image.
//
// Parameters:
//   - img: The input image to be flipped horizontally. Must not be nil.
//
// Returns:
//   - image.Image: A new NRGBA64 image containing the horizontally flipped image.
//     The dimensions of the returned image are the same as the input image.
func FlipHorizontal(img image.Image) image.Image {

	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
		newX := bounds.Max.X - 1
		for x := 0; x < bounds.Max.X; x++ {
			pixel := img.At(newX, y)
			newImage.Set(x, y, pixel)
			newX--
		}
	}

	return newImage
}

// FlipVertical flips the given image vertically (mirror image along the horizontal axis).
// Each pixel in the result image is positioned at a vertically mirrored location
// from its position in the original image.
//
// Parameters:
//   - img: The input image to be flipped vertically. Must not be nil.
//
// Returns:
//   - image.Image: A new NRGBA64 image containing the vertically flipped image.
//     The dimensions of the returned image are the same as the input image.
func FlipVertical(img image.Image) image.Image {

	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	for x := 0; x < bounds.Max.X; x++ {
		newY := bounds.Max.Y - 1
		for y := 0; y < bounds.Max.Y; y++ {
			pixel := img.At(x, newY)
			newImage.Set(x, y, pixel)
			newY--
		}
	}

	return newImage
}
