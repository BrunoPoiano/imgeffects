package pointillism

import (
	"image"
	"math/rand"
	"sync"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// PointillismLuminanceBased applies points randomly based on the luminance of each pixel.
//
// Parameters:
//   - img: The input image
//   - points: the larger this number, more points the image gets
//   - scalling: max size of the brush: 1-30
//
// Returns:
//   - image.Image
func PointillismLuminanceBased(img image.Image, points, scalling int) image.Image {
	bounds := img.Bounds()
	points = utils.ClampGeneric(points, 10, 99999999999999)

	pointFunc := func(start, end int, newImage *image.RGBA64, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := start; i < end; i++ {
			x := rand.Intn(bounds.Dx())
			y := rand.Intn(bounds.Dy())
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()

			lum := utils.Luminance16bit(r, g, b) * float64(scalling)

			radius := int(lum / (3 * 65535) * 5)
			radius = utils.ClampGeneric(radius, 1, scalling)
			if radius == 1 {
				radius = rand.Intn(5)
			}
			radius_calc := radius * radius

			for dy := -radius; dy <= radius; dy++ {
				for dx := -radius; dx <= radius; dx++ {
					if dx*dx+dy*dy <= radius_calc {
						newImage.Set(x+dx, y+dy, color)
					}
				}
			}
		}
	}

	return utils.ParallelExecution(utils.ParallelExecutionStruct{Image: img, Function: pointFunc, EndSize: points})
}
