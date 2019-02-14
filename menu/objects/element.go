package objects

import (
	"github.com/sparkoo/sparkengine-examples/menu/conf"
	"github.com/sparkoo/sparkengine/core/scene"
)

type Element struct {
	*scene.Base
}

var elementPixels []scene.Pixel

const (
	eX = 100
	eY = 5
)

func init() {
	elementPixels = make([]scene.Pixel, eX*eY)
	c := 0
	for i := 0; i < eX; i++ {
		for j := 0; j < eY; j++ {
			elementPixels[c] = scene.NewPixel(i, j, [4]byte{127, 127, 127, 127})
			c++
		}
	}
}

func NewElement(ypos float64) *Element {
	y := float64(conf.SHEIGHT) * 0.1 * ypos
	return &Element{Base: scene.NewBase(conf.MENU_XPOS, y, eX, eY)}
}

func (*Element) GetPixels() []scene.Pixel {
	return elementPixels
}
