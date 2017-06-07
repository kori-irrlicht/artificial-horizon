package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/kori-irrlicht/artificial-dream/core"
)

func main() {
	l := logrus.New()
	l.Formatter = &logrus.JSONFormatter{}

	core.GameLoop(newGame())

}
