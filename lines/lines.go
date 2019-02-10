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

	//for y := 0; y < HEIGHT; y += 20 {
	//	s.AddObject(line.NewLine(MIDX, MIDY, MIDX+100, float64(y), scene.RandomColor(255)))
	//}
	s.AddObject(line.NewLine(MIDX, MIDY, MIDX+100, MIDY+100, scene.RandomColor(255)))
	//s.AddObject(line.NewLine(MIDX, MIDY, MIDX+100, MIDY-100, scene.RandomColor(255)))
	//s.AddObject(line.NewLine(MIDX, MIDY, MIDX+100, MIDY, scene.RandomColor(255)))
	//s.AddObject(line.NewLine(MIDX, MIDY, MIDX, MIDY-100, scene.RandomColor(255)))
	//s.AddObject(line.NewLine(MIDX, MIDY, MIDX, MIDY+100, scene.RandomColor(255)))
	//s.AddObject(line.NewLine(MIDX, MIDY, MIDX-100, MIDY-100, scene.RandomColor(255)))
	//s.AddObject(line.NewLine(MIDX, MIDY, MIDX-100, MIDY, scene.RandomColor(255)))
	//s.AddObject(line.NewLine(MIDX, MIDY, MIDX-100, MIDY+100, scene.RandomColor(255)))
}

func main() {
	g.Start(s)
}
