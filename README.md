# image-effects
A collection of image effects algorithms written in Go.

## Installation

```bash
go get github.com/BrunoPoiano/imgeffects
```

## Available Effects

### Blur Effects
- `GaussianBlur(img image.Image, size int) image.Image`

### Dithering Effects
- `ErrorDiffusionDithering(img image.Image, algorithm string, level int) image.Image`
  - Supported algorithms: "floyd-steinberg", "false-floyd-steinberg", "jarvis-judice-ninke", "stucki", "atkinson", "sierra", "two-row-seirra", "sierra-lite"
- `OrderedDithering(img image.Image, level, size int) image.Image`

### Flip Operations
- `FlipHorizontal(img image.Image) image.Image`
- `FlipVertical(img image.Image) image.Image`

### HSL Adjustments
- `AdjustHue(img image.Image, change int) image.Image` - change range: 0-360 degrees
- `AdjustSaturation(img image.Image, change float64) image.Image` - change range: -1.0 to 1.0
- `AdjustLuminance(img image.Image, change float64) image.Image` - change range: -1.0 to 1.0

### Filters
- `KuwaharaFilter(img image.Image, size int) image.Image` - size range: 1-20

## License

MIT License - See LICENSE file for details.
