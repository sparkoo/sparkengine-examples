package main

import (
	"fmt"
	"github.com/sparkoo/sparkengine/scene"
	"github.com/sparkoo/sparkengine/scene/image"
)

var vikingPixels [60][]scene.Pixel

// sprites taken from https://opengameart.org/content/viking-sprite
type Viking struct {
	*scene.Base

	currentAnimationFrame int
}

func NewViking(x float64, y float64, w int, h int) *Viking {
	return &Viking{Base: scene.NewBase(x, y, w, h), currentAnimationFrame: 0}
}

func init() {
	for i := range vikingPixels {
		vikingPixels[i] = image.LoadFullImage(fmt.Sprintf("viking_Sprites/viking_idle%04d.png", i + 1))
	}
}

func (o *Viking) GetPixels() []scene.Pixel {
	return vikingPixels[o.currentAnimationFrame]
}

func (o *Viking) TickAnimation() {
	if o.currentAnimationFrame+1 == len(vikingPixels) {
		o.currentAnimationFrame = 0
	} else {
		o.currentAnimationFrame++
	}
}
