package main

import (
	"time"

	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/kori-irrlicht/artificial-dream/core"
)

var _ core.Game = &game{}

type game struct {
	window *glfw.Window
}

func (g *game) Update() {}
func (g *game) Input()  {}
func (g *game) Render() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	g.window.SwapBuffers()
	glfw.PollEvents()
}
func (g *game) FrameTime() time.Duration {
	return 16 * time.Millisecond
}

func (g *game) Now() time.Time {
	return time.Now()
}

func (g *game) Running() bool {
	return !g.window.ShouldClose()
}

func newGame() (core.Game, error) {
	game := &game{}

	vidMode := glfw.GetPrimaryMonitor().GetVideoMode()

	window, err := glfw.CreateWindow(vidMode.Width, vidMode.Height, "Artificial Horizon", glfw.GetPrimaryMonitor(), nil)
	if err != nil {
		return nil, err
	}
	window.MakeContextCurrent()
	gl.ClearColor(0.8, 0.3, 0.01, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	game.window = window
	return game, nil
}
