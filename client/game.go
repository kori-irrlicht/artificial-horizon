package main

import (
	"fmt"
	"time"

	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/kori-irrlicht/artificial-horizon/core"
	"github.com/kori-irrlicht/artificial-horizon/network"
)

var _ core.Game = &game{}

type game struct {
	window     *glfw.Window
	controller core.Controller
}

func (g *game) Update() {}
func (g *game) Input() {
	if g.controller.IsDown(KeyDown) {
		fmt.Println("Down")
	}
	if g.controller.IsDown(KeyUp) {
		fmt.Println("Up")
	}
}
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

	kcm, _ := core.NewKeyCallbackManager(window)
	mapping := core.KeyboardMapping{
		{KeyUp, glfw.KeyW},
		{KeyDown, glfw.KeyS},
		{KeyLeft, glfw.KeyA},
		{KeyRight, glfw.KeyD},
	}
	game.controller, _ = core.NewKeyboardController(kcm, mapping)

	conn, err := network.Dial("127.0.0.1:42424", "http://localhost")
	if err != nil {
		return nil, err
	}

	conn.Write([]byte(`Hallo`))

	return game, nil
}
