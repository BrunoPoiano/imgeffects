package contrast

import (
	"image"
	"image/color"
	"runtime"
	"sync"

	"github.com/BrunoPoiano/imgeffects/blur"
)

// UnsharpMasking applies the Unsharp Masking algorithm filter to an image.
//
// Parameters:
//   - img: The input image
//   - variation: [0,1]
//   - blurLevel: [1,10]
//
// Returns:
//   - image.Image
func UnsharpMasking(img image.Image, variation float64, blurLevel int) image.Image {

	bounds := img.Bounds()
	newImage := image.NewRGBA64(bounds)
	bluredImage := blur.GaussianBlur(img, blurLevel)

	stretch := func(value, bluredValue uint32) uint16 {

		clamp := func(val int32) uint16 {
			switch {
			case val < 0:
				return 0
			case val > 65535:
				return 65535
			default:
				return uint16(val)
			}
		}

		diff := int32(value) - int32(bluredValue)
		result := int32(value) + int32(float64(diff)*variation)
		return clamp(result)
	}

	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for y := 0; y < bounds.Max.Y/2; y++ {
			for x := 0; x < bounds.Max.X; x++ {
				//normal Image
				r, g, b, a := img.At(x, y).RGBA()
				// blured Image
				br, bg, bb, _ := bluredImage.At(x, y).RGBA()

				newImage.Set(x, y, color.RGBA64{
					stretch(r, br),
					stretch(g, bg),
					stretch(b, bb),
					uint16(a),
				})
			}
		}
	}()
	go func() {
		defer wg.Done()
		for y := bounds.Max.Y / 2; y < bounds.Max.Y; y++ {
			for x := 0; x < bounds.Max.X; x++ {
				//normal Image
				r, g, b, a := img.At(x, y).RGBA()
				// blured Image
				br, bg, bb, _ := bluredImage.At(x, y).RGBA()

				newImage.Set(x, y, color.RGBA64{
					stretch(r, br),
					stretch(g, bg),
					stretch(b, bb),
					uint16(a),
				})
			}
		}
	}()
	wg.Wait()
	return newImage
}
