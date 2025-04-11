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

// GenerateAscii converts an input image to ASCII art representation.
//
// The function maps pixel brightness values to characters from the provided
// ASCII character set, where darker pixels use characters from the beginning
// of the string and brighter pixels use characters from the end.
//
// Parameters:
//   - img: The source image to convert to ASCII art
//   - asciiChars: String of ASCII characters to use (ordered from darkest to lightest)
//   - lineHeight: The vertical spacing between lines in the output
//   - fontSize: The font size to use for ASCII characters
//   - useColor: Whether to preserve the original image colors (true) or use white text (false)
//
// Returns:
//   - A 2D slice of AsciiImage structs representing the converted image, where each element
//     contains character, styling, and color information for a single position.
//     Returns nil if asciiChars is empty.
//
// Each AsciiImage in the result contains:
//   - Char: The ASCII character for this position
//   - FontSize: The calculated font size
//   - LineHeight: The calculated line height
//   - Color: CSS-compatible color string (RGB format or "#fff")
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
