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
	ball    = objects.NewBall(conf.SWIDTH*.5, conf.SHEIGHT*.5, -5, randSpeed())
)

var game *core.Game
var running = false

func main() {
	game = core.NewGame(core.NewConf(conf.FPS, conf.FPS*2, conf.SWIDTH, conf.SHEIGHT))

	s := scene.NewScene(tick)
	s.AddObject(player1)
	s.AddObject(player2)
	s.AddObject(ball)

	s.AddEventListener(sdl.MOUSEMOTION, player1.PlayerMoveEvent)
	s.AddEventListener(sdl.MOUSEBUTTONDOWN, startGame)

	game.Start(s)
}

func tick() {
	if !running {
		return
	}
	moveBall(ball)
	playerAi(ball, player2)

	if scene.Collides(player1, ball) {
		ball.BounceX(randSpeed())
	}

	if scene.Collides(player2, ball) {
		ball.BounceX(randSpeed())
	}
}

func startGame(_ sdl.Event) {
	running = true
}

func playerAi(b *objects.Ball, p *objects.Player) {
	if b.GetYoffset() > p.GetYoffset() + (p.GetYsize() / 2) {
		p.Move(rand.Float64() * 5)
	} else {
		p.Move(rand.Float64() * -5)
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

func randSpeed() float64 {
	return (rand.Float64() * 11) - 5
}
