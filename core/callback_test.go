package core

import (
	"testing"

	"github.com/go-gl/glfw/v3.2/glfw"
	. "github.com/smartystreets/goconvey/convey"
)

func callback(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
	return false
}

type callBacker struct {
	callback glfw.KeyCallback
}

func (cb *callBacker) SetKeyCallback(kb glfw.KeyCallback) glfw.KeyCallback {
	cb.callback = kb
	return nil
}

func TestCallback(t *testing.T) {
	Convey("Init Keycallback with nil ", t, func() {
		kc, err := NewKeyCallbackManager(nil)
		Convey("Should return nil and an error", func() {
			So(kc, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
	})

	Convey("Init Keycallback with valid parameter", t, func() {
		cb := &callBacker{}
		kcm, err := NewKeyCallbackManager(cb)
		kc := kcm.(*keyCallbackManager)
		Convey("Should create new manager and return no error", func() {
			So(kc, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
		Convey("Adding a callback", func() {
			err := kc.AddKeyCallback(callback)
			Convey("keyCallbacks should contain the function", func() {
				So(len(kc.keyCallbacks), ShouldEqual, 1)
				So(kc.keyCallbacks[0], ShouldEqual, callback)
				So(err, ShouldBeNil)
			})
		})
		Convey("Adding nil to callback", func() {
			err := kc.AddKeyCallback(nil)
			Convey("keyCallbacks should be unchanged", func() {
				So(len(kc.keyCallbacks), ShouldEqual, 0)
				So(err, ShouldNotBeNil)
			})
		})

		Convey("Initializing a new SetKeyCallbacker", func() {
			So(cb.callback, ShouldNotBeNil)
			Convey("Adding a callback", func() {
				called := false
				kc.AddKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
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
				kc.AddKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
					amount++
					return true
				})
				kc.AddKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
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
				kc.AddKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
					amount++
					return false
				})
				kc.AddKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
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
