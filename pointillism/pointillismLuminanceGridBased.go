package pointillism

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// PointillismLuminanceGridBased transforms an image into a pointillism-style artwork
// using a grid-based approach where points are sized according to luminance values.
// The algorithm creates a visual effect similar to pointillist paintings, where
// discrete dots of color are applied to form a pattern that creates the overall image.
//
// Parameters:
//   - img: The input image to be transformed
//   - scalling: Controls the maximum radius of the points (1-100)
//     Higher values create larger points but increase processing time
//   - direction: Determines the traversal pattern of the algorithm:
//     "up" - processes from bottom to top, left to right
//     "down" - processes from top to bottom, right to left
//     "left" - processes from right to left, bottom to top
//     "right" - processes from left to right, bottom to top
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
