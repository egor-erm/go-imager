package main

import (
	goimage "github.com/egor-erm/goimager/imager"
)

func main() {
	img := goimage.New("eg.png", 100, 100)
	img.DrowHexRect(0, 0, 99, 99, "#00FFD5")
	img.Save()
}
