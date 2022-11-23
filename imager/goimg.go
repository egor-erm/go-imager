package goimage

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	gocolor "github.com/egor-erm/goimager/manager"
	"github.com/go-gl/mathgl/mgl32"
)

type goimage struct {
	Name  string
	image *image.RGBA
}

func New(name string, xmax int, ymax int) *goimage {
	b := image.Rect(0, 0, xmax, ymax)
	image := image.NewRGBA(b)

	return &goimage{name, image}
}

func NewWithCorners(name string, xmin int, ymin int, xmax int, ymax int) *goimage {
	b := image.Rect(xmin, ymin, xmax, ymax)
	image := image.NewRGBA(b)

	return &goimage{name, image}
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

func (gimg *goimage) SetPixel(x int, y int, color color.RGBA) {
	gimg.image.SetRGBA(x, y, color)
}

func (gimg *goimage) SetPixelByVector(vec mgl32.Vec2, color color.RGBA) {
	gimg.image.SetRGBA(int(vec.X()), int(vec.Y()), color)
}

func (gimg *goimage) SetHexPixel(x int, y int, c string) {
	color := gocolor.HexToRGBA(c)
	gimg.SetPixel(x, y, color)
}

func (gimg *goimage) SetHexPixelByVector(vec mgl32.Vec2, c string) {
	color := gocolor.HexToRGBA(c)
	gimg.SetPixel(int(vec.X()), int(vec.Y()), color)
}

func (gimg *goimage) SetHexAlphaPixel(x int, y int, c string, alpha uint8) {
	color := gocolor.HexAlphaToRGBA(c, alpha)
	gimg.SetPixel(x, y, color)
}

func (gimg *goimage) SetHexAlphaPixelByVector(vec mgl32.Vec2, c string, alpha uint8) {
	color := gocolor.HexAlphaToRGBA(c, alpha)
	gimg.SetPixel(int(vec.X()), int(vec.Y()), color)
}

func (gimg *goimage) ClearPixel(x int, y int) {
	gimg.SetPixel(x, y, color.RGBA{0, 0, 0, 0})
}

func (gimg *goimage) ClearPixelByVectors(vec mgl32.Vec2) {
	gimg.SetPixel(int(vec.X()), int(vec.Y()), color.RGBA{0, 0, 0, 0})
}

func (gimg *goimage) DrowRect(x1 int, y1 int, x2 int, y2 int, color color.RGBA) {
	for y := y1; y <= y2; y++ {
		for x := x1; x < x2; x++ {
			gimg.SetPixel(x, y, color)
		}
	}
}

func (gimg *goimage) DrowHexRect(x1 int, y1 int, x2 int, y2 int, c string) {
	color := gocolor.HexToRGBA(c)
	for y := y1; y <= y2; y++ {
		for x := x1; x < x2; x++ {
			gimg.SetPixel(x, y, color)
		}
	}
}

func (gimg *goimage) DrowHexAlphaRect(x1 int, y1 int, x2 int, y2 int, c string, alpha uint8) {
	color := gocolor.HexAlphaToRGBA(c, alpha)
	for y := y1; y <= y2; y++ {
		for x := x1; x < x2; x++ {
			gimg.SetPixel(x, y, color)
		}
	}
}

func (gimg *goimage) FillAll(color color.RGBA) {
	for y := gimg.image.Bounds().Min.Y; y < gimg.image.Bounds().Max.Y; y++ {
		for x := gimg.image.Bounds().Min.X; x < gimg.image.Bounds().Max.X; x++ {
			gimg.SetPixel(x, y, color)
		}
	}
}

func (gimg *goimage) FillAllHex(c string) {
	color := gocolor.HexToRGBA(c)
	for y := gimg.image.Bounds().Min.Y; y < gimg.image.Bounds().Max.Y; y++ {
		for x := gimg.image.Bounds().Min.X; x < gimg.image.Bounds().Max.X; x++ {
			gimg.SetPixel(x, y, color)
		}
	}
}

func (gimg *goimage) FillAllHexAlpha(c string, alpha uint8) {
	color := gocolor.HexAlphaToRGBA(c, alpha)
	for y := gimg.image.Bounds().Min.Y; y < gimg.image.Bounds().Max.Y; y++ {
		for x := gimg.image.Bounds().Min.X; x < gimg.image.Bounds().Max.X; x++ {
			gimg.SetPixel(x, y, color)
		}
	}
}

func (gimg *goimage) GetPixel(x int, y int) color.Color {
	return gimg.image.At(x, y)
}

func (gimg *goimage) GetRawImage() *image.RGBA {
	return gimg.image
}
