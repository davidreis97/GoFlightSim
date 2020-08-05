package controller

import (
	"github.com/g3n/engine/window"
	"github.com/g3n/engine/core"
)

// Keyboard holds the KeyState of the keyboard
type Keyboard struct {
	keyState *window.KeyState
}

// Input stores the inputs given by the user
type Input struct {
	Thrust float32
	Pitch float32
	Yaw float32
	Roll float32
}

// InitKeyboard initializes the Keyboard controler for the given window (core.IDispatcher)
func InitKeyboard(win core.IDispatcher) Keyboard {
	return Keyboard{keyState: window.NewKeyState(win)}
}

// ProcessInput generates an Input object according to the keys pressed
func (k Keyboard) ProcessInput() Input {
	input := Input{0,0,0,0}

	// Thrust
	if k.keyState.Pressed(window.KeyA){
		input.Thrust++
	}
	if k.keyState.Pressed(window.KeyZ){
		input.Thrust--
	}

	// Pitch
	if k.keyState.Pressed(window.KeyUp){
		input.Pitch++
	}
	if k.keyState.Pressed(window.KeyDown){
		input.Pitch--
	}

	// Yaw
	if k.keyState.Pressed(window.KeyE){
		input.Yaw++
	}
	if k.keyState.Pressed(window.KeyQ){
		input.Yaw--
	}

	// Roll
	if k.keyState.Pressed(window.KeyRight){
		input.Roll++
	}
	if k.keyState.Pressed(window.KeyLeft){
		input.Roll--
	}

	return input
}