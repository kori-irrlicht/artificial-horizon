package main

import (
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/kori-irrlicht/artificial-dream/core"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	l := logrus.New()
	l.Formatter = &logrus.JSONFormatter{}

	err := glfw.Init(gl.ContextWatcher)
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()
	g, err := newGame()
	if err != nil {
		panic(err)
	}
	core.GameLoop(g)

}
