package main

import (
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/kori-irrlicht/artificial-dream/core"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	l := logrus.New()
	l.Formatter = &logrus.JSONFormatter{}

	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	if err = gl.Init(); err != nil {
		panic(err)
	}

	g, err := newGame()
	if err != nil {
		panic(err)
	}
	core.GameLoop(g)

}
