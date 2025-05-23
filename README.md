# image-effects
A collection of image effects algorithms written in Go with parallel processing support.

## Installation

```bash
go get github.com/BrunoPoiano/imgeffects
```

## Available Effects

### Blur Effects
  - `blur.GaussianBlur`

  ![gaussianBlur](https://github.com/user-attachments/assets/781a5e9b-b876-4416-928f-6a71ba4f317c)

  - `blur.Median`

  ![Median](https://github.com/user-attachments/assets/a882c295-a509-4f75-ae44-03473e39efc1)


### Contrast Effects
  - `contrast.LogarithmicTransformation`

  ![LogarithmicTransformation](https://github.com/user-attachments/assets/9ff83635-d0ab-4c81-88e2-8f6a9d8bf6d3)

  - `contrast.UnsharpMasking`

  ![UnsharpMasking](https://github.com/user-attachments/assets/6f499089-ba63-44c7-beb7-515f5f5b42d7)

  - `contrast.LinearContrastStretching`

  ![LinearContrastStretching](https://github.com/user-attachments/assets/7ca5fc00-01e9-40a7-8ea3-cc120d845fd1)

  - `contrast.LinearContrastStretchingGrayscale`

  ![LinearContrastStretchingGrayscale](https://github.com/user-attachments/assets/bdcacd47-b6bb-488a-8ca9-e90f5c2cb70f)

### Dithering Effects
  - `dithering.ErrorDiffusionDithering`
    - Supported algorithms:
      - floyd-steinberg

      ![floyd-steinberg](https://github.com/user-attachments/assets/cc9ca473-3f61-4255-9338-e3bbc707b8cc)


      - false-floyd-steinberg

      ![false-floyd-steinberg](https://github.com/user-attachments/assets/c649cdbd-d07e-4082-be5d-a3bbde17b5ce)


      - jarvis-judice-ninke

      ![jarvis-judice-ninke](https://github.com/user-attachments/assets/a1b9af6c-42eb-4241-8f49-34da62bf81b9)


      - stucki

      ![stucki](https://github.com/user-attachments/assets/1fea8a7e-d010-4460-b570-109d86c75899)


      - atkinson

      ![atkinson](https://github.com/user-attachments/assets/57fcfe16-9873-4731-8001-c56445ccfba9)


      - sierra

      ![sierra](https://github.com/user-attachments/assets/6560ec2b-e0a8-4237-9187-5b25e8394b14)


      - two-row-seirra

      ![two-row-seirra](https://github.com/user-attachments/assets/cb0b91f6-7a68-4914-9496-47d5d7d422c5)


      - sierra-lite

      ![sierra-lite](https://github.com/user-attachments/assets/9cf1ee20-30bb-4621-b98d-867b795da8db)

      - none

      ![none](https://github.com/user-attachments/assets/62c81d59-7a09-4a49-a6ea-3c0339c898f0)


  - `dithering.OrderedDithering`

  ![orderedDithering](https://github.com/user-attachments/assets/a98f6d3e-ee00-435d-9b2c-956f9250e3e6)

### Flip Operations
  - `flip.FlipHorizontal`

  ![flipHorizontal](https://github.com/user-attachments/assets/fb1f5dc9-f33c-445c-9403-c0f676f894b5)

  - `flip.FlipVertical`

  ![flipVertical](https://github.com/user-attachments/assets/15ff1b8b-baa6-41cd-b976-858da0f261ab)

### HSL Adjustments
  - `hsl.Hue`

  ![hue](https://github.com/user-attachments/assets/5fa805ea-3c5c-4f73-a92b-c3e718096e9f)

  - `hsl.Saturation`

  ![saturation](https://github.com/user-attachments/assets/803800d7-fd4a-4dbc-addf-03c1874a4dfc)

  - `hsl.Luminance`

  ![luminance](https://github.com/user-attachments/assets/f225c7eb-a8b9-4600-85b1-f2eb44b240be)

### RGB
  - `rgb.adjustLevels(100, 0, 0)`

  ![AdjustLevels-red](https://github.com/user-attachments/assets/5d8e30ba-5e7e-459e-ae9c-1a6e2dff6d22)

  - `rgb.adjustLevels(0, 100, 0)`

  ![AdjustLevels-green](https://github.com/user-attachments/assets/b6cf6c47-84aa-44b7-8ca9-b7f9a8d03f81)

  - `rgb.adjustLevels(0, 0, 100)`

  ![AdjustLevels-blue](https://github.com/user-attachments/assets/80ec4ceb-35d7-46a2-a78f-137c8f168198)

### Threshold
  - `threshold.MultiThreshold`

  ![MultiThreshold](https://github.com/user-attachments/assets/77c76e1b-7bee-4b56-aea2-58ab3b5e1eb0)

  - `threshold.MultiThresholdColor`

  ![MultiThresholdColor](https://github.com/user-attachments/assets/d6342cc7-5ceb-42c6-93eb-07bbf7b79b77)

  - `threshold.GlobalThreshold`

  ![GlobalThreshold](https://github.com/user-attachments/assets/dd77ebe1-6cec-41c5-a99f-238c0c6b26ec)

  - `threshold.GlobalThresholdColor`
  
  ![GlobalThresholdColor](https://github.com/user-attachments/assets/627d12a5-4306-4940-becd-6e5748e8089d)

  - `threshold.ThresholdRGB`

  ![ThresholdRGB](https://github.com/user-attachments/assets/d0445348-8e25-4710-9a93-0fed49c68753)

  - `threshold.MultiThresholdRGB`

  ![MultiThresholdRGB](https://github.com/user-attachments/assets/347b8f97-1a7c-4f17-bb6a-7e895a75e334)

### Edge Detection
  - `edgedetection.DifferenceOfGaussians`

  ![DifferenceOfGaussians](https://github.com/user-attachments/assets/2d8b0df0-e694-4149-b24c-632b9e1358b9)

  - `edgedetection.LaplacianOfGaussian`

  ![LaplacianOfGaussian](https://github.com/user-attachments/assets/96817817-c26a-446b-abc0-baed7195962d)


  - `edgedetection.KernelOperatorBased`
    - sobel

    ![KernelOperatorBased(srcImg, sobel)](https://github.com/user-attachments/assets/95ad76b6-8f38-4cc6-9fca-b579abae8476)

    - prewitt

    ![KernelOperatorBased(srcImg, prewitt)](https://github.com/user-attachments/assets/dbf44385-aa29-4941-b302-4155135532ac)

    - robert-cross

    ![KernelOperatorBased(srcImg, robert-cross)](https://github.com/user-attachments/assets/92dafeda-c93c-4fd2-9767-b1f008148e60)

    - scharr

    ![KernelOperatorBased(srcImg, scharr)](https://github.com/user-attachments/assets/f642f84d-7b1f-48da-86ac-085b5814495d)

### Pointillism
  - `pointillism.halftone - black and white`
  
  ![Halftone_black](https://github.com/user-attachments/assets/15f1fd19-4239-4b38-b715-841ba6173310)

  - `pointillism.halftone - Color`

  ![Halftone_color](https://github.com/user-attachments/assets/4cef79c3-3026-415b-87a1-a2779966c1bd)

  - `pointillism.PointillismGridBased`

  ![PointillismGridBased](https://github.com/user-attachments/assets/bec98442-c11d-49e9-80d7-0fe102015289)

  - `pointillism.PointillismLuminanceBased`

  ![PointillismLuminanceBased](https://github.com/user-attachments/assets/e959b2ee-75c2-41fe-9366-99ff32498437)

  - `pointillism.PointillismLuminanceGridBased`
    - up

     ![PointillismLuminanceGridBased-up](https://github.com/user-attachments/assets/a71cc2be-4b03-4f56-b888-29d2cb1eb054)

    - down

     ![PointillismLuminanceGridBased-down](https://github.com/user-attachments/assets/996e6e04-cb7d-4155-865b-66ecfc792de1)

    - left

    ![PointillismLuminanceGridBased-left](https://github.com/user-attachments/assets/7c4b0cd4-c4fd-4e42-bc63-347d7cf26dc7)

    - right

    ![PointillismLuminanceGridBased-right](https://github.com/user-attachments/assets/4a31d61a-a90d-43d5-89fd-63980325450b)

### Filters

  - `filter.VoronoiPixelation`

  ![VoronoiPixelation(](https://github.com/user-attachments/assets/08db4b9b-e853-4f58-839e-286119f97839)

  - `filter.SolarizeEffect`

  ![SolarizeEffect](https://github.com/user-attachments/assets/b03ae250-8d78-474f-ba0c-d7d3a49486d2)

  - `filter.ChromaticAberration`

  ![ChromaticAberration](https://github.com/user-attachments/assets/be0a3c01-50a5-40dd-afe7-90580afb68bd)

  - `filter.GammaCorrection`

  ![GammaCorrection](https://github.com/user-attachments/assets/cec2d686-05d1-41f5-baaf-e5da5199831d)

  - `filter.KuwaharaFilter`

  ![kawahara](https://github.com/user-attachments/assets/23329558-ab98-4998-8c60-a37ef0a3251c)

  - `filter.GrayScale16`

  ![GrayScale16](https://github.com/user-attachments/assets/9dfa73fc-5921-4465-930f-65b96060fdc5)

  - `filter.GrayScale`

  ![GrayScale](https://github.com/user-attachments/assets/686ccce1-565d-468f-a4c7-910be01b119d)

### Lines

  - `LinesHorizontal - Color`
  
  ![linesHorizontalColor](https://github.com/user-attachments/assets/32eb7d6a-bddf-481b-a02d-0d37595373e1)

  - `LinesHorizontal`
  
  ![linesHorizontalblack](https://github.com/user-attachments/assets/51034756-f19d-493c-93bc-13e9a01eca17)
    
  - `LinesVertical Color`

  ![linesVerticalcolor](https://github.com/user-attachments/assets/204c8e3f-7047-4429-a195-ac4151bccea8)

  - `LinesVertical`

  ![linesVerticalblack](https://github.com/user-attachments/assets/10eac720-2ef8-4821-9c09-4f85ed94063b)

## Noise

  - `noise.NoiseGenerator`

  ![noiseGenerator](https://github.com/user-attachments/assets/86402539-88ad-4182-a0b9-eee4a05101ba)

  - `noise.NoiseGeneratorColor`

  ![noiseGeneratorColor](https://github.com/user-attachments/assets/3a5d5903-f175-46f0-a759-2ceb6955bd57)

  - `noise.NoiseGeneratorGrayScale`

  ![noiseGeneratorGrayScale](https://github.com/user-attachments/assets/08014298-b136-456a-9904-7090af556163)

  - `noise.BlendingNoiseToImage`
    - gray

    ![mergeTwoImages-gray](https://github.com/user-attachments/assets/1fb06417-04ee-4ebc-b2a9-03c30122491e)

    - color

    ![mergeTwoImages-color](https://github.com/user-attachments/assets/5e54e4d8-7eab-43d6-81d3-ba12ce186aae)

    - default

    ![mergeTwoImages-default](https://github.com/user-attachments/assets/a7afbf15-f227-44bb-b28b-ef694d456cd1)

## Resize
  - `resize.NearestNeighbor`

  ![NearestNeighbor](https://github.com/user-attachments/assets/a91cc3d9-4009-4020-a7ec-827923312d89)

  - `resize.BypolarInterpolate`

  ![BypolarInterpolate](https://github.com/user-attachments/assets/488966b9-acfd-44d0-a822-70bb16eaf6a6)

## Ascii
  - `ascii.GenerateAscii`

## HelperFunctions
  - `hsl.HSLToRGB`
  - `hsl.RGBToHSL`
  - `kuwahara.RGBToHSV`
  - `resize.NewAspectRatio`

## License
MIT License - See LICENSE file for details.
