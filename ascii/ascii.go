package ascii

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

type AsciiImage struct {
	Char       string
	FontSize   int
	LineHeight int
	Color      string
}

// GenerateAscii generate an ascii art based on the input image.
//
// Parameters:
//   - img: The input image
//   - asciiChars: string of ascii characters to use
//   - newWidth: The desired width of the resized image
//   - newHeight: The desired height of the resized image
//   - useColor: if the ascii image should use color
//
// Returns:
//
//	[][]AsciiImage{
//		Char       string
//		FontSize   int
//		LineHeight int
//		Color      string
//	}
func GenerateAscii(img image.Image, asciiChars string, lineHeight, fontSize int, useColor bool) [][]AsciiImage {

	if len(asciiChars) == 0 || asciiChars == "" {
		return nil
	}

	density := []rune(asciiChars)

	newFontSize, newLineHeight := resizeAscii(lineHeight, fontSize)
	bounds := img.Bounds()

	asciiImage := make([][]AsciiImage, bounds.Max.Y)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {

		line := make([]AsciiImage, bounds.Max.X)
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			px := img.At(x, y)
			gr := color.GrayModel.Convert(px)
			gray := gr.(color.Gray)

			intensity := float64(gray.Y) / 255.0
			charIndex := math.Floor(float64(len(density)-1) * intensity)

			jsImgInfo := AsciiImage{
				Char:       string(density[int(charIndex)]),
				FontSize:   newFontSize,
				LineHeight: newLineHeight,
			}

			if useColor {
				r, g, b, _ := px.RGBA()
				colorCSS := fmt.Sprintf("rgb(%d,%d,%d)", r>>8, g>>8, b>>8)
				jsImgInfo.Color = colorCSS
			} else {
				jsImgInfo.Color = "#fff"
			}

			line[x] = jsImgInfo
		}
		asciiImage[y] = line
	}

	return asciiImage
}

func resizeAscii(lineHeight, fontSize int) (int, int) {

	imageDefaultSize := 100

	lineHeightRatio := float64(lineHeight) / float64(fontSize)
	inverseFactor := float64(imageDefaultSize) / float64(fontSize)
	newFontSize := float64(fontSize) * inverseFactor
	newLineHeight := newFontSize * lineHeightRatio

	newFontSize = math.Max(newFontSize, float64(fontSize))
	newLineHeight = math.Max(newLineHeight, float64(lineHeight))

	return int(newFontSize), int(newLineHeight)
}
