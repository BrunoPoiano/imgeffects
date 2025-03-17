# image-effects
A collection of image effects algorithms written in Go.

## Installation

```bash
go get github.com/BrunoPoiano/imgeffects
```

## Available Effects

### Blur Effects
  - `GaussianBlur`
  ![gaussianBlur](https://github.com/user-attachments/assets/781a5e9b-b876-4416-928f-6a71ba4f317c)

  - `Median`

### Dithering Effects
  - `ErrorDiffusionDithering`
    - Supported algorithms:
      - floyd-steinberg
      - false-floyd-steinberg
      - jarvis-judice-ninke
      - stucki
      - atkinson
      - sierra
      - two-row-seirra
      - sierra-lite

  - `OrderedDithering`
  ![orderedDithering](https://github.com/user-attachments/assets/a98f6d3e-ee00-435d-9b2c-956f9250e3e6)

### Flip Operations
  - `FlipHorizontal`
  ![flipHorizontal](https://github.com/user-attachments/assets/fb1f5dc9-f33c-445c-9403-c0f676f894b5)

  - `FlipVertical`
  ![flipVertical](https://github.com/user-attachments/assets/15ff1b8b-baa6-41cd-b976-858da0f261ab)

### HSL Adjustments
  - `Hue` - 0-360 degrees
  ![hue](https://github.com/user-attachments/assets/5fa805ea-3c5c-4f73-a92b-c3e718096e9f)

  - `Saturation` - -1.0 to 1.0
  ![saturation](https://github.com/user-attachments/assets/803800d7-fd4a-4dbc-addf-03c1874a4dfc)

  - `Luminance` - -1.0 to 1.0
  ![luminance](https://github.com/user-attachments/assets/f225c7eb-a8b9-4600-85b1-f2eb44b240be)

### Filters
  - `KuwaharaFilter` - 1-20
  ![kawahara](https://github.com/user-attachments/assets/23329558-ab98-4998-8c60-a37ef0a3251c)

### Filters
  - `GrayScale16`
  - `GrayScale`

## Ascii
  - `GenerateAscii`

## HelperFunctions
  - `hsl.HSLToRGB`
  - `hsl.RGBToHSL`
  - `kuwahara.RGBToHSV`
  - `resize.NewAspectRatio`

## License
MIT License - See LICENSE file for details.
