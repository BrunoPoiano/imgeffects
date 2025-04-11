package hsl

import (
	"image"
	"image/color"
	"math"

	"github.com/BrunoPoiano/imgeffects/utils"
)

// HSLToRGB converts HSL (Hue, Saturation, Luminance) values to RGB color space.
// Hue is specified in degrees (0-360), while saturation and luminance are in the range 0.0 to 1.0.
// Returns 16-bit per channel RGB values as uint32 (range 0-65535).
//
// Parameters:
//   - h: Hue angle in degrees (0-360)
//   - s: Saturation value (0.0-1.0)
//   - l: Luminance value (0.0-1.0)
//
// Returns:
//   - r, g, b: RGB color channels as uint32 values (0-65535)
func HSLToRGB(h, s, l float64) (r, g, b uint32) {
	var fr, fg, fb float64

	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - c/2

	switch {
	case h < 60:
		fr, fg, fb = c, x, 0
	case h < 120:
		fr, fg, fb = x, c, 0
	case h < 180:
		fr, fg, fb = 0, c, x
	case h < 240:
		fr, fg, fb = 0, x, c
	case h < 300:
		fr, fg, fb = x, 0, c
	default:
		fr, fg, fb = c, 0, x
	}

	r = uint32((fr + m) * 65535)
	g = uint32((fg + m) * 65535)
	b = uint32((fb + m) * 65535)

	return r, g, b
}

// RGBToHSL converts RGB color values to HSL (Hue, Saturation, Luminance) color space.
// Takes 16-bit per channel RGB values and returns HSL components.
//
// Parameters:
//   - r, g, b: RGB color channels as uint32 values (0-65535)
//
// Returns:
//   - h: Hue angle in degrees (0-360)
//   - s: Saturation value (0.0-1.0)
//   - l: Luminance value (0.0-1.0)
func RGBToHSL(r, g, b uint32) (h, s, l float64) {
	fr := float64(r) / 65535.0
	fg := float64(g) / 65535.0
	fb := float64(b) / 65535.0

	max := math.Max(math.Max(fr, fg), fb)
	min := math.Min(math.Min(fr, fg), fb)
	l = (max + min) / 2.0

	if max == min {
		h, s = 0, 0 // Achromatic
	} else {
		delta := max - min
		s = delta / (1 - math.Abs(2*l-1))

		switch max {
		case fr:
			h = math.Mod((fg-fb)/delta, 6)
		case fg:
			h = (fb-fr)/delta + 2
		case fb:
			h = (fr-fg)/delta + 4
		}

		h *= 60
		if h < 0 {
			h += 360
		}
	}

	return h, s, l
}

// Hue adjusts the hue component of an image while preserving saturation and luminance.
// This function shifts the hue angle of each pixel by the specified amount.
//
// Parameters:
//   - img: The input image to be processed
//   - change: Hue shift in degrees (0-360), will be clamped to valid range
//
// Returns:
//   - image.Image
func Hue(img image.Image, change int) image.Image {
	change = utils.ClampGeneric(change, 0, 360)
	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			h, s, l := RGBToHSL(r, g, b)
			h = float64((int(h) + change) % 360)
			rr, gg, bb := HSLToRGB(h, s, l)

			newImage.Set(x, y, color.NRGBA64{uint16(rr), uint16(gg), uint16(bb), uint16(a)})
		}
	}

	return newImage

}

// Saturation adjusts the color saturation of an image while preserving hue and luminance.
// Positive values increase saturation while negative values decrease it.
//
// Parameters:
//   - img: The input image
//   - change: Saturation adjustment (-1.0 to 1.0)
//
// Returns:
//   - image.Image
func Saturation(img image.Image, change float64) image.Image {
	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)
	change = math.Max(-1, math.Min(1, change))

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			h, s, l := RGBToHSL(r, g, b)
			s = math.Max(0, math.Min(1, s*(1+change)))
			rr, gg, bb := HSLToRGB(h, s, l)

			newImage.Set(x, y, color.NRGBA64{uint16(rr), uint16(gg), uint16(bb), uint16(a)})
		}
	}

	return newImage

}

// Luminance adjusts the brightness or luminance of an image while preserving hue and saturation.
// Positive values brighten the image, while negative values darken it.
//
// Parameters:
//   - img: The input image to be processed
//   - change: Luminance adjustment factor (-1.0 to 1.0)
//
// Returns:
//   - image.Image
func Luminance(img image.Image, change float64) image.Image {

	bounds := img.Bounds()
	newImage := image.NewNRGBA64(bounds)
	change = math.Max(-1, math.Min(1, change))

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			h, s, l := RGBToHSL(r, g, b)
			l = math.Max(0, math.Min(1, l*(1+change)))
			rr, gg, bb := HSLToRGB(h, s, l)

			newImage.Set(x, y, color.NRGBA64{uint16(rr), uint16(gg), uint16(bb), uint16(a)})
		}
	}

	return newImage

}
