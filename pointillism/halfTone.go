package pointillism

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// Halftone applies a halftone effect to an image.
// It simulates the halftone printing technique by creating a pattern of dots that vary in size based on the luminance of the source image.
//
// Parameters:
//   - img: The input image
//   - dotSize: The size of each halftone dot cell in pixels: 1-20
//   - useColor: When true, uses the average color of each cell for the dots;
//     when false, uses solid black dots
//
// Returns:
//   - image.Image
func Halftone(img image.Image, dotSize int, useColor bool) image.Image {

	dotSize = utils.ClampGeneric(dotSize, 1, 20)
	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	for y := 0; y < bounds.Max.Y; y += dotSize {
		for x := 0; x < bounds.Max.X; x += dotSize {

			var grid []color.Color

			for dy := 0; dy < dotSize; dy++ {
				for dx := 0; dx < dotSize; dx++ {
					px := x + dx
					py := y + dy

					if px < 0 || py < 0 || px >= bounds.Max.X || py >= bounds.Max.Y {
						continue
					}
					grid = append(grid, img.At(px, py))
				}
			}

			colorAverage := utils.ColorAverage(grid)
			r, g, b, _ := colorAverage.RGBA()
			gray := utils.Luminance16bit(r, g, b)

			radius := int((1 - gray/65535) * float64(dotSize) / 2)
			radiusCalc := radius * radius

			centerX := x + dotSize/2
			centerY := y + dotSize/2

			for dy := -radius; dy <= radius; dy++ {
				for dx := -radius; dx <= radius; dx++ {
					if dx*dx+dy*dy <= radiusCalc {
						px := centerX + dx
						py := centerY + dy
						if useColor {
							newImage.Set(px, py, colorAverage)
						} else {
							newImage.Set(px, py, color.Black)
						}
					}
				}
			}
		}
	}
	return newImage
}

// HalftoneDiagonal applies a diagonal halftone effect to an image.
// It creates a pattern of dots with a diagonal offset, producing a unique halftone effect
// where each row of dots is slightly offset from the previous one.
//
// Parameters:
//   - img: The input image
//   - dotSize: The size of each halftone dot cell in pixels: 1-20
//   - useColor: When true, uses the average color of each cell for the dots;
//     when false, uses solid black dots
//
// Returns:
//   - image.Image
func HalftoneDiagonal(img image.Image, dotSize int, useColor bool) image.Image {

	dotSize = utils.ClampGeneric(dotSize, 1, 20)
	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	xvalue := 0
	for y := 0; y < bounds.Max.Y; y += dotSize {
		for x := xvalue; x < bounds.Max.X; x += dotSize {

			var grid []color.Color

			for dy := 0; dy < dotSize; dy++ {
				for dx := 0; dx < dotSize; dx++ {
					px := x + dx
					py := y + dy

					if px < 0 || py < 0 || px >= bounds.Max.X || py >= bounds.Max.Y {
						continue
					}
					grid = append(grid, img.At(px, py))
				}
			}

			colorAverage := utils.ColorAverage(grid)
			r, g, b, _ := colorAverage.RGBA()
			gray := utils.Luminance16bit(r, g, b)

			radius := int((1 - gray/65535) * float64(dotSize) / 2)
			radiusCalc := radius * radius

			centerX := x + dotSize/2
			centerY := y + dotSize/2

			for dy := -radius; dy <= radius; dy++ {
				for dx := -radius; dx <= radius; dx++ {
					if dx*dx+dy*dy <= radiusCalc {
						px := centerX + dx
						py := centerY + dy
						if useColor {
							newImage.Set(px, py, colorAverage)
						} else {
							newImage.Set(px, py, color.Black)
						}
					}
				}

			}
		}

		if xvalue >= dotSize {
			xvalue = 0
		} else {
			xvalue++
		}

	}
	return newImage
}
