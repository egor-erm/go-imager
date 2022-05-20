package main

import (
	"fmt"
	"image/color"
)

func main() {
	c := color.RGBA{100, 100, 100, 255}
	fmt.Printf("#%X%X%X", c.R, c.G, c.B)
}
