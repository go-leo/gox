package imagex

import "image"

// Rotate90 旋转90度
func Rotate90(oldImage image.Image) image.Image {
	newImage := image.NewRGBA(image.Rect(0, 0, oldImage.Bounds().Dy(), oldImage.Bounds().Dx()))
	for x := oldImage.Bounds().Min.Y; x < oldImage.Bounds().Max.Y; x++ {
		for y := oldImage.Bounds().Max.X - 1; y >= oldImage.Bounds().Min.X; y-- {
			newImage.Set(oldImage.Bounds().Max.Y-x, y, oldImage.At(y, x))
		}
	}
	return newImage
}

// Rotate180 旋转180度
func Rotate180(oldImage image.Image) image.Image {
	newImage := image.NewRGBA(image.Rect(0, 0, oldImage.Bounds().Dx(), oldImage.Bounds().Dy()))
	for x := oldImage.Bounds().Min.X; x < oldImage.Bounds().Max.X; x++ {
		for y := oldImage.Bounds().Min.Y; y < oldImage.Bounds().Max.Y; y++ {
			newImage.Set(oldImage.Bounds().Max.X-x, oldImage.Bounds().Max.Y-y, oldImage.At(x, y))
		}
	}
	return newImage
}

// Rotate270 旋转270度
func Rotate270(oldImage image.Image) image.Image {
	newImage := image.NewRGBA(image.Rect(0, 0, oldImage.Bounds().Dy(), oldImage.Bounds().Dx()))
	for x := oldImage.Bounds().Min.Y; x < oldImage.Bounds().Max.Y; x++ {
		for y := oldImage.Bounds().Max.X - 1; y >= oldImage.Bounds().Min.X; y-- {
			newImage.Set(x, oldImage.Bounds().Max.X-y, oldImage.At(y, x))
		}
	}
	return newImage
}
