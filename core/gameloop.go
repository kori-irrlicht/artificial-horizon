package core

import "time"

// GameLoop implements a fixed timestep gameloop
func GameLoop(g Game) {
	dt := g.FrameTime()

	currentTime := g.Now()
	acc := 0 * time.Millisecond

	for g.Running() {
		newTime := g.Now()
		diff := newTime.Sub(currentTime)
		currentTime = newTime

		acc += diff
		g.Input()
		for acc >= dt {
			g.Update()
			acc -= dt
		}

		g.Render()
	}

}

// Game contains the game logic
type Game interface {
	Update()
	Input()
	Render()
	Running() bool
	Now() time.Time
	FrameTime() time.Duration
}
