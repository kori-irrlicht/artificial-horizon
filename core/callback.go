package core

import (
	"errors"

	"github.com/goxjs/glfw"
)

type KeyCallback func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool

type SetKeyCallbacker interface {
	SetKeyCallback(glfw.KeyCallback)
}

var keyCallbacks []KeyCallback

// AddKeyCallback adds a KeyCallback
// Returns an error if kb is nil
func AddKeyCallback(kb KeyCallback) (err error) {
	if kb == nil {
		return errors.New("Can't pass 'nil' to AddKeyCallback")
	}
	keyCallbacks = append(keyCallbacks, kb)
	return
}

// InitKeyCallback adds a keycallback to the kb parameter
// This callback calls every callback registered with AddKeyCallback until one
// handles the event and returns true
func InitKeyCallback(kb SetKeyCallbacker) {
	kb.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		for _, cb := range keyCallbacks {
			if cb(w, key, scancode, action, mods) {
				return
			}
		}
	})
}
