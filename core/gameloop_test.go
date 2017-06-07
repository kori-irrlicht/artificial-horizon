package core

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

type game struct {
	isInput             int
	isUpdate            int
	isUpdateAfterInput  bool
	isRender            int
	isRenderAfterUpdate bool
	isRunning           chan bool
	time                chan time.Time
}

func (g *game) Input() {
	g.isInput++
	g.isUpdateAfterInput = false
}

func (g *game) Update() {
	g.isUpdate++
	g.isUpdateAfterInput = true
	g.isRenderAfterUpdate = false
}

func (g *game) Render() {
	g.isRender++
	g.isRenderAfterUpdate = true
}

func (g *game) Running() bool {
	return <-g.isRunning
}

func (g *game) Now() time.Time {
	return <-g.time
}

func (g *game) FrameTime() time.Duration {
	return 16 * time.Millisecond
}

func TestGameLoop(t *testing.T) {

	Convey("Main is called", t, func() {
		g := &game{}
		g.isRunning = make(chan bool, 4)
		g.time = make(chan time.Time, 4)
		Convey("Game is not running", func() {
			g.isRunning <- false
			g.time <- time.Unix(0, 0)
			GameLoop(g)
			Convey("Input is not run", func() {
				So(g.isInput, ShouldEqual, 0)
			})
			Convey("Update is not run", func() {
				So(g.isUpdate, ShouldEqual, 0)

			})
			Convey("Render is not run", func() {
				So(g.isRender, ShouldEqual, 0)
			})

		})

		Convey("Game is running", func() {
			g.isRunning <- true
			g.isRunning <- false
			g.time <- time.Unix(0, 0)
			g.time <- time.Unix(0, 16000000)
			GameLoop(g)
			Convey("Input is run", func() {
				So(g.isInput, ShouldEqual, 1)
			})
			Convey("Update is run", func() {
				So(g.isUpdate, ShouldEqual, 1)
				Convey("After Input", func() {
					So(g.isUpdateAfterInput, ShouldBeTrue)
				})

			})
			Convey("Render is run", func() {
				So(g.isRender, ShouldEqual, 1)
				Convey("After Update", func() {
					So(g.isRenderAfterUpdate, ShouldBeTrue)
				})
			})
		})

		Convey("Game is running 3 times", func() {
			g.isRunning <- true
			g.isRunning <- true
			g.isRunning <- true
			g.isRunning <- false
			g.time <- time.Unix(0, 0)
			g.time <- time.Unix(0, 16000000)
			g.time <- time.Unix(0, 16000000*2)
			g.time <- time.Unix(0, 16000000*3)
			GameLoop(g)
			Convey("Input is run 3 times", func() {
				So(g.isInput, ShouldEqual, 3)
			})
			Convey("Update is run 3 times", func() {
				So(g.isUpdate, ShouldEqual, 3)
			})
			Convey("Render is run 3 times", func() {
				So(g.isRender, ShouldEqual, 3)
			})
		})

		Convey("Last frame time took to long", func() {
			g.isRunning <- true
			g.isRunning <- false
			g.time <- time.Unix(0, 0)
			g.time <- time.Unix(0, 16*2*1000*1000)
			GameLoop(g)
			Convey("Input is run 1 times", func() {
				So(g.isInput, ShouldEqual, 1)
			})
			Convey("Update is run 2 times", func() {
				So(g.isUpdate, ShouldEqual, 2)
			})
			Convey("Render is run 1 times", func() {
				So(g.isRender, ShouldEqual, 1)
			})

		})

	})
}
