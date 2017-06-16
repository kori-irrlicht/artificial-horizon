package core

import (
	"errors"

	"github.com/go-gl/glfw/v3.2/glfw"
)

type KeyCallback func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool

type SetKeyCallbacker interface {
	SetKeyCallback(glfw.KeyCallback) glfw.KeyCallback
}

type KeyCallbackManager interface {
	AddKeyCallback(KeyCallback) error
}
type keyCallbackManager struct {
	keyCallbacks []KeyCallback
}

// AddKeyCallback adds a KeyCallback
// Returns an error if kb is nil
func (kc *keyCallbackManager) AddKeyCallback(kb KeyCallback) (err error) {
	if kb == nil {
		return errors.New("Can't pass 'nil' to AddKeyCallback")
	}
	kc.keyCallbacks = append(kc.keyCallbacks, kb)
	return
}

// InitKeyCallback adds a keycallback to the kb parameter
// This callback calls every callback registered with AddKeyCallback until one
// handles the event and returns true
func NewKeyCallbackManager(kb SetKeyCallbacker) (KeyCallbackManager, error) {
	if kb == nil {
		return nil, errors.New("Can't pass 'nil' to NewKeyCallback")
	}
	kc := &keyCallbackManager{}
	kb.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		for _, cb := range kc.keyCallbacks {
			if cb(w, key, scancode, action, mods) {
				return
			}
		}
	})
	return kc, nil
}
