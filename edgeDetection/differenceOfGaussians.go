package edgedetection

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/blur"
	"github.com/BrunoPoiano/imgeffects/utils"
)

// DifferenceOfGaussians applies edge detection using the Difference of Gaussians (DoG) method.
// The DoG algorithm works by subtracting two differently blurred versions of the original image.
// Areas where pixel values change rapidly (edges) will be highlighted in the resulting image.
//
// Parameters:
//   - img: The input image to perform edge detection on
//   - img_one_blur: Blur intensity for the first Gaussian filter (range: 0 to 20)
//     Lower values preserve more detail
//   - img_two_blur: Blur intensity for the second Gaussian filter (range: 0 to 20)
//     Higher values detect stronger edges
//
// The difference between the two blur intensities controls the size and type of
// edges that will be detected. Larger differences highlight broader edges.
//
// Returns:
//   - image.Image: A grayscale image with detected edges
func DifferenceOfGaussians(img image.Image, img_one_blur, img_two_blur int) image.Image {
	bounds := img.Bounds()
	newImage := image.NewGray(bounds)

	img_blured_one := blur.GaussianBlur(img, utils.ClampGeneric(img_one_blur, 0, 20))
	img_blured_two := blur.GaussianBlur(img, utils.ClampGeneric(img_two_blur, 0, 20))

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {

			b_r, b_g, b_b, _ := img_blured_one.At(x, y).RGBA()
			b_pixel := utils.Luminance8bit(b_r, b_g, b_b)

			bb_r, bb_g, bb_b, _ := img_blured_two.At(x, y).RGBA()
			bb_pixel := utils.Luminance8bit(bb_r, bb_g, bb_b)

			newImage.Set(x, y, color.Gray{uint8(b_pixel - bb_pixel)})
		}
	}

	return newImage
}
