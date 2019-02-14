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

	IDLE = 1
	RUN  = 2
	JUMP = 3
)

var frameCounts = map[int]int{
	IDLE: 10,
	RUN:  12,
}

type Character struct {
	*scene.Base

	currentAnimation   int
	currentFrame       int
	animationDirection int
}

func init() {
	pixels = make(map[int][][]scene.Pixel, 1)

	// load idle
	pixels[IDLE] = make([][]scene.Pixel, frameCounts[IDLE])
	for i := 0; i < frameCounts[IDLE]; i++ {
		pixels[IDLE][i] = image.LoadFullImage(fmt.Sprintf("assets/Idle/%02d.png", i))
	}
}

func NewCharacter(x float64, y float64) *Character {
	return &Character{
		Base:               scene.NewBase(x, y, SPRITE_WIDTH, SPRITE_HEIGHT),
		currentFrame:       0,
		currentAnimation:   IDLE,
		animationDirection: 1}
}

func (ch *Character) GetPixels() []scene.Pixel {
	//log.Println(ch.currentFrame)
	return pixels[ch.currentAnimation][ch.currentFrame]
}

func (ch *Character) Tick() {
	if ch.currentFrame+1 >= frameCounts[ch.currentAnimation] {
		ch.animationDirection = -1
	}
	if ch.currentFrame-1 < 0 {
		ch.animationDirection = 1
	}

	ch.currentFrame += ch.animationDirection
}
