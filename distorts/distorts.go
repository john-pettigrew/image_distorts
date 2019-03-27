package distorts

import (
	"image"
	"image/color"
	"math/rand"
)

type pixelColor struct {
	r, g, b, a uint32
}

func (c pixelColor) RGBA() (uint32, uint32, uint32, uint32) {
	return c.r, c.g, c.b, c.a
}

//ChromaticAberation returns a copy of input with a random chromatic aberation effect applied
func ChromaticAberation(input image.Image) image.Image {
	rand.Seed(42)
	offsetMax := 20
	rOffsetX := rand.Intn(offsetMax) - offsetMax/2
	rOffsetY := rand.Intn(offsetMax) - offsetMax/2

	bounds := input.Bounds()
	newImg := image.NewRGBA(bounds)

	var currentPixelColor color.Color
	var currentROffsetX, currentROffsetY int
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			_, g, b, a := input.At(x, y).RGBA()
			currentROffsetX = x + rOffsetX
			if currentROffsetX > bounds.Max.X {
				currentROffsetX = currentROffsetX - bounds.Max.X
			} else if currentROffsetX < 0 {
				currentROffsetX = currentROffsetX + bounds.Max.X
			}

			currentROffsetY = y + rOffsetY
			if currentROffsetY > bounds.Max.Y {
				currentROffsetY = currentROffsetY - bounds.Max.Y
			} else if currentROffsetY < 0 {
				currentROffsetY = currentROffsetY + bounds.Max.Y
			}
			rOffsetColor, _, _, _ := input.At(currentROffsetX, currentROffsetY).RGBA()
			currentPixelColor = pixelColor{
				r: rOffsetColor,
				g: g,
				b: b,
				a: a,
			}
			newImg.Set(x, y, currentPixelColor)

		}
	}
	return newImg
}
