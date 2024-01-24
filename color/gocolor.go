package color

import (
	"fmt"
	"image/color"
)

func RGBAtoHex(c color.RGBA) string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

func HexToRGBA(hex string) color.RGBA {
	var c color.RGBA
	c.A = 255
	switch len(hex) {
	case 7:
		_, _ = fmt.Sscanf(hex, "#%02X%02X%02X", &c.R, &c.G, &c.B)
		return c
	case 4:
		_, _ = fmt.Sscanf(hex, "#%1X%1X%1X", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
		return c
	default:
		fmt.Printf("invalid length, must be 7 or 4")
		return c
	}
}

func HexAlphaToRGBA(hex string, alpha uint8) color.RGBA {
	var c color.RGBA
	c.A = uint8(alpha)
	switch len(hex) {
	case 7:
		_, _ = fmt.Sscanf(hex, "#%02X%02X%02X", &c.R, &c.G, &c.B)
		return c
	case 4:
		_, _ = fmt.Sscanf(hex, "#%1X%1X%1X", &c.R, &c.G, &c.B)
		c.R *= 17
		c.G *= 17
		c.B *= 17
		return c
	default:
		fmt.Printf("invalid length, must be 7 or 4")
		return c
	}
}
