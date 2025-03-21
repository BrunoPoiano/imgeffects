package pointillism

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// PointillismLuminanceGridBased transforms an image to a grid of points based on luminance.
//
// Parameters:
//   - img: The input image
//   - scalling: max size of the brush: 1-100: the larger the value it takes longer to compute
//   - direction: direction of the brush: [up, down, left, right]
//
// Returns:
//   - image.Image
func PointillismLuminanceGridBased(img image.Image, scalling int, direction string) image.Image {

	scalling = utils.ClampGeneric(scalling, 1, 100)
	box := 3
	edge := box / 2

	bounds := img.Bounds()
	newImage := image.NewRGBA64(bounds)

	directionFunc := func(x, y int) {

		var grid []color.Color

		for dy := 0; dy < box; dy++ {
			for dx := 0; dx < box; dx++ {
				px := x + dx - edge
				py := y + dy - edge

				if px >= 0 && px < bounds.Max.X && py >= 0 && py < bounds.Max.Y {
					grid = append(grid, img.At(px, py))
				}
			}
		}

		averageColor := utils.ColorAverage(grid)
		r, g, b, _ := averageColor.RGBA()
		luminance := utils.Luminance16bit(r, g, b) * float64(scalling)
		radius := int(luminance / (3 * 65535) * 5)
		radius = utils.ClampGeneric(int(radius), 1, scalling)

		if radius == 1 {
			radius = rand.Intn(scalling / 2)
		}

		radius_calc := radius * radius

		for dy := -radius; dy <= radius; dy++ {
			for dx := -radius; dx <= radius; dx++ {
				if dx*dx+dy*dy <= radius_calc {
					newImage.Set(x+dx, y+dy, averageColor)
				}
			}
		}
	}

	switch direction {
	case "up":
		for y := 0; y < bounds.Max.Y; y += edge {
			for x := 0; x < bounds.Max.X; x += edge {
				directionFunc(x, y)
			}
		}

	case "down":
		for y := bounds.Max.Y; y > 0; y -= edge {
			for x := bounds.Max.X; x > 0; x -= edge {
				directionFunc(x, y)
			}
		}

	case "left":
		for x := 0; x < bounds.Max.X; x += edge {
			for y := bounds.Max.Y; y > 0; y -= edge {
				directionFunc(x, y)
			}
		}

	case "right":
		for x := bounds.Max.X; x > 0; x -= edge {
			for y := bounds.Max.Y; y > 0; y -= edge {
				directionFunc(x, y)
			}
		}
	}

	return newImage
}
