package core

import "github.com/goxjs/glfw"

type AbstractKey int

type Controller interface {
	IsDown(AbstractKey) bool
}

type keyboardController struct{}

func (kc *keyboardController) IsDown(ak AbstractKey) bool {
	return false
}

func NewKeyboardController(kcm KeyCallbackManager) Controller {
	kc := &keyboardController{}

	kcm.AddKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) bool {
		return false
	})
	return kc

}
