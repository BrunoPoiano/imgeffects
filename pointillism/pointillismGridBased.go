package pointillism

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// PointillismGridBased applies a pointillism effect to an image using a grid-based approach.
//
// It divides the image into a grid. For each grid cell, it calculates the average color
// of a square region centered within that cell. It then draws a circle at the center
// of the cell using this average color. The `radius` parameter controls the size of the
// sampling region and the drawn circles, effectively determining the size of the "points".
//
// Parameters:
//   - img: The input image (image.Image) to process.
//   - radius: An integer specifying the radius of the sampling area and the resulting points (circles).
//     This value is automatically clamped between 1 and 20 (inclusive). Larger values result in larger points and a coarser effect.
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

			startX := x - edge
			startY := y - edge
			endX := startX + radius
			endY := startY + radius

			for curY := startY; curY < endY; curY++ {
				for curX := startX; curX < endX; curX++ {
					clampedX := utils.ClampGeneric(curX, bounds.Min.X, bounds.Max.X-1)
					clampedY := utils.ClampGeneric(curY, bounds.Min.Y, bounds.Max.Y-1)
					grid = append(grid, img.At(clampedX, clampedY))
				}
			}

			var c color.Color
			if len(grid) > 0 {
				c = utils.ColorAverage(grid)
			} else {
				clampedX := utils.ClampGeneric(x, bounds.Min.X, bounds.Max.X-1)
				clampedY := utils.ClampGeneric(y, bounds.Min.Y, bounds.Max.Y-1)
				c = img.At(clampedX, clampedY)
			}

			for dy := -radius; dy <= radius; dy++ {
				for dx := -radius; dx <= radius; dx++ {
					if dx*dx+dy*dy <= radius_calc {
						px := x + dx
						py := y + dy
						if px >= bounds.Min.X && px < bounds.Max.X && py >= bounds.Min.Y && py < bounds.Max.Y {
							newImage.Set(px, py, c)
						}
					}
				}
			}
		}
	}

	return newImage
}
