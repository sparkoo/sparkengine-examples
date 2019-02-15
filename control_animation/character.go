package main

import (
	"github.com/sparkoo/sparkengine/core/scene"
	"github.com/sparkoo/sparkengine/core/scene/image"
)

var pixels map[int][][]scene.Pixel

const (
	SPRITE_WIDTH  = 32
	SPRITE_HEIGHT = 32

	IDLE     = 1
	RUN      = 2
	JUMP     = 3
	ATTACK_1 = 4

	//characterFile = "BronzeKnight.png"
	characterFile = "Peasant.png"
)

var frameCounts = map[int]int{
	IDLE:     10,
	RUN:      10,
	ATTACK_1: 10,
}

type Character struct {
	*scene.Base

	currentAnimation int
	currentFrame     int

	moving bool
}

func init() {
	pixels = make(map[int][][]scene.Pixel, 3)

	// load idle
	pixels[IDLE] = make([][]scene.Pixel, frameCounts[IDLE])
	for i := 0; i < frameCounts[IDLE]; i++ {
		pixels[IDLE][i] = image.LoadImage("assets/"+characterFile, 32*i, 0, 32, 32)
	}

	// load running
	pixels[RUN] = make([][]scene.Pixel, frameCounts[RUN])
	for i := 0; i < frameCounts[RUN]; i++ {
		pixels[RUN][i] = image.LoadImage("assets/"+characterFile, 32*i, 64, 32, 32)
	}

	// load running
	pixels[ATTACK_1] = make([][]scene.Pixel, frameCounts[ATTACK_1])
	for i := 0; i < frameCounts[ATTACK_1]; i++ {
		pixels[ATTACK_1][i] = image.LoadImage("assets/"+characterFile, 32*i, 96, 32, 32)
	}
}

func NewCharacter(x float64, y float64) *Character {
	return &Character{
		Base:             scene.NewBase(x, y, SPRITE_WIDTH, SPRITE_HEIGHT),
		currentFrame:     0,
		currentAnimation: IDLE}
}

func (ch *Character) GetPixels() []scene.Pixel {
	return pixels[ch.currentAnimation][ch.currentFrame]
}

func (ch *Character) Tick() {
	if ch.currentFrame+1 >= frameCounts[ch.currentAnimation] {
		ch.currentFrame = 0
	}

	ch.currentFrame += 1

	if ch.moving {
		ch.MoveBy(1, 0)
	}
}

func (ch *Character) Moving(moving bool) {
	ch.moving = moving
}

func (ch *Character) ChangeAction(action int) {
	if ch.currentAnimation != action {
		ch.currentFrame = 0
		ch.currentAnimation = action
	}
}
