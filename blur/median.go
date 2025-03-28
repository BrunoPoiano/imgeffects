package blur

import (
	"image"
	"image/color"
	"sort"
	"sync"

	"github.com/BrunoPoiano/imgeffects/utils"
)

func sortColor(window []color.Color) color.Color {
	var r, g, b, a []uint32

	for _, c := range window {
		rr, gg, bb, aa := c.RGBA()
		r = append(r, rr)
		g = append(g, gg)
		b = append(b, bb)
		a = append(a, aa)
	}

	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	sort.Slice(g, func(i, j int) bool { return g[i] < g[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	medianIndex := len(window) / 2
	return color.NRGBA64{
		R: uint16(r[medianIndex]),
		G: uint16(g[medianIndex]),
		B: uint16(b[medianIndex]),
		A: uint16(a[medianIndex]),
	}
}

// Median apply a median blur filter to a image.
//
// Parameters:
//   - image
//   - box: 3-30: the bigger the number it takes longer to process
//
// Returns:
//   - image.Image
func Median(img image.Image, box int) image.Image {
	bounds := img.Bounds()
	box = utils.ClampGeneric(box, 3, 30)
	edgex := box / 2
	edgey := box / 2

	medianFunc := func(start, end int, newImage *image.RGBA64, wg *sync.WaitGroup) {
		defer wg.Done()
		for y := start; y < end; y++ {
			for x := 0; x < bounds.Max.X-edgex; x++ {
				var window []color.Color
				for dy := 0; dy < box; dy++ {
					for dx := 0; dx < box; dx++ {
						window = append(window, img.At(x+dx-edgex, y+dy-edgey))
					}
				}
				newImage.Set(x, y, sortColor(window))
			}
		}
	}
	return utils.ParallelExecution(utils.ParallelExecutionStruct{Image: img, Function: medianFunc})
}
