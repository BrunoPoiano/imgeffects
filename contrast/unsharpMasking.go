package contrast

import (
	"image"
	"image/color"

	"github.com/BrunoPoiano/imgeffects/blur"
	"github.com/BrunoPoiano/imgeffects/utils"
)

// UnsharpMasking applies the Unsharp Masking algorithm to sharpen an image by subtracting
// a blurred version from the original, enhancing fine details and edge definition.
//
// Unsharp masking works by creating a blurred copy of the image, then amplifying the
// difference between the original and blurred version to enhance image details.
//
// Parameters:
//   - img: The input image to be sharpened
//   - variation: A value between 0.0 and 1.0 that controls sharpening intensity.
//     At 0.0, no sharpening occurs, while 1.0 applies maximum sharpening effect.
//   - blurLevel: An integer from 1 to 20 (clamped internally) that determines the
//     radius of the Gaussian blur applied. Lower values produce finer detail enhancement,
//     while higher values affect larger features.
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
