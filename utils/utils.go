package utils

import (
	"image"
	"runtime"
	"sync"
)

type ParallelExecutionStruct struct {
	Image    image.Image
	Function func(int, int, *image.RGBA64, *sync.WaitGroup)
}

func ParallelExecution(exec ParallelExecutionStruct) image.Image {
	var wg sync.WaitGroup

	bounds := exec.Image.Bounds()
	newImage := image.NewRGBA64(bounds)
	cpus_available := runtime.NumCPU()

	if cpus_available < 4 {
		wg.Add(1)
		go exec.Function(0, bounds.Max.Y, newImage, &wg)

	} else {
		cpus_available--
		println("using", cpus_available, "cpus")

		wg.Add(cpus_available)
		div := (bounds.Max.Y + cpus_available - 1) / cpus_available
		start := 0

		for i := 0; i < cpus_available; i++ {
			end := start + div
			if end > bounds.Max.Y {
				end = bounds.Max.Y
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
