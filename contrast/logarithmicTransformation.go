package contrast

import (
	"image"
	"image/color"
	"math"
	"runtime"
	"sync"
)

// LogarithmicTransformation applies logarithmic transformation to an image.
//
// Parameters:
//   - img: The input image
//   - variation: [-1,1]
//
// Returns:
//   - image.Image
func LogarithmicTransformation(img image.Image, variation float64) image.Image {

	bounds := img.Bounds()
	newImage := image.NewRGBA64(bounds)

	calc := func(value uint32) uint16 {
		switch variation {
		case 0:
			variation = 0.01
		case -1:
			variation = -0.99
		}
		v := float64(value) / 65535.0
		transformed := math.Log(1+variation*v) / math.Log(1+variation)
		return uint16(transformed * 65535.0)
	}

	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for y := 0; y < bounds.Max.Y/2; y++ {
			for x := 0; x < bounds.Max.X; x++ {
				r, g, b, a := img.At(x, y).RGBA()

				newImage.Set(x, y, color.RGBA64{
					calc(r),
					calc(g),
					calc(b),
					uint16(a),
				})
			}
		}
	}()
	go func() {
		defer wg.Done()
		for y := bounds.Max.Y / 2; y < bounds.Max.Y; y++ {
			for x := 0; x < bounds.Max.X; x++ {
				r, g, b, a := img.At(x, y).RGBA()

				newImage.Set(x, y, color.RGBA64{
					stretch(r),
					stretch(g),
					stretch(b),
					uint16(a),
				})
			}
		}
	}()
	wg.Wait()
	return newImage
}
