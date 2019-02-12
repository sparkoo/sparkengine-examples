package main

import (
	"github.com/sparkoo/sparkengine-examples/menu/conf"
	"github.com/sparkoo/sparkengine-examples/menu/objects"
	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/scene"
	"github.com/veandco/go-sdl2/sdl"
	"log"
)

var game *core.Game

var elM1 = []scene.Object{
	objects.NewElement(1),
	objects.NewElement(2),
	objects.NewElement(3),
	objects.NewElement(4),
}

var elM2 = []scene.Object{
	objects.NewElement(1),
	objects.NewElement(2),
}

var cursorM1 = objects.NewCursor(0, len(elM1))
var cursorM2 = objects.NewCursor(0, len(elM2))

var menu1 = scene.NewScene(scene.NoopTick)
var menu2 = scene.NewScene(scene.NoopTick)

func init() {
	game = core.NewGame(core.NewConf(conf.FPS, conf.FPS*2, conf.SWIDTH, conf.SHEIGHT))

	menu1.AddObjects(elM1...)
	menu1.AddObject(cursorM1)
	menu1.AddEventListener(func(event sdl.Event) {
		menuMoveEvent(cursorM1, event)
	})
	menu1.AddEventListener(menu1Enter)

	menu2.AddObjects(elM2...)
	menu2.AddObject(cursorM2)
	menu2.AddEventListener(func(event sdl.Event) {
		menuMoveEvent(cursorM2, event)
	})
	menu2.AddEventListener(menu2Enter)
}

func main() {
	game.Start(menu1)
}

func menuMoveEvent(c *objects.Cursor, event sdl.Event) {
	switch t := event.(type) {
	case *sdl.KeyboardEvent:
		if t.Keysym.Scancode == 82 && t.State == 1 { // up
			moveCursor(c, -1)
		} else if t.Keysym.Scancode == 81 && t.State == 1 { // down
			moveCursor(c, 1)
		}
	}
}

func moveCursor(c *objects.Cursor, diff int) {
	potPos := c.Pos + diff
	if potPos < 0 {
		potPos = 0
	} else if potPos >= c.Fields-1 {
		potPos = c.Fields - 1
	}
	c.Pos = potPos
	c.MoveToPos(c.Pos)
}

func menu1Enter(event sdl.Event) {
	switch t := event.(type) {
	case *sdl.KeyboardEvent:
		if t.Keysym.Scancode == 40 && t.State == 1 { // enter
			if cursorM1.Pos >= len(elM1)-1 {
				game.Quit()
			} else if cursorM1.Pos == 0 {
				game.SwitchScene(menu2, false)
			}
		}
	}
}

func menu2Enter(event sdl.Event) {
	switch t := event.(type) {
	case *sdl.KeyboardEvent:
		if t.Keysym.Scancode == 40 && t.State == 1{ // enter
			log.Println(cursorM2.Pos, cursorM2.Fields)
			if cursorM2.Pos >= cursorM2.Fields-1 {
				game.SwitchScene(menu1, false)
			}
		}
	}
}
