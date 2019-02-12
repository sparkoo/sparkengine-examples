package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/scene"
	"github.com/sparkoo/sparkengine/scene/olib/shape"
	"math/rand"
	"time"
)

//noinspection ALL
const (
	FPS    = 30
	WIDTH  = 640
	HEIGHT = 480
)

//noinspection ALL
var (
	g *core.Game
	c *core.Conf
	s *scene.Scene
)

func init() {
	c = core.NewConf(FPS, FPS*2, WIDTH, HEIGHT)
	g = core.NewGame(c)
	s = scene.NewScene(addRandomLineTick)

	rand.Seed(time.Now().UnixNano())
}

func addRandomLineTick(_ int64, _ int64) {
	s.AddObject(shape.NewLine(
		rand.Float64()*WIDTH,
		rand.Float64()*HEIGHT,
		rand.Float64()*WIDTH,
		rand.Float64()*HEIGHT,
		scene.RandomColor(50)))
}

func main() {
	g.Start(s)
}
