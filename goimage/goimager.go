package goimage

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func Open(name string) (*goimage, error) {
	file, err := os.Open(name)
	if err != nil {
		return &goimage{}, fmt.Errorf("error opening file: %s", err)
	}

	img, err := png.Decode(file)
	if err != nil {
		return &goimage{}, fmt.Errorf("error decoding image: %s", err)
	}

	rgbaing := imageToRGBA(img)
	file.Close()

	return &goimage{name, rgbaing}, nil
}

func imageToRGBA(src image.Image) *image.RGBA {
	if dst, ok := src.(*image.RGBA); ok {
		return dst
	}

	b := src.Bounds()
	img := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(img, img.Bounds(), src, b.Min, draw.Src)
	return img
}

func (gimg *goimage) Save() error {
	file, err := os.Create(gimg.Name)

	if err != nil || file == nil {
		file, err = os.Open(gimg.Name)
		if err != nil {
			return fmt.Errorf("error opening file: %s", err)
		}
	}

	err = png.Encode(file, gimg.image)
	if err != nil {
		return fmt.Errorf("error encoding image: %s", err)
	}

	file.Close()
	return nil
}

func (gimg *goimage) SaveNewImage(name string) error {
	file, err := os.Create(name)

	if err != nil || file == nil {
		file, err = os.Open(name)
		if err != nil {
			return fmt.Errorf("error opening file: %s", err)
		}
	}

	err = png.Encode(file, gimg.image)
	if err != nil {
		return fmt.Errorf("error encoding image: %s", err)
	}

	file.Close()
	return nil
}