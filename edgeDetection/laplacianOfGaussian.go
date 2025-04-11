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
// The LoG operator first applies Gaussian blur to reduce noise, then uses the Laplacian
// operator to detect edges by finding zero crossings in the second derivative of the image.
// This method is effective at finding edges and ignoring noise.
//
// Parameters:
//   - img: The input image to detect edges on
//   - blur_level: Controls the Gaussian blur intensity (range: 0 to 20)
//     Lower values preserve more detail, higher values reduce noise
//   - scaling: Amplification factor for edge intensity (range: 5 to 20)
//     Higher values make edges more pronounced in the output
//
// Returns:
//   - image.Image: A grayscale image with detected edges
func LaplacianOfGaussian(img image.Image, blur_level int, scaling float64) image.Image {
	bounds := img.Bounds()
	newImage := image.NewGray(bounds)

	blur_level = utils.ClampGeneric(blur_level, 0, 20)
	scaling = float64(utils.ClampGeneric(int(scaling), 5, 20))

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
