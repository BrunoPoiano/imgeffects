package flip

import (
	"image"
)

// FlipHorizontal flips the given image horizontally.
//
// Parameters:
//   - img: The input image
//
// Returns:
//   - A new image.Image with the horizontal flip applied
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

// FlipVertical flips the given image vertically.
//
// Parameters:
//   - img: The input image
//
// Returns:
//   - A new image.Image with the vertical flip applied
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
