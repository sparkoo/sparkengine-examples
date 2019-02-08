package objects

import (
	"github.com/sparkoo/pong/conf"
	"github.com/sparkoo/sparkengine/scene"
	"github.com/veandco/go-sdl2/sdl"
)

type Player struct {
	xpos float64
	ypos float64
}

func NewPlayer(xpos float64, ypos float64) *Player {
	return &Player{xpos, ypos}
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

func (p *Player) GetXoffset() int {
	return int(p.xpos)
}

func (p *Player) GetYoffset() int {
	return int(p.ypos)
}

func (p *Player) GetXsize() int {
	return playerWidth
}

func (p *Player) GetYsize() int {
	return playerLength
}

func (p *Player) GetPixels() []scene.Pixel {
	return playerPixels
}

func (p *Player) Move(y float64) {
	ynew := p.ypos + y
	if ynew < 0 {
		p.MoveTo(p.xpos, 0)
	} else if ynew + playerLength > conf.SHEIGHT {
		p.MoveTo(p.xpos, conf.SHEIGHT - playerLength)
	} else {
		p.MoveTo(p.xpos, ynew)
	}
}

func (p *Player) MoveBy(x float64, y float64) {
	p.xpos += x
	p.ypos += y
}

func (p *Player) MoveTo(x float64, y float64) {
	p.xpos = x
	p.ypos = y
}

func (p *Player) PlayerMoveEvent(event sdl.Event) {
	switch e := event.(type) {
	case *sdl.MouseMotionEvent:
		p.Move(float64(e.YRel))
	}
}