package pointillism

import (
	"image"
	"math/rand"
	"sync"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// PointillismLuminanceBased creates a pointillism effect on the input image.
// It achieves this by drawing a specified number of circular points ('dots')
// at random locations on a new canvas. The color of each dot is sampled
// from the corresponding pixel in the original image. The size (radius) of
// each dot is proportional to the luminance of the sampled pixel, adjusted
// by the 'scaling' factor, and constrained to a minimum and maximum size.
// The effect is generated using parallel processing for efficiency.
//
// Parameters:
//   - img: The source image (image.Image) to apply the effect to.
//   - points: The total number of circular points to draw. This value is clamped
//     internally to a minimum of 10. Higher values result in a denser effect.
//   - scaling: Controls the maximum size of the points and acts as a scaling factor
//     in the luminance-to-radius calculation. Effectively sets the upper limit
//     for the point radius (clamped between 1 and 30).
//
// Returns:
//   - image.Image.
func PointillismLuminanceBased(img image.Image, points, scaling int) image.Image {
	bounds := img.Bounds()
	points = utils.ClampGeneric(points, 10, 99999999999999)
	scaling = utils.ClampGeneric(scaling, 1, 30)

	pointFunc := func(start, end int, newImage *image.RGBA64, wg *sync.WaitGroup) {
		defer wg.Done()
		for i := start; i < end; i++ {
			x := rand.Intn(bounds.Dx())
			y := rand.Intn(bounds.Dy())
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()

			lum := utils.Luminance16bit(r, g, b) * float64(scaling)

			radius := int(lum / (3 * 65535) * 5)
			radius = utils.ClampGeneric(radius, 1, scaling)
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
