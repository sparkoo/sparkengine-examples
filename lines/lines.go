package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/scene"
	"github.com/sparkoo/sparkengine/scene/olib/line"
	"math/rand"
	"time"
)

const (
	FPS    = 30
	WIDTH  = 320
	HEIGHT = 240
	MIDX   = float64(WIDTH) / 2
	MIDY   = float64(HEIGHT) / 2
)

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

	l1 := line.NewLine(MIDX, MIDY, MIDX+100, MIDY+100, scene.RandomColor(255))
	l2 := line.NewLine(MIDX, MIDY, MIDX+100, MIDY-100, scene.RandomColor(255))
	l3 := line.NewLine(MIDX, MIDY, MIDX+100, MIDY, scene.RandomColor(255))
	l4 := line.NewLine(MIDX, MIDY, MIDX, MIDY-100, scene.RandomColor(255))
	l5 := line.NewLine(MIDX, MIDY, MIDX, MIDY+100, scene.RandomColor(255))
	l6 := line.NewLine(MIDX, MIDY, MIDX-100, MIDY-100, scene.RandomColor(255))
	l7 := line.NewLine(MIDX, MIDY, MIDX-100, MIDY, scene.RandomColor(255))
	l8 := line.NewLine(MIDX, MIDY, MIDX-100, MIDY+100, scene.RandomColor(255))

	s.AddObjects(
		l1,
		l2,
		l3,
		l4,
		l5,
		l6,
		l7,
		l8)
}

func main() {
	g.Start(s)
}
