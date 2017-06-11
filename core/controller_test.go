package core

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type testKeyCallbackManager struct {
	registered bool
}

func (tkcm *testKeyCallbackManager) AddKeyCallback(KeyCallback) error {
	tkcm.registered = true
	return nil
}

func TestController(t *testing.T) {
	Convey("Creating a new KeyboardController", t, func() {
		kcm := &testKeyCallbackManager{}
		kc := NewKeyboardController(kcm)
		So(kc, ShouldNotBeNil)

		Convey("It registers a new keycallback listener", func() {
			So(kcm.registered, ShouldBeTrue)
		})
	})
}
