package main

import (
	"github.com/sparkoo/sparkengine/core/scene"
	"github.com/sparkoo/sparkengine/core/scene/image"
)

var animationPixels [8][]scene.Pixel

type Obj struct {
	*scene.Base

	currentAnimationFrame int
}

func NewObj(x float64, y float64, w int, h int) *Obj {
	return &Obj{Base: scene.NewBase(x, y, w, h), currentAnimationFrame: 0}
}

func init() {
	for i := range animationPixels {
		animationPixels[i] = image.LoadImage("loading.png", 0, 100*i, 50, 100)
	}
}

func (o *Obj) GetPixels() []scene.Pixel {
	return animationPixels[o.currentAnimationFrame]
}

func (o *Obj) TickAnimation() {
	if o.currentAnimationFrame+1 == len(animationPixels) {
		o.currentAnimationFrame = 0
	} else {
		o.currentAnimationFrame++
	}
}
