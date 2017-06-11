package core

import (
	"testing"

	"github.com/goxjs/glfw"
	. "github.com/smartystreets/goconvey/convey"
)

type testKeyCallbackManager struct {
	kc         KeyCallback
	registered bool
}

func (tkcm *testKeyCallbackManager) AddKeyCallback(kc KeyCallback) error {
	tkcm.kc = kc
	tkcm.registered = true
	return nil
}

const testKey = 4

func TestController(t *testing.T) {
	Convey("Creating a new KeyboardController", t, func() {
		kcm := &testKeyCallbackManager{}
		mapping := KeyboardMapping{{
			testKey, glfw.Key4,
		}}
		kc, err := NewKeyboardController(kcm, mapping)
		So(kc, ShouldNotBeNil)
		So(err, ShouldBeNil)

		Convey("It registers a new keycallback listener", func() {
			So(kcm.registered, ShouldBeTrue)
			Convey("The key is not pressed", func() {
				So(kc.IsDown(testKey), ShouldBeFalse)
			})
			Convey("Pressing an unknown key", func() {
				handled := kcm.kc(nil, 0, 0, glfw.Press, 0)
				Convey("It shouldn't be handled", func() {
					So(kc.IsDown(testKey), ShouldBeFalse)
					So(handled, ShouldBeFalse)
				})
			})
			Convey("Releasing an unknown key", func() {
				handled := kcm.kc(nil, 0, 0, glfw.Release, 0)
				Convey("It shouldn't be handled", func() {
					So(handled, ShouldBeFalse)
				})

			})
			Convey("Pressing the registered key", func() {
				handled := kcm.kc(nil, glfw.Key4, 0, glfw.Press, 0)
				Convey("It should be handled", func() {
					So(kc.IsDown(testKey), ShouldBeTrue)
					So(handled, ShouldBeTrue)
				})
				Convey("Releasing the registered key", func() {
					handled := kcm.kc(nil, glfw.Key4, 0, glfw.Release, 0)
					Convey("It should be handled", func() {
						So(kc.IsDown(testKey), ShouldBeFalse)
						So(handled, ShouldBeTrue)
					})
				})
			})
		})
	})
	Convey("Creating a new KeyboardController with invalid input", t, func() {
		Convey("Manager and nil", func() {
			kc, err := NewKeyboardController(&testKeyCallbackManager{}, nil)
			So(kc, ShouldBeNil)
			So(err, ShouldNotBeNil)

		})
		Convey("nil and Keymapping", func() {
			kc, err := NewKeyboardController(nil, make(KeyboardMapping, 0))
			So(kc, ShouldBeNil)
			So(err, ShouldNotBeNil)

		})
		Convey("nil and nil", func() {
			kc, err := NewKeyboardController(nil, nil)
			So(kc, ShouldBeNil)
			So(err, ShouldNotBeNil)

		})

	})
}
