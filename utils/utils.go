package utils

import (
	"image"
	"image/color"
	"runtime"
	"sync"
)

// ParallelExecutionStruct contains parameters for parallel image processing.
//
// Fields:
//   - Image: The source image to be processed
//   - Function: The function to apply to image regions (params: start y, end y, output image, waitgroup)
//   - EndSize: Optional limit for the y-coordinate (defaults to image height if <= 0)
type ParallelExecutionStruct struct {
	Image    image.Image
	Function func(int, int, *image.RGBA64, *sync.WaitGroup)
	EndSize  int
}

// ParallelExecution processes an image in parallel using multiple CPU cores.
// It distributes the workload by dividing the image into horizontal strips
// and processing each strip concurrently with the provided function.
//
// Parameters:
//   - exec: ParallelExecutionStruct containing the image, processing function, and optional size limit
//
// Returns:
//   - image.Image: The resulting processed RGBA64 image
func ParallelExecution(exec ParallelExecutionStruct) image.Image {
	var wg sync.WaitGroup
	bounds := exec.Image.Bounds()
	newImage := image.NewRGBA64(bounds)
	cpus_available := runtime.NumCPU()
	endSize := bounds.Max.Y

	if exec.EndSize > 0 {
		endSize = exec.EndSize
	}

	if cpus_available < 4 {
		wg.Add(1)
		go exec.Function(0, endSize, newImage, &wg)

	} else {
		cpus_available--
		println("using", cpus_available, "cpus")

		wg.Add(cpus_available)
		div := (endSize + cpus_available - 1) / cpus_available
		start := 0

		for i := 0; i < cpus_available; i++ {
			end := start + div
			if end > endSize {
				end = endSize
			}
			go exec.Function(start, end, newImage, &wg)
			start = end
		}
	}

	wg.Wait()
	return newImage
}

// Luminance8bit calculates the luminance of an image (0-255) using the formula:
// L = 0.299*R + 0.587*G + 0.114*B
//
// Parameters:
//   - r,g,b: uint32 representing the red, green, and blue components of a color
//
// Returns:
//   - float64 representing the luminance of the color
func Luminance8bit(r, g, b uint32) float64 {
	return float64(r>>8)*0.299 + float64(g>>8)*0.587 + float64(b>>8)*0.114
}

// Luminance16bit calculates the luminance of an image (0-65535) using the formula:
// L = 0.299*R + 0.587*G + 0.114*B
//
// Parameters:
//   - r,g,b: uint32 representing the red, green, and blue components of a color
//
// Returns:
//   - float64 representing the luminance of the color
func Luminance16bit(r, g, b uint32) float64 {
	return float64(r)*0.299 + float64(g)*0.587 + float64(b)*0.114
}

// Clamp8bit clamp the value between 0 and 255:
//
// Parameters:
//   - value: int
//
// Returns:
//   - uint8
func Clamp8bit(value int) uint8 {
	switch {
	case value < 0:
		return 0
	case value > 255:
		return 255
	default:
		return uint8(value)
	}
}

// Clamp16bit clamp the value between 0 and 65535:
//
// Parameters:
//   - value: int32
//
// Returns:
//   - uint16
func Clamp16bit(val int32) uint16 {
	switch {
	case val < 0:
		return 0
	case val > 65535:
		return 65535
	default:
		return uint16(val)
	}
}

// ClampFloat64 clamps a float64 value between a minimum and maximum value.
//
// Parameters:
//   - value: float64 representing the value to be clamped
//   - min: float64 representing the minimum value
//   - max: float64 representing the maximum value
//
// Returns:
//   - float64 representing the clamped value
func ClampFloat64(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// ClampGeneric constrains an integer value between a minimum and maximum value.
//
// Parameters:
//   - value: int representing the value to be clamped
//   - min: int representing the minimum allowed value
//   - max: int representing the maximum allowed value
//
// Returns:
//   - int representing the clamped value
func ClampGeneric(value, min, max int) int {
	switch {
	case value < min:
		return min
	case value > max:
		return max
	default:
		return value
	}
}

// ColorAverage calculates the average color of a slice of colors:
//
// Parameters:
//   - colors: []color.Color
//
// Returns:
//   - color.Color
func ColorAverage(colors []color.Color) color.Color {

	c_l := uint32(len(colors))
	var r, g, b, a uint32

	for _, c := range colors {
		rr, gg, bb, aa := c.RGBA()
		r += rr
		g += gg
		b += bb
		a += aa
	}

	return color.NRGBA64{
		R: uint16(r / c_l),
		G: uint16(g / c_l),
		B: uint16(b / c_l),
		A: uint16(a / c_l),
	}
}
