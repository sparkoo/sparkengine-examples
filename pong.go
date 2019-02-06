package main

import (
	"github.com/sparkoo/pong/conf"
	"github.com/sparkoo/pong/objects"
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/scene"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
)

var (
	player1 = objects.NewPlayer(conf.SWIDTH*.05, conf.SHEIGHT*.5)
	player2 = objects.NewPlayer(conf.SWIDTH*.95, conf.SHEIGHT*.5)
)

func main() {
	game := core.NewGame(core.NewConf(conf.FPS, conf.FPS*2, conf.SWIDTH, conf.SHEIGHT))

	s := scene.NewScene(tick)
	s.AddObject(player1)
	s.AddObject(player2)

	s.AddEventListener(sdl.MOUSEMOTION, player1.PlayerMoveEvent)

	game.Start(s)
}

func tick() {
	player2.Move(rand.Int31n(11) - 5)
}
