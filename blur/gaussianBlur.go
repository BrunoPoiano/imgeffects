package blur

import (
	"image"
	"image/color"
	"math"
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

func applyHorizontalBlur(img image.Image, kernel []float64) *image.RGBA64 {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	result := image.NewRGBA64(bounds)

	padding := len(kernel) / 2

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var r, g, b, a float64

			for kx := 0; kx < len(kernel); kx++ {
				sx := x + kx - padding

				if sx < 0 {
					sx = -sx
				}
				if sx >= width {
					sx = 2*width - sx - 1
				}

				sCol := img.At(sx, y)
				sr, sg, sb, sa := sCol.RGBA()

				weight := kernel[kx]
				r += float64(sr) * weight
				g += float64(sg) * weight
				b += float64(sb) * weight
				a += float64(sa) * weight
			}

			result.SetRGBA64(x, y, color.RGBA64{
				R: uint16(r),
				G: uint16(g),
				B: uint16(b),
				A: uint16(a),
			})
		}
	}

	return result
}

func applyVerticalBlur(img image.Image, kernel []float64) *image.RGBA64 {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	result := image.NewRGBA64(bounds)

	padding := len(kernel) / 2

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var r, g, b, a float64

			for ky := 0; ky < len(kernel); ky++ {
				sy := y + ky - padding

				if sy < 0 {
					sy = -sy
				}
				if sy >= height {
					sy = 2*height - sy - 1
				}

				sCol := img.At(x, sy)
				sr, sg, sb, sa := sCol.RGBA()

				weight := kernel[ky]
				r += float64(sr) * weight
				g += float64(sg) * weight
				b += float64(sb) * weight
				a += float64(sa) * weight
			}

			result.SetRGBA64(x, y, color.RGBA64{
				R: uint16(r),
				G: uint16(g),
				B: uint16(b),
				A: uint16(a),
			})
		}
	}

	return result
}

// GaussianBlur apply a Gaussian blur filter to a image.
//
// Parameters:
//   - image
//   - size 0 to 20
//
// Returns:
//   - image.Image
func GaussianBlur(img image.Image, size int) image.Image {

	kernel := createKernel(size)
	horizontalBlur := applyHorizontalBlur(img, kernel)
	newImage := applyVerticalBlur(horizontalBlur, kernel)
	return newImage

}
