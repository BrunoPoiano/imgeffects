package lines

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// LinesHorizontal creates an image with horizontal lines effect.
// It processes the image by analyzing each horizontal row of pixels in blocks of dotSize,
// calculates the average color of each block, and draws circles with radius proportional
// to the brightness of that area.
//
// Parameters:
//   - img: The source image to apply the effect to
//   - dotSize: Size of the grid for sampling (1-20)
//   - useColor: Whether to use color from the original image or just black
//
// Returns:
//   - image.Image
func LinesHorizontal(img image.Image, dotSize int, useColor bool) image.Image {

	dotSize = utils.ClampGeneric(dotSize, 1, 20)
	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
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

// LinesVertical creates an image with vertical lines effect.
// It processes the image by analyzing each vertical column of pixels in blocks of dotSize,
// calculates the average color of each block, and draws circles with radius proportional
// to the brightness of that area.
//
// Parameters:
//   - img: The source image to apply the effect to
//   - dotSize: Size of the grid for sampling (1-20)
//   - useColor: Whether to use color from the original image or just black
//
// Returns:
//   - image.Image
func LinesVertical(img image.Image, dotSize int, useColor bool) image.Image {

	dotSize = utils.ClampGeneric(dotSize, 1, 20)
	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	for y := 0; y < bounds.Max.Y; y += dotSize {
		for x := 0; x < bounds.Max.X; x++ {

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

// LinesDiagonal creates an image with diagonal lines effect.
// It processes the image by analyzing diagonal patterns of pixels in blocks of dotSize,
// calculates the average color of each block, and draws circles with radius proportional
// to the brightness of that area.
//
// Parameters:
//   - img: The source image to apply the effect to
//   - dotSize: Size of the grid for sampling (1-20)
//   - useColor: Whether to use color from the original image or just black
//
// Returns:
//   - image.Image
func LinesDiagonal(img image.Image, dotSize int, useColor bool) image.Image {

	dotSize = utils.ClampGeneric(dotSize, 1, 20)
	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	xvalue := 0

	for y := 0; y < bounds.Max.Y; y++ {
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
