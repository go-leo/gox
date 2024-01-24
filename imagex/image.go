package imagex

import (
	"golang.org/x/image/bmp"
	"golang.org/x/image/webp"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"mime/multipart"
)

func Png(dest io.Writer, src io.Reader, transformers ...Transformer) error {
	img, err := png.Decode(src)
	if err != nil {
		return err
	}
	return png.Encode(dest, Transform(img, transformers...))
}

func Jpeg(dest io.Writer, src io.Reader, opts *jpeg.Options, transformers ...Transformer) error {
	img, err := jpeg.Decode(src)
	if err != nil {
		return err
	}
	return jpeg.Encode(dest, Transform(img, transformers...), opts)
}

func Gif(dest io.Writer, src io.Reader, opt *gif.Options, transformers ...Transformer) error {
	img, err := gif.Decode(src)
	if err != nil {
		return err
	}
	return gif.Encode(dest, Transform(img, transformers...), opt)
}

func Bmp(dest io.Writer, src io.Reader, transformers ...Transformer) error {
	img, err := bmp.Decode(src)
	if err != nil {
		return err
	}
	return bmp.Encode(dest, Transform(img, transformers...))
}

func Webp(dest io.Writer, src multipart.File, transformers ...Transformer) error {
	img, err := webp.Decode(src)
	if err != nil {
		return err
	}
	return png.Encode(dest, Transform(img, transformers...))
}
