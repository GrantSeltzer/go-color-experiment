package colorexp

import "image"

type pixelManipulation func(pixel) pixel

// ApplyToAllPixels blah
func ApplyToAllPixels(img image.RGBA, fn pixelManipulation) error {

	pixels := loadPixels(img)
	newPixArray := []uint8{}

	for _, pixel := range pixels {
		pixel = fn(pixel)
		newPixArray = append(newPixArray, pixel.r)
		newPixArray = append(newPixArray, pixel.g)
		newPixArray = append(newPixArray, pixel.b)
		newPixArray = append(newPixArray, pixel.a)
	}

	img.Pix = newPixArray
	saveImage(img)
	return nil
}

func Invert(pix pixel) pixel {
	pix.r = pix.g
	return pix
}
