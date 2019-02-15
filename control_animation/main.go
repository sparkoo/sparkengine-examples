package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/core/scene"
	"github.com/sparkoo/sparkengine/core/scene/image"
	"github.com/sparkoo/sparkengine/core/scene/olib/sprite"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	g := core.NewGame(core.NewConf30T960W())

	character := NewCharacter(0, 400)
	background := sprite.NewSprite(0, 0, 960, 540, func() []scene.Pixel {
		return image.LoadFullImage("assets/background.png")
	})

	s := scene.NewScene(func(gameTickCounter int64, sceneTickCounter int64) {
			character.Tick()
	})

	s.AddEventListener(func(event sdl.Event) {
		switch e := event.(type) {
		case *sdl.KeyboardEvent:
			//log.Println(e.Keysym.Scancode)

			if e.Keysym.Scancode == 79 && e.State == 1 {
				character.ChangeAction(RUN)
				character.Moving(true)
			} else if e.Keysym.Scancode == 79 && e.State == 0 {
				character.ChangeAction(IDLE)
				character.Moving(false)
			}
			if e.Keysym.Scancode == 224 && e.State == 1 {
				character.ChangeAction(ATTACK_1)
			} else if e.Keysym.Scancode == 224 && e.State == 0 {
				character.ChangeAction(IDLE)
			}
		}
	})

	s.AddObjects(background, character)

	g.Start(s)
}
