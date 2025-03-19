package edgedetection

import (
	"image"
	"image/color"
	"math"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// KernelOperatorBased applies edge detection using a kernel operator.
//
// Supported kernels:
//   - sobel
//   - prewitt
//   - robert-cross
//   - scharr
//
// Parameters:
//   - img: The input image
//   - kernel: The name of the kernel to use
//
// Returns:
//   - image.Image
func KernelOperatorBased(img image.Image, kernel string) image.Image {
	bounds := img.Bounds()
	newImage := image.NewGray(bounds)

	gx := [][]int{}
	gy := [][]int{}

	switch kernel {
	case "sobel":
		gx = [][]int{
			{-1, 0, 1},
			{-2, 0, 2},
			{-1, 0, 1},
		}
		gy = [][]int{
			{-1, -2, -1},
			{0, 0, 0},
			{1, 2, 1},
		}

	case "prewitt":
		gx = [][]int{
			{-1, 0, 1},
			{-1, 0, 1},
			{-1, 0, 1},
		}
		gy = [][]int{
			{-1, -1, -1},
			{0, 0, 0},
			{1, 1, 1},
		}
	case "robert-cross":
		gx = [][]int{
			{1, 0},
			{0, -1},
		}
		gy = [][]int{
			{0, 1},
			{-1, 0},
		}
	case "scharr":
		gx = [][]int{
			{-3, 0, 3},
			{-10, 0, 10},
			{-3, 0, 3},
		}
		gy = [][]int{
			{-3, -10, -3},
			{0, 0, 0},
			{3, 10, 3},
		}
	}

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {

			var sumx, sumy float64
			kernelLen := len(gx) - 1
			for yy := -1; yy < kernelLen; yy++ {
				for xx := -1; xx < kernelLen; xx++ {
					r, g, b, _ := img.At(x+xx, y+yy).RGBA()
					pixel := utils.Luminance8bit(r, g, b)

					sumx += pixel * float64(gx[yy+1][xx+1])
					sumy += pixel * float64(gy[yy+1][xx+1])

				}
			}

			gradient_mag := math.Sqrt(sumx*sumx + sumy*sumy)

			newImage.Set(x, y, color.Gray{uint8(gradient_mag)})
		}
	}

	return newImage

}
