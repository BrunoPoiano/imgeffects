package contrast

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/blur"
	"github.com/BrunoPoiano/imgeffects/utils"
)

// UnsharpMasking applies the Unsharp Masking algorithm filter to an image.
//
// Parameters:
//   - img: The input image
//   - variation: [0,1]: controls sharpening strength, where 0 is no effect and 1 is maximum
//   - blurLevel: [1,10]: controls the amount of blur applied before applying the mask
//
// Returns:
//   - image.Image
func UnsharpMasking(img image.Image, variation float64, blurLevel int) image.Image {
	blurLevel = utils.ClampGeneric(blurLevel, 1, 20)
	bounds := img.Bounds()
	newImage := image.NewRGBA64(bounds)
	bluredImage := blur.GaussianBlur(img, blurLevel)

	stretch := func(value, bluredValue uint32) uint16 {

		diff := int32(value) - int32(bluredValue)
		result := int32(value) + int32(float64(diff)*variation)
		return utils.Clamp16bit(result)
	}

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			br, bg, bb, _ := bluredImage.At(x, y).RGBA()

			newImage.Set(x, y, color.RGBA64{
				stretch(r, br),
				stretch(g, bg),
				stretch(b, bb),
				uint16(a),
			})
		}
	}

	return newImage
}
