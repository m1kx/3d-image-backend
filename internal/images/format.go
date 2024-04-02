package images

import (
	"image"
	"image/draw"
)

func SplitImage(img *image.Image) []image.Image {
	sImgs := make([]image.Image, 3)
	bounds := (*img).Bounds()
	for i := range len(sImgs) {
		sImg := image.NewRGBA(image.Rect(0, 0, bounds.Dx()/3, bounds.Dy()))
		draw.Draw(sImg, bounds, (*img), image.Point{(bounds.Dx() / 3) * i, 0}, draw.Src)
		sImgs[i] = sImg
	}
	return sImgs
}
