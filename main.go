package main

import (
	goimage "github.com/egor-erm/goimager/imager"
)

func main() {
	img := goimage.New("eg.png", 100, 100)
	img.FillAllHex("#66FF00")
	img.DrowHexRect(10, 10, 89, 89, "#00FFD5")
	img.Save()
}
