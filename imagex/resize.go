package imagex

import (
	"golang.org/x/image/draw"
	"image"
)

func Resize(width, height int, scaler draw.Scaler, op draw.Op, opts *draw.Options) Transformer {
	return func(img image.Image) image.Image {
		rect := image.Rect(0, 0, width, height)
		dst := image.NewRGBA(rect)
		scaler.Scale(dst, rect, img, img.Bounds(), op, opts)
		return dst
	}
}
