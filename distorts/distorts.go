package distorts

import (
	"image"
	"image/color"
	"math/rand"
	"time"
)

type pixelColor struct {
	r, g, b, a uint32
}

func (c pixelColor) RGBA() (uint32, uint32, uint32, uint32) {
	return c.r, c.g, c.b, c.a
}

type offset struct {
  X, Y int
}

//ChromaticAberation returns a copy of input with a random chromatic aberation effect applied
func ChromaticAberation(input image.Image) image.Image {
	bounds := input.Bounds()
	newImg := image.NewRGBA(bounds)

	rand.Seed(time.Now().UTC().UnixNano())
	smallestMax := bounds.Max.X
	if bounds.Max.Y < smallestMax {
		smallestMax = bounds.Max.Y
	}
	offsetMax := smallestMax / 10

    rOffset := offset{X: rand.Intn(offsetMax) - offsetMax/2, Y: rand.Intn(offsetMax) - offsetMax/2}
    gOffset := offset{X: rand.Intn(offsetMax) - offsetMax/2, Y: rand.Intn(offsetMax) - offsetMax/2}
    bOffset := offset{X: rand.Intn(offsetMax) - offsetMax/2, Y: rand.Intn(offsetMax) - offsetMax/2}

	var currentPixelColor color.Color
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {

			rOffsetColor, _, _, _ := getColorsAtOffset(input, x, y, rOffset.X, rOffset.Y)
			_, gOffsetColor, _, _ := getColorsAtOffset(input, x, y, gOffset.X, gOffset.Y)
			_, _, bOffsetColor, _ := getColorsAtOffset(input, x, y, bOffset.X, bOffset.Y)
			_, _, _, a := input.At(x, y).RGBA()
			currentPixelColor = pixelColor{
				r: rOffsetColor,
				g: gOffsetColor,
				b: bOffsetColor,
				a: a,
			}
			newImg.Set(x, y, currentPixelColor)

		}
	}
	return newImg
}

//PixelShift randomly shifts bands of pixels
func PixelShift(input image.Image) image.Image {
  rand.Seed(time.Now().UTC().UnixNano())

  bounds := input.Bounds()
  newImg := image.NewRGBA(bounds)

  offsetMax := bounds.Max.Y / 5

  offsetAmount := rand.Intn(offsetMax)
  offsetStart := rand.Intn(bounds.Max.X)
  offsetRange := rand.Intn(bounds.Max.X / 3)

  var currentPixelColor color.Color
  var r, g, b, a uint32
  for x := 0; x < bounds.Max.X; x++{
    for y := 0; y < bounds.Max.Y; y++{
      if x >= offsetStart && x < offsetStart + offsetRange{
        r, g, b, a = getColorsAtOffset(input, x, y, 0, offsetAmount)
      }else{
        r, g, b, a = input.At(x, y).RGBA()
      }
      currentPixelColor = pixelColor{
          r: r,
          g: g,
          b: b,
          a: a,
      }
      newImg.Set(x, y, currentPixelColor)
    }
  }
  return newImg
}

func getColorsAtOffset(input image.Image, x int, y int, offsetX int, offsetY int) (uint32, uint32, uint32, uint32) {
	bounds := input.Bounds()
	currentOffsetX := x + offsetX
	if currentOffsetX > bounds.Max.X {
		currentOffsetX = currentOffsetX - bounds.Max.X
	} else if currentOffsetX < 0 {
		currentOffsetX = currentOffsetX + bounds.Max.X
	}

	currentOffsetY := y + offsetY
	if currentOffsetY > bounds.Max.Y {
		currentOffsetY = currentOffsetY - bounds.Max.Y
	} else if currentOffsetY < 0 {
		currentOffsetY = currentOffsetY + bounds.Max.Y
	}
	return input.At(currentOffsetX, currentOffsetY).RGBA()
}
