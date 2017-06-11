package core

import (
	"testing"

	"github.com/goxjs/glfw"
	. "github.com/smartystreets/goconvey/convey"
)

type testKeyCallbackManager struct {
	registered bool
}

func (tkcm *testKeyCallbackManager) AddKeyCallback(KeyCallback) error {
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
