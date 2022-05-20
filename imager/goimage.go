package goimage

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type goimage struct {
	name  string
	image *image.RGBA
}

func New(name string, xmax int, ymax int) goimage {
	b := image.Rect(0, 0, xmax, ymax)
	image := image.NewRGBA(b)

	return goimage{name, image}
}

func NewWithCorners(name string, xmin int, ymin int, xmax int, ymax int) goimage {
	b := image.Rect(xmin, ymin, xmax, ymax)
	image := image.NewRGBA(b)

	return goimage{name, image}
}

func (gimg *goimage) Save() error {
	file, err := os.Create(gimg.name)

	if err != nil || file == nil {
		file, err = os.Open(gimg.name)
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

func (gimg *goimage) SetPixel(x int, y int, color color.RGBA) {
	gimg.image.SetRGBA(x, y, color)
}

func (gimg *goimage) SetHEXPixel(x int, y int, c string) {
	color := gocolor.HEXtoRGBA(c)
	gimg.image.SetRGBA(x, y, color)
}

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
