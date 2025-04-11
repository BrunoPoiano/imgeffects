package blur

import (
	"image"
	"image/color"
	"math"
	"sync"

	"github.com/BrunoPoiano/imgeffects/utils"
)

func createKernel(size int) []float64 {

	if size%2 == 0 {
		size++
	}

	sigma := float64(size) / 6.0
	kernel := make([]float64, size)
	center := size / 2
	sum := 0.0

	for i := 0; i < size; i++ {
		x := float64(i - center)
		kernel[i] = math.Exp(-(x*x)/(2*sigma*sigma)) / (math.Sqrt(2*math.Pi) * sigma)
		sum += kernel[i]
	}

	for i := 0; i < size; i++ {
		kernel[i] /= sum
	}

	return kernel
}

func applyHorizontalBlur(img image.Image, kernel []float64) image.Image {
	bounds := img.Bounds()
	padding := len(kernel) / 2

	horizontalFunc := func(start, end int, newImage *image.RGBA64, wg *sync.WaitGroup) {
		defer wg.Done()
		for y := start; y < end; y++ {
			for x := 0; x < bounds.Max.X; x++ {
				var r, g, b, a float64

				for kx := 0; kx < len(kernel); kx++ {
					sx := x + kx - padding

					if sx < 0 {
						sx = -sx
					}
					if sx >= bounds.Max.X {
						sx = 2*bounds.Max.X - sx - 1
					}

					sCol := img.At(sx, y)
					sr, sg, sb, sa := sCol.RGBA()

					weight := kernel[kx]
					r += float64(sr) * weight
					g += float64(sg) * weight
					b += float64(sb) * weight
					a += float64(sa) * weight
				}

				newImage.SetRGBA64(x, y, color.RGBA64{
					R: uint16(r),
					G: uint16(g),
					B: uint16(b),
					A: uint16(a),
				})
			}
		}
	}
	return utils.ParallelExecution(utils.ParallelExecutionStruct{Image: img, Function: horizontalFunc})

}

func applyVerticalBlur(img image.Image, kernel []float64) image.Image {
	bounds := img.Bounds()
	padding := len(kernel) / 2

	verticalFunc := func(start, end int, newImage *image.RGBA64, wg *sync.WaitGroup) {
		defer wg.Done()
		for y := start; y < end; y++ {
			for x := 0; x < bounds.Max.X; x++ {
				var r, g, b, a float64

				for ky := 0; ky < len(kernel); ky++ {
					sy := y + ky - padding

					if sy < 0 {
						sy = -sy
					}
					if sy >= bounds.Max.Y {
						sy = 2*bounds.Max.Y - sy - 1
					}

					sCol := img.At(x, sy)
					sr, sg, sb, sa := sCol.RGBA()

					weight := kernel[ky]
					r += float64(sr) * weight
					g += float64(sg) * weight
					b += float64(sb) * weight
					a += float64(sa) * weight
				}

				newImage.SetRGBA64(x, y, color.RGBA64{
					R: uint16(r),
					G: uint16(g),
					B: uint16(b),
					A: uint16(a),
				})
			}
		}
	}

	return utils.ParallelExecution(utils.ParallelExecutionStruct{Image: img, Function: verticalFunc})
}

// GaussianBlur applies a Gaussian blur filter to an image, creating a smooth blurring effect.
// The algorithm uses separable 1D convolutions for horizontal and vertical passes, which is
// significantly more efficient than a direct 2D convolution. Processing is done in parallel
// to maximize performance on multi-core systems.
//
// Parameters:
//   - img: The source image to be blurred (image.Image interface)
//   - level: The blur intensity, ranging from 0 (no blur) to 30 (maximum blur)
//     Higher values create a stronger blur effect but may impact performance
//
// Returns:
//   - image.Image: A new image with the blur effect applied
//
// Note: The level parameter is automatically clamped to the valid range [0, 30].
func GaussianBlur(img image.Image, level int) image.Image {
	level = utils.ClampGeneric(level, 0, 30)
	kernel := createKernel(level)
	horizontalBlur := applyHorizontalBlur(img, kernel)
	newImage := applyVerticalBlur(horizontalBlur, kernel)
	return newImage
}
