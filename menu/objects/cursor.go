package objects

import (
	"github.com/sparkoo/sparkengine-examples/menu/conf"
	"github.com/sparkoo/sparkengine/core/scene"
)

var cursorPixels []scene.Pixel

const (
	cSize = eY
)

func init() {
	cursorPixels = make([]scene.Pixel, cSize*cSize)
	c := 0
	for i := 0; i < cSize; i++ {
		for j := 0; j < cSize; j++ {
			cursorPixels[c] = scene.NewPixel(i, j, [4]byte{127, 127, 127, 127})
			c++
		}
	}
}

type Cursor struct {
	*scene.Base
	Pos    int
	Fields int
}

func NewCursor(ypos float64, fields int) *Cursor {
	y := float64(conf.SHEIGHT) * 0.1 * (ypos + 1)
	return &Cursor{Base: scene.NewBase(conf.CURSOR_YPOS, y, eX, eY), Pos: int(ypos), Fields: fields}
}

func (c *Cursor) MoveToPos(ypos int) {
	x, _ := c.GetPos()
	y := float64(conf.SHEIGHT) * 0.1 * (float64(ypos) + 1)
	c.MoveTo(x, y)
}

func (*Cursor) GetPixels() []scene.Pixel {
	return cursorPixels
}
