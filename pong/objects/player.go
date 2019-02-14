package objects

import (
	"github.com/sparkoo/sparkengine-examples/pong/conf"
	"github.com/sparkoo/sparkengine/core/scene"
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	*scene.Base
}

func NewPlayer(xpos float64, ypos float64) *Player {
	return &Player{Base: scene.NewBase(xpos, ypos, playerWidth, playerLength)}
}

const playerWidth = 5
const playerLength = conf.PLAYERSIZE

var playerPixels []scene.Pixel

func init() {
	playerPixels = make([]scene.Pixel, playerWidth*playerLength)
	c := 0
	for i := 0; i < playerWidth; i++ {
		for j := 0; j < playerLength; j++ {
			playerPixels[c] = scene.NewPixel(i, j, [4]byte{127, 127, 127, 127})
			c++
		}
	}
}

func (p *Player) GetPixels() []scene.Pixel {
	return playerPixels
}

func (p *Player) Move(y float64) {
	xpos, ypos := p.Base.GetPos()
	ynew := ypos + y
	if ynew < 0 {
		p.Base.MoveTo(xpos, 0)
	} else if ynew+playerLength > conf.SHEIGHT {
		p.Base.MoveTo(xpos, conf.SHEIGHT-playerLength)
	} else {
		p.Base.MoveTo(xpos, ynew)
	}
}

func (p *Player) PlayerMoveEvent(event sdl.Event) {
	switch e := event.(type) {
	case *sdl.MouseMotionEvent:
		p.Move(float64(e.YRel))
	}
}
