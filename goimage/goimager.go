package goimage

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strings"
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

func ExportExpanded(name string, multi int) error {
	img, err := Open(name)
	if err != nil {
		return err
	}

	tags := strings.Split(name, ".")
	if len(tags) != 2 {
		return fmt.Errorf("format error: " + name)
	}

	exp := New(tags[0]+"-exp."+tags[1], img.image.Bounds().Max.X*multi, img.image.Bounds().Max.Y*multi)
	for x := 0; x < img.GetRawImage().Bounds().Max.X; x++ {
		for y := 0; y < img.GetRawImage().Bounds().Max.Y; y++ {
			r, g, b, a := img.GetPixel(x, y).RGBA()

			if a == 0 {
				continue
			}

			color := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}

			exp.DrowRect(x*multi, y*multi, x*multi+multi-1, y*multi+multi-1, color)
		}
	}

	return exp.Save()
}

func ExportAll(folder string, multi int) error {
	files, err := os.ReadDir(folder)
	if err != nil {
		return err
	}

	for _, f := range files {
		err = ExportExpanded(folder+"/"+f.Name(), multi)
		if err != nil {
			return err
		}
	}

	return nil
}

func Pixelate(name string, multi int) (*goimage, error) {
	img, err := Open(name)
	if err != nil {
		return nil, err
	}

	tags := strings.Split(name, ".")
	if len(tags) != 2 {
		return nil, fmt.Errorf("format error: " + name)
	}

	temp := New(tags[0]+"-px."+tags[1], img.image.Bounds().Max.X/multi, img.image.Bounds().Max.Y/multi)
	for x := 0; x < img.GetRawImage().Bounds().Max.X; x += multi {
		if x+multi > img.GetRawImage().Bounds().Max.X {
			break
		}

		for y := 0; y < img.GetRawImage().Bounds().Max.Y; y += multi {
			if y+multi > img.GetRawImage().Bounds().Max.Y {
				break
			}

			var mr, mg, mb uint32
			for ix := x; ix < x+multi; ix++ {
				for iy := y; iy < y+multi; iy++ {
					rgb := img.GetPixel(ix, iy).(color.RGBA)

					mr += uint32(rgb.R)
					mg += uint32(rgb.G)
					mb += uint32(rgb.B)
				}
			}

			color := color.RGBA{uint8(mr / uint32(multi*multi)), uint8(mg / uint32(multi*multi)), uint8(mb / uint32(multi*multi)), 255}
			temp.SetPixel(x/multi, y/multi, color)
		}
	}

	return temp, nil
}

func PixelateSave(name string, multi int) error {
	img, err := Pixelate(name, multi)
	if err != nil {
		return err
	}

	return img.Save()
}
