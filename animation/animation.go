package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/core/scene"
)

func main() {
	g := core.NewGame(core.NewConf30T640W())

	loadingSprite := NewObj(0, 0, 50, 100)
	viking := NewViking(100, 100, 210, 222)

	s := scene.NewScene(func(gameTickCounter int64, sceneTickCounter int64) {
		if sceneTickCounter > 0 && sceneTickCounter%5 == 0 {
			loadingSprite.TickAnimation()
		}
		viking.TickAnimation()
	})
	s.AddObjects(loadingSprite, viking)

	g.Start(s)
}
