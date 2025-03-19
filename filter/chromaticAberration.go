package filter

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// Chromatic Aberration Mimics the lens distortion where red, green, and blue channels misalign.
//
// Parameters:
//   - img: The input image
//   - x_offset: offset for the x axis: 1-10
//   - y_offset: offset for the y axis: 1-10
//
// Returns:
//   - image.Image
func ChromaticAberration(img image.Image, x_offset, y_offset int) image.Image {
	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	x_offset = utils.ClampGeneric(x_offset, 1, 20)
	y_offset = utils.ClampGeneric(y_offset, 1, 20)

	blue_y_offset := y_offset * -1
	blue_x_offset := x_offset * -1

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {

			r, _, _, a := img.At(x+x_offset, y+y_offset).RGBA()
			_, g, _, _ := img.At(x, y).RGBA()
			_, _, b, _ := img.At(x+blue_x_offset, y+blue_y_offset).RGBA()

			newImage.Set(x, y, color.NRGBA64{
				uint16(r),
				uint16(g),
				uint16(b),
				uint16(a),
			})
		}
	}

	return newImage

}
