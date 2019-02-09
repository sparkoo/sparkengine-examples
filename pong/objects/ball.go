package objects

import (
	"github.com/sparkoo/sparkengine/scene"
)

type Ball struct {
	*scene.Base
	xvel float64
	yvel float64
}

func NewBall(xpos float64, ypos float64, xvel float64, yvel float64) *Ball {
	return &Ball{scene.NewBase(xpos, ypos, size, size), xvel, yvel}
}

const size = 5

var ballPixels []scene.Pixel

func init() {
	ballPixels = make([]scene.Pixel, size*size)
	c := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			ballPixels[c] = scene.NewPixel(i, j, [4]byte{127, 127, 127, 127})
			c++
		}
	}
}

func (b *Ball) GetPixels() []scene.Pixel {
	return ballPixels
}

func (b *Ball) Move() {
	b.Base.MoveBy(b.xvel, b.yvel)
}

func (b *Ball) GetVel() (float64, float64) {
	return b.xvel, b.yvel
}

func (b *Ball) SetXVel(xvel float64) {
	b.xvel = xvel
}

func (b *Ball) SetYVel(yvel float64) {
	b.yvel = yvel
}

func (b *Ball) BounceX(yvel float64) {
	b.xvel *= -1
	b.yvel = yvel
}