package edgedetection

import (
	"image"
	"image/color"
	"math"

	"github.com/BrunoPoiano/imgeffects/blur"
	"github.com/BrunoPoiano/imgeffects/utils"
)

// LaplacianOfGaussian applies edge detection using the Laplacian of Gaussian (LoG) method.
//
// Parameters:
//   - img: The input image
//   - blur_level: level of blur for the image: 0 to 20
//   - scaling: scaling factor for the edge detection: 5 to 20
//
// Returns:
//   - image.Image
func LaplacianOfGaussian(img image.Image, blur_level int, scaling float64) image.Image {
	bounds := img.Bounds()
	newImage := image.NewGray(bounds)
	bluredImage := blur.GaussianBlur(img, blur_level)

	kernel := [][]int{
		{0, 1, 0},
		{1, -4, 1},
		{0, 1, 0},
	}

	for y := 1; y < bounds.Max.Y-1; y++ {
		for x := 1; x < bounds.Max.X-1; x++ {

			var sum float64
			for ky := 0; ky < 3; ky++ {
				for kx := 0; kx < 3; kx++ {
					r, g, b, _ := bluredImage.At(x+kx-1, y+ky-1).RGBA()
					pixel := utils.Luminance8bit(r, g, b)
					sum += pixel * float64(kernel[ky][kx])
				}
			}

			val := math.Abs(sum) * scaling
			if val > 255 {
				val = 255
			}

			newImage.Set(x, y, color.Gray{uint8(val)})
		}
	}

	return newImage

}
