package core

import (
	"testing"

	"github.com/goxjs/glfw"
	. "github.com/smartystreets/goconvey/convey"
)

func callback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
	return false
}

type callBacker struct {
	callback glfw.KeyCallback
}

func (cb *callBacker) SetKeyCallback(kb glfw.KeyCallback) {
	cb.callback = kb
}

func TestCallback(t *testing.T) {
	Convey("Clear callback array", t, func() {
		keyCallbacks = make([]KeyCallback, 0)
		Convey("Adding a callback", func() {
			err := AddKeyCallback(callback)
			Convey("keyCallbacks should contain the function", func() {
				So(len(keyCallbacks), ShouldEqual, 1)
				So(keyCallbacks[0], ShouldEqual, callback)
				So(err, ShouldBeNil)
			})
		})
		Convey("Adding nil to callback", func() {
			err := AddKeyCallback(nil)
			Convey("keyCallbacks should be unchanged", func() {
				So(len(keyCallbacks), ShouldEqual, 0)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("Initializing a new SetKeyCallbacker", func() {
			cb := &callBacker{}
			InitKeyCallback(cb)
			So(cb.callback, ShouldNotBeNil)
			Convey("Adding a callback", func() {
				called := false
				AddKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
					called = true
					return false
				})

				Convey("And calling it", func() {
					cb.callback(nil, 0, 0, 0, 0)
					So(called, ShouldBeTrue)
				})

			})
			Convey("Adding two callbacks, first true, second false", func() {
				amount := 0
				AddKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
					amount++
					return true
				})
				AddKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
					amount++
					return false
				})

				Convey("Only one is called", func() {
					cb.callback(nil, 0, 0, 0, 0)
					So(amount, ShouldEqual, 1)
				})

			})
			Convey("Adding two callbacks, first false, second true", func() {
				amount := 0
				AddKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
					amount++
					return false
				})
				AddKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
					amount++
					return true
				})

				Convey("Both are called", func() {
					cb.callback(nil, 0, 0, 0, 0)
					So(amount, ShouldEqual, 2)
				})

			})
		})

	})

}
