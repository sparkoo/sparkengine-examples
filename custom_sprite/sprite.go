package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/core/scene"
	"github.com/sparkoo/sparkengine/core/scene/olib/sprite"
)

func main() {
	g := core.NewGame(core.NewConf(30, 60, 640, 480))
	s := scene.NewScene(scene.NoopTick)

	s.AddObject(sprite.NewSprite(100, 100, 15, 10, initPixels))

	g.Start(s)
}
