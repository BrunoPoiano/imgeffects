package pointillism

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// PointillismGridBased transforms an image into a grid of points.
//
// Parameters:
//   - img: The input image
//   - radius: size of the point: 1-20
//
// Returns:
//   - image.Image
func PointillismGridBased(img image.Image, radius int) image.Image {
	radius = utils.ClampGeneric(radius, 1, 20)
	bounds := img.Bounds()
	newImage := image.NewRGBA64(bounds)

	edge := radius / 2
	step := radius * 2
	radius_calc := radius * radius

	for y := 0; y < bounds.Max.Y; y += step {
		for x := 0; x < bounds.Max.X; x += step {

			var grid []color.Color

			for dy := 0; dy < radius; dy++ {
				for dx := 0; dx < radius; dx++ {
					px := x + dx - edge
					py := y + dy - edge

					grid = append(grid, img.At(px, py))
				}
			}

			c := utils.ColorAverage(grid)

			for dy := -radius; dy <= radius; dy++ {
				for dx := -radius; dx <= radius; dx++ {
					if dx*dx+dy*dy <= radius_calc {
						px := x + dx
						py := y + dy
						newImage.Set(px, py, c)
					}
				}
			}
		}
	}

	return newImage
}
