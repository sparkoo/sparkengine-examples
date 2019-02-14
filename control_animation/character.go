package main

import (
	"fmt"
	"github.com/sparkoo/sparkengine/core/scene"
	"github.com/sparkoo/sparkengine/core/scene/image"
)

var pixels map[int][][]scene.Pixel

const (
	SPRITE_WIDTH  = 315
	SPRITE_HEIGHT = 320

	IDLE     = 1
	RUN      = 2
	JUMP     = 3
	ATTACK_1 = 4
)

var frameCounts = map[int]int{
	IDLE:     10,
	RUN:      12,
	ATTACK_1: 16,
}

type Character struct {
	*scene.Base

	currentAnimation int
	currentFrame     int
}

func init() {
	pixels = make(map[int][][]scene.Pixel, 3)

	// load idle
	pixels[IDLE] = make([][]scene.Pixel, frameCounts[IDLE])
	for i := 0; i < frameCounts[IDLE]; i++ {
		pixels[IDLE][i] = image.LoadFullImage(fmt.Sprintf("assets/Idle/%02d.png", i))
	}

	// load running
	pixels[RUN] = make([][]scene.Pixel, frameCounts[RUN])
	for i := 1; i <= frameCounts[RUN]; i++ {
		pixels[RUN][i-1] = image.LoadFullImage(fmt.Sprintf("assets/Running/Run_%d.png", i))
	}

	// load running
	pixels[ATTACK_1] = make([][]scene.Pixel, frameCounts[ATTACK_1])
	for i := 0; i < frameCounts[ATTACK_1]; i++ {
		pixels[ATTACK_1][i] = image.LoadFullImage(fmt.Sprintf("assets/Attack 1/Attack_%03d.png", i))
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
}

func (ch *Character) ChangeAction(action int) {
	if ch.currentAnimation != action {
		ch.currentFrame = 0
		ch.currentAnimation = action
	}
}
