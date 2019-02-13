package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/scene"
)

func main() {
	g := core.NewGame(core.NewConf30T320W())

	loadingSprite := NewObj(0, 0, 50, 100)

	s := scene.NewScene(func(gameTickCounter int64, sceneTickCounter int64) {
		if sceneTickCounter > 0 && sceneTickCounter%5 == 0 {
			loadingSprite.TickAnimation()
		}
	})
	s.AddObject(loadingSprite)

	g.Start(s)
}
