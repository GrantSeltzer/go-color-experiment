package colorexp

import (
	"bytes"
	"image"
	"image/jpeg"
	"log"
	"os"
)

type pixel struct {
	r, g, b, a uint32
}

func LoadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return img
}

func loadPixels(img image.RGBA) []pixel {
	bounds := img.Bounds()
	pixels := make([]pixel, bounds.Dx()*bounds.Dy())

	for i := 0; i < bounds.Dx()*bounds.Dy(); i++ {
		x := i % bounds.Dx()
		y := i / bounds.Dx()
		r, g, b, a := img.At(x, y).RGBA()
		pixels[i].r = r
		pixels[i].g = g
		pixels[i].b = b
		pixels[i].a = a
	}
	return pixels
}

func SaveImage(img image.Image) error {
	var buf bytes.Buffer
	jpeg.Encode(&buf, img, nil)
	x, err := os.Create("exportedImage.jpeg")
	if err != nil {
		return err
	}
	x.Write(buf.Bytes())
	return nil
}
