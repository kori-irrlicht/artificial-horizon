package core

import (
	"errors"

	"github.com/goxjs/glfw"
)

type AbstractKey int

type Controller interface {
	IsDown(AbstractKey) bool
}

type keyboardController struct {
	keyState map[AbstractKey]bool
	mapping  KeyboardMapping
}

func (kc *keyboardController) IsDown(ak AbstractKey) bool {
	return kc.keyState[ak]
}

func NewKeyboardController(kcm KeyCallbackManager, mapping KeyboardMapping) (Controller, error) {
	if kcm == nil || mapping == nil {
		return nil, errors.New("Can't pass 'nil' to NewKeyboardController")
	}
	kc := &keyboardController{}
	kc.mapping = mapping
	kc.keyState = make(map[AbstractKey]bool)

	kcm.AddKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
		handled := false
		if action == glfw.Press {
			for _, v := range kc.mapping.Find(key) {
				kc.keyState[v] = true
				handled = true

			}
		}
		if action == glfw.Release {
			for _, v := range kc.mapping.Find(key) {
				kc.keyState[v] = false
				handled = true

			}
		}
		return handled
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

// Find finds all abstract keys, key is mapped to
func (km KeyboardMapping) Find(key glfw.Key) (res []AbstractKey) {
	for _, v := range km {
		if v.Key == key {
			res = append(res, v.Abstract)
		}
	}
	return

}
