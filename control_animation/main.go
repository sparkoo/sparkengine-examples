package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/core/scene"
	"github.com/sparkoo/sparkengine/core/scene/image"
	"github.com/sparkoo/sparkengine/core/scene/olib/sprite"
)

func main() {
	g := core.NewGame(core.NewConf30T960W())

	character := NewCharacter(0, 0)
	background := sprite.NewSprite(0, 0, 960, 540, func() []scene.Pixel {
		return image.LoadFullImage("assets/background.png")
	})

	s := scene.NewScene(func(gameTickCounter int64, sceneTickCounter int64) {
		if gameTickCounter%3 == 0 {
			character.Tick()
		}
	})

	s.AddObjects(background, character)

	g.Start(s)
}
