package goimage

import (
	"image"
	"image/color"

	gocolor "github.com/egor-erm/go-imager/color"
	"github.com/go-gl/mathgl/mgl32"
)

type goimage struct {
	Name  string
	image *image.RGBA
}

func New(name string, xmax, ymax int) *goimage {
	b := image.Rect(0, 0, xmax, ymax)
	image := image.NewRGBA(b)

	return &goimage{name, image}
}

func NewWithCorners(name string, xmin, ymin, xmax, ymax int) *goimage {
	b := image.Rect(xmin, ymin, xmax, ymax)
	image := image.NewRGBA(b)

	return &goimage{name, image}
}

func (gimg *goimage) SetPixel(x, y int, color color.RGBA) {
	gimg.image.SetRGBA(x, y, color)
}

func (gimg *goimage) SetPixelColor(x, y int, color color.RGBA) {
	gimg.image.SetRGBA(x, y, color)
}

func (gimg *goimage) SetPixelByVector(vec mgl32.Vec2, color color.RGBA) {
	gimg.image.SetRGBA(int(vec.X()), int(vec.Y()), color)
}

func (gimg *goimage) SetHexPixel(x, y int, c string) {
	color := gocolor.HexToRGBA(c)
	gimg.SetPixel(x, y, color)
}

func (gimg *goimage) SetHexPixelByVector(vec mgl32.Vec2, c string) {
	color := gocolor.HexToRGBA(c)
	gimg.SetPixel(int(vec.X()), int(vec.Y()), color)
}

func (gimg *goimage) SetHexAlphaPixel(x, y int, c string, alpha uint8) {
	color := gocolor.HexAlphaToRGBA(c, alpha)
	gimg.SetPixel(x, y, color)
}

func (gimg *goimage) SetHexAlphaPixelByVector(vec mgl32.Vec2, c string, alpha uint8) {
	color := gocolor.HexAlphaToRGBA(c, alpha)
	gimg.SetPixel(int(vec.X()), int(vec.Y()), color)
}

func (gimg *goimage) ClearPixel(x, y int) {
	gimg.SetPixel(x, y, color.RGBA{0, 0, 0, 0})
}

func (gimg *goimage) ClearPixelByVectors(vec mgl32.Vec2) {
	gimg.SetPixel(int(vec.X()), int(vec.Y()), color.RGBA{0, 0, 0, 0})
}

func (gimg *goimage) DrowRect(x1, y1, x2, y2 int, color color.RGBA) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			gimg.SetPixel(x, y, color)
		}
	}
}

func (gimg *goimage) DrowHexRect(x1, y1, x2, y2 int, c string) {
	color := gocolor.HexToRGBA(c)
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			gimg.SetPixel(x, y, color)
		}
	}
}

func (gimg *goimage) DrowHexAlphaRect(x1, y1, x2, y2 int, c string, alpha uint8) {
	color := gocolor.HexAlphaToRGBA(c, alpha)
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
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

func (gimg *goimage) GetPixel(x, y int) color.Color {
	return gimg.image.At(x, y)
}

func (gimg *goimage) GetRawImage() *image.RGBA {
	return gimg.image
}

func (gimg *goimage) DrawImage(img *goimage, x, y int) {
	for xc := 0; xc < img.GetRawImage().Bounds().Max.X; xc++ {
		for yc := 0; yc < img.GetRawImage().Bounds().Max.Y; yc++ {
			r, g, b, a := img.GetPixel(xc, yc).RGBA()

			if a == 0 {
				continue
			}

			color := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			gimg.SetPixel(x+xc, y+yc, color)
		}
	}
}

func (gimg *goimage) DrawMirroredXImage(img *goimage, x, y int) {
	for xc := 0; xc < img.GetRawImage().Bounds().Max.X; xc++ {
		for yc := 0; yc < img.GetRawImage().Bounds().Max.Y; yc++ {
			r, g, b, a := img.GetPixel(xc, yc).RGBA()

			if a == 0 {
				continue
			}

			color := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			gimg.SetPixel(x+(img.GetRawImage().Bounds().Max.X-xc), y+yc, color)
		}
	}
}

func (gimg *goimage) DrawMirroredYImage(img *goimage, x, y int) {
	for xc := 0; xc < img.GetRawImage().Bounds().Max.X; xc++ {
		for yc := 0; yc < img.GetRawImage().Bounds().Max.Y; yc++ {
			r, g, b, a := img.GetPixel(xc, yc).RGBA()

			if a == 0 {
				continue
			}

			color := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			gimg.SetPixel(x+xc, y+(img.GetRawImage().Bounds().Max.Y-yc), color)
		}
	}
}

func (gimg *goimage) DrawMirroredXYImage(img *goimage, x, y int) {
	for xc := 0; xc < img.GetRawImage().Bounds().Max.X; xc++ {
		for yc := 0; yc < img.GetRawImage().Bounds().Max.Y; yc++ {
			r, g, b, a := img.GetPixel(xc, yc).RGBA()

			if a == 0 {
				continue
			}

			color := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			gimg.SetPixel(x+(img.GetRawImage().Bounds().Max.X-xc), y+(img.GetRawImage().Bounds().Max.Y-yc), color)
		}
	}
}
