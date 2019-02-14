package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/core/scene"
)

func main() {
	g := core.NewGame(core.NewConf30T960W())

	character := NewCharacter(0, 0)

	s := scene.NewScene(func(gameTickCounter int64, sceneTickCounter int64) {
		if gameTickCounter%3 == 0 {
			character.Tick()
		}
	})

	s.AddObject(character)

	g.Start(s)
}
