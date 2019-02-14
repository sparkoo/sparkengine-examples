package main

import (
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/core/scene"
	"github.com/sparkoo/sparkengine/core/sound"
	"github.com/veandco/go-sdl2/sdl"
)

var snd *sound.Sound

func main() {
	g := core.NewGame(core.NewConf30T320W())

	snd = sound.LoadSound("foom_0.wav")

	s := scene.NewScene(scene.NoopTick)

	s.AddEventListener(func(event sdl.Event) {
		switch e := event.(type) {
		case *sdl.KeyboardEvent:
			if e.Keysym.Scancode == 82 && e.State == 1 { // up
				snd.Play(0)
			}
		}
	})

	g.Start(s)
}
