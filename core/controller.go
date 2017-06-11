package core

import (
	"errors"

	"github.com/goxjs/glfw"
)

type AbstractKey int

type Controller interface {
	IsDown(AbstractKey) bool
}

type keyboardController struct{}

func (kc *keyboardController) IsDown(ak AbstractKey) bool {
	return false
}

func NewKeyboardController(kcm KeyCallbackManager, mapping KeyboardMapping) (Controller, error) {
	if kcm == nil || mapping == nil {
		return nil, errors.New("Can't pass 'nil' to NewKeyboardController")
	}
	kc := &keyboardController{}

	kcm.AddKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
		return false
	})
	return kc, nil

}

// KeyMapping maps a key on the keyboard to a logical key
type KeyMapping struct {
	Abstract AbstractKey
	Key      glfw.Key
}

// KeyboardMapping is a collection of KeyMappings
type KeyboardMapping []KeyMapping
