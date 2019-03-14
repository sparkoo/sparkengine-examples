package main

import (
	"math/rand"

	"github.com/sparkoo/sparkengine/core"
	"github.com/sparkoo/sparkengine/core/scene"
	"github.com/sparkoo/sparkengine/core/scene/olib/shape"
)

const screenWidth = 640
const screenHeight = 480

func main() {
	g := core.NewGame(core.NewConf30T640W())

	balls := generateRandomBalls(100)

	s := scene.NewScene(func(gameTickCounter int64, sceneTickCounter int64) {
		for _, b := range balls {
			x, y := b.GetPos()

			xpot := int(x + b.vx)
			if xpot < 0 || xpot+b.GetXsize() > screenWidth {
				b.vx = b.vx * -1
			}

			ypot := int(y + b.vy)
			if ypot < 0 || ypot+b.GetYsize() > screenHeight {
				b.vy = b.vy * -1
			}
			b.MoveBy(b.vx, b.vy)
		}
	})

	for _, b := range balls {
		s.AddObject(b)
	}

	g.Start(s)
}

func generateRandomBalls(n int) []movingBall {
	balls := make([]movingBall, n)
	for i := 0; i < n; i++ {
		balls[i] = movingBall{
			Ball: shape.NewBall(320, 240, rand.Intn(30), scene.RandomColor(255)),
			vx:   (rand.Float64() - .5) * 5,
			vy:   (rand.Float64() - .5) * 5}
	}
	return balls
}

type movingBall struct {
	*shape.Ball
	vx float64
	vy float64
}
