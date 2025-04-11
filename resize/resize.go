package resize

import (
	"image"
	"image/color"
	"math"
)

// NewAspectRatio calculates the new dimensions for an image while preserving the aspect ratio.
// Given a new target width, this function computes the proportionally correct height
// to maintain the original image's proportions, preventing distortion.
//
// Parameters:
//   - img: The input image whose aspect ratio should be preserved
//   - newWidth: The desired width of the resized image (in pixels)
//
// Returns:
//   - int: The requested new width (same as input newWidth)
//   - int: The calculated new height that preserves aspect ratio
func NewAspectRatio(img image.Image, newWidth int) (int, int) {
	imgBounds := img.Bounds()
	aspectRatio := float64(newWidth) / float64(imgBounds.Dx())
	newHeight := int(float64(imgBounds.Dy()) * aspectRatio)

	return newWidth, newHeight
}

// NearestNeighbor resizes an image using nearest neighbor interpolation.
// This is the simplest and fastest resizing algorithm that works by selecting
// the color of the nearest pixel in the original image for each pixel in the
// resized image. While efficient, it may produce pixelated results, especially
// when significantly enlarging images.
//
// Parameters:
//   - img: The input image to be resized
//   - newWidth: The desired width of the resized image (in pixels)
//   - newHeight: The desired height of the resized image (in pixels)
//
// Returns:
//   - image.Image
func NearestNeighbor(img image.Image, newWidth, newHeight int) image.Image {
	bounds := img.Bounds()
	newImage := image.NewNRGBA64(image.Rect(0, 0, newWidth, newHeight))

	for y := 0; y < bounds.Max.Y; y++ {
		newY := y * bounds.Max.Y / newHeight
		for x := 0; x < bounds.Max.X; x++ {
			newX := x * bounds.Max.X / newWidth
			color := img.At(newX, newY)
			newImage.Set(x, y, color)
		}
	}

	return newImage
}

// BypolarInterpolate resizes an image using Bypolar interpolation.
// This method calculates pixel values in the resized image by interpolating
// between the four nearest pixels in the original image. It provides smoother
// results compared to nearest neighbor interpolation, making it suitable for
// both enlarging and reducing images while maintaining better visual quality.
//
// Parameters:
//   - img: The input image to be resized
//   - newWidth: The desired width of the resized image (in pixels)
//   - newHeight: The desired height of the resized image (in pixels)
//
// Returns:
//   - image.Image
func BypolarInterpolate(img image.Image, newWidth, newHeight int) image.Image {
	oldWidth, oldHeight := img.Bounds().Max.X, img.Bounds().Max.Y
	newImage := image.NewNRGBA64(image.Rect(0, 0, newWidth, newHeight))

	ratioX := float64(oldWidth) / float64(newWidth)
	ratioY := float64(oldHeight) / float64(newHeight)

	bilinearInterpolateColor := func(c00, c10, c01, c11 color.Color, dx, dy float64) color.RGBA64 {
		r00, g00, b00, _ := c00.RGBA()
		r10, g10, b10, _ := c10.RGBA()
		r01, g01, b01, _ := c01.RGBA()
		r11, g11, b11, _ := c11.RGBA()

		r := (1-dx)*(1-dy)*float64(r00) + dx*(1-dy)*float64(r10) + (1-dx)*dy*float64(r01) + dx*dy*float64(r11)
		g := (1-dx)*(1-dy)*float64(g00) + dx*(1-dy)*float64(g10) + (1-dx)*dy*float64(g01) + dx*dy*float64(g11)
		b := (1-dx)*(1-dy)*float64(b00) + dx*(1-dy)*float64(b10) + (1-dx)*dy*float64(b01) + dx*dy*float64(b11)

		return color.RGBA64{uint16(r), uint16(g), uint16(b), 65535.0}
	}

	for y := 0; y < newHeight; y++ {
		srcY := float64(y) * ratioY
		y0 := int(math.Floor(srcY))
		y1 := int(math.Min(float64(y0+1), float64(oldHeight-1)))
		dy := srcY - float64(y0)

		for x := 0; x < newWidth; x++ {
			srcX := float64(x) * ratioX
			x0 := int(math.Floor(srcX))
			x1 := int(math.Min(float64(x0+1), float64(oldWidth-1)))
			dx := srcX - float64(x0)

			c00 := img.At(x0, y0)
			c10 := img.At(x1, y0)
			c01 := img.At(x0, y1)
			c11 := img.At(x1, y1)

			newImage.SetRGBA64(x, y, bilinearInterpolateColor(c00, c10, c01, c11, dx, dy))
		}
	}
	return newImage
}
