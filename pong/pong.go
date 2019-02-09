package main

import (
	"github.com/sparkoo/sparkengine-examples/pong/conf"
	"github.com/sparkoo/sparkengine-examples/pong/objects"
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/scene"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

var (
	player1 = objects.NewPlayer(conf.SWIDTH*.05, conf.SHEIGHT*.5)
	player2 = objects.NewPlayer(conf.SWIDTH*.95, conf.SHEIGHT*.5)
	ball    = objects.NewBall(conf.SWIDTH*.5, conf.SHEIGHT*.5, -5, 5)
)

var game *core.Game

func main() {
	game = core.NewGame(core.NewConf(conf.FPS, conf.FPS*2, conf.SWIDTH, conf.SHEIGHT))

	s := scene.NewScene(tick)
	s.AddObject(player1)
	s.AddObject(player2)
	s.AddObject(ball)

	s.AddEventListener(sdl.MOUSEMOTION, player1.PlayerMoveEvent)

	game.Start(s)
}

func tick() {

	player2.Move((rand.Float64() * 11) - 5)
	moveBall(ball)

	if scene.Collides(player1, ball) {
		ball.BounceX((rand.Float64() * 11) - 5)
	}

	if scene.Collides(player2, ball) {
		ball.BounceX((rand.Float64() * 11) - 5)
	}
}

func moveBall(b *objects.Ball) {
	xvel, yvel := b.GetVel()
	xpos, ypos := b.GetPos()
	xPot := int(xpos + xvel)
	if xPot < 0 || xPot+b.GetXsize() >= conf.SWIDTH {
		b.SetXVel(xvel * -1)
		game.Quit()
	}

	yPot := int(ypos + yvel)
	if yPot < 0 || yPot+b.GetYsize() >= conf.SHEIGHT {
		b.SetYVel(yvel * -1)
	}

	b.Move()
}