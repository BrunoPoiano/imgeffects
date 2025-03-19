package filter

import (
	"image"
	"image/color"
	"math"
	"math/rand"
)

type Point struct {
	X, Y  int
	Color color.Color
}

func distance(p1, p2 Point) float64 {
	dx := float64(p2.X - p1.X)
	dy := float64(p2.Y - p1.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

func calcPoints(img image.Image, seed int) []Point {
	bounds := img.Bounds()

	var seeds []Point
	for i := 0; i < seed; i++ {
		x := rand.Intn(bounds.Dx())
		y := rand.Intn(bounds.Dy())

		seeds = append(seeds, Point{X: x, Y: y, Color: img.At(x, y)})
	}

	return seeds
}

// VoronoiPixelation places blocks of pixels with irregular, mosaic-like shapes based on a Voronoi diagram.
//
// Parameters:
//   - img: The input image
//   - seed: the larger this number, the more seeds will be placed, resulting in a more detailed and complex pattern, but it takes longer to compute.
//
// Returns:
//   - image.Image
func VoronoiPixelation(img image.Image, seed int) image.Image {
	bounds := img.Bounds()
	newImage := image.NewRGBA64(bounds)
	seeds := calcPoints(img, seed)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {

			minDistance := math.MaxFloat64
			var nearestColor color.Color

			for _, seed := range seeds {
				distance := distance(Point{X: x, Y: y}, seed)
				if distance < minDistance {
					minDistance = distance
					nearestColor = seed.Color
				}
			}

			newImage.Set(x, y, nearestColor)
		}
	}

	return newImage

}
