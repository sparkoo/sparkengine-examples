package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/core/scene"
	"github.com/sparkoo/sparkengine/core/scene/olib/shape"
	"math/rand"
	"time"
)

//noinspection ALL
const (
	FPS    = 30
	WIDTH  = 320
	HEIGHT = 240
	MIDX   = float64(WIDTH) / 2
	MIDY   = float64(HEIGHT) / 2
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
	s = scene.NewScene(scene.NoopTick)

	rand.Seed(time.Now().UnixNano())

	// right
	for y := int(MIDY - 100); y < int(MIDY + 100); y += 10 {
		s.AddObject(shape.NewLine(MIDX, MIDY, MIDX+100, float64(y), scene.RandomColor(255)))
	}

	// left
	for y := int(MIDY - 100); y < int(MIDY + 100); y += 10 {
		s.AddObject(shape.NewLine(MIDX, MIDY, MIDX-100, float64(y), scene.RandomColor(255)))
	}

	// top
	for x := int(MIDX - 100); x < int(MIDX + 100); x += 10 {
		s.AddObject(shape.NewLine(MIDX, MIDY, float64(x), MIDY-100, scene.RandomColor(255)))
	}

	// bottom
	for x := int(MIDX - 100); x < int(MIDX + 100); x += 10 {
		s.AddObject(shape.NewLine(MIDX, MIDY, float64(x), MIDY+100, scene.RandomColor(255)))
	}
}

func main() {
	g.Start(s)
}
