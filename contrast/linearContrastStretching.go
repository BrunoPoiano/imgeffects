package contrast

import (
	"image"
	"image/color"
)

func LinearContrastStretchingGrayscale(img image.Image) image.Image {

	bounds := img.Bounds()
	newImage := image.NewRGBA(bounds)

	//intensity
	I_min, I_max := uint8(255), uint8(0)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := uint8((r + g + b) / 3 >> 8)
			if uint8(gray) < I_min {
				I_min = uint8(gray)
			}

			if uint8(gray) > I_max {
				I_max = uint8(gray)
			}
		}
	}

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := uint8((r + g + b) / 3 >> 8)
			I_out := ((I_max - I_min) / 255) * (gray - I_min)

			newImage.Set(x, y, color.Gray{I_out})
		}
	}

	return newImage
}

func LinearContrastStretching(img image.Image) image.Image {

	bounds := img.Bounds()
	newImage := image.NewRGBA(bounds)

	//intensity
	R_min, R_max := uint8(255), uint8(0)
	G_min, G_max := uint8(255), uint8(0)
	B_min, B_max := uint8(255), uint8(0)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			rr := uint8(r >> 8)
			gg := uint8(g >> 8)
			bb := uint8(b >> 8)

			if rr < R_min {
				R_min = rr
			}
			if rr > R_max {
				R_max = rr
			}

			if gg < G_min {
				G_min = gg
			}
			if gg > G_max {
				G_max = gg
			}

			if bb < B_min {
				B_min = bb
			}
			if bb > B_max {
				B_max = bb
			}
		}
	}

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			rr := uint8(r >> 8)
			gg := uint8(g >> 8)
			bb := uint8(b >> 8)
			aa := uint8(a >> 8)

			stretch := func(value, min, max uint8) uint8 {

				if max == min {
					return value
				}
				return uint8(255 * (float64(value-min) / float64(max-min)))
			}

			newImage.Set(x, y, color.RGBA{
				stretch(rr, R_min, R_max),
				stretch(gg, G_min, G_max),
				stretch(bb, B_min, B_max),
				aa,
			})
		}
	}

	return newImage
}
