package filter

import (
	"image"
	"image/color"
	"math"
	"math/rand"
	"sync"

	"github.com/BrunoPoiano/imgeffects/utils"
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

// VoronoiPixelation applies a Voronoi diagram-based pixelation effect to an image.
// This creates an abstract, mosaic-like pattern where the image is divided into
// irregular cell regions, each filled with a single color sampled from the original image.
//
// The algorithm works by randomly placing seed points throughout the image and
// assigning each pixel to the closest seed point, inheriting its color.
// This creates organic, cell-like regions that form a distinctive artistic effect.
//
// Parameters:
//   - img: The input image to be processed
//   - seed: Number of seed points to generate. Higher values (10-1000) create more
//     detailed patterns with smaller cells, while lower values result in larger,
//     more abstract shapes. Values are clamped between 1 and 100000.
//     Increasing this value significantly will increase processing time.
//
// Returns:
//   - image.Image: The processed image with the Voronoi pixelation effect applied
//
// The function uses parallel processing for improved performance on multi-core systems.
func VoronoiPixelation(img image.Image, seed int) image.Image {
	bounds := img.Bounds()
	seeds := calcPoints(img, seed)
	seed = utils.ClampGeneric(seed, 1, 100000)

	voronoiPixelationFunction := func(start, end int, newImage *image.RGBA64, wg *sync.WaitGroup) {
		defer wg.Done()
		for y := start; y < end; y++ {
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
	}

	return utils.ParallelExecution(utils.ParallelExecutionStruct{Image: img, Function: voronoiPixelationFunction})
}
