package main

import (
	"fmt"
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

func newGame() core.Game {
	game := &game{}

	err := glfw.Init(gl.ContextWatcher)
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()
	fmt.Printf("OpenGL: %s %s %s; %v samples.\n", gl.GetString(gl.VENDOR), gl.GetString(gl.RENDERER), gl.GetString(gl.VERSION), gl.GetInteger(gl.SAMPLES))
	fmt.Printf("GLSL: %s.\n", gl.GetString(gl.SHADING_LANGUAGE_VERSION))

	vidMode := glfw.GetPrimaryMonitor().GetVideoMode()

	window, err := glfw.CreateWindow(vidMode.Width, vidMode.Height, "Artificial Horizon", glfw.GetPrimaryMonitor(), nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	gl.ClearColor(0.8, 0.3, 0.01, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	game.window = window
	return game
}
