package imagex

import "image"

type Transformer func(image.Image) image.Image

func Transform(img image.Image, transformers ...Transformer) image.Image {
	for i := len(transformers) - 1; i >= 0; i-- {
		img = transformers[i](img)
	}
	return img
}
