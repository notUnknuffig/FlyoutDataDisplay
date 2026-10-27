package state

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var ConnectionInputProcessed = false

var mapping []int32 = []int32{
	// Bottom Buttons
	rl.KeyF1,
	rl.KeyF2,
	rl.KeyF3,
	rl.KeyF4,
	rl.KeyF5,

	// Left	 Side Buttons
	rl.KeyUp,
	rl.KeyDown,
	rl.KeyF6,
	rl.KeyF7,
	rl.KeyF8,

	// Top Buttons
	rl.KeyLeft,
	rl.KeyRight,
	rl.KeyF9,
	rl.KeyF10,
	rl.KeyF11,

	// Right Side Buttons
	rl.KeyF9,
	rl.KeyF10,
	rl.KeyF12,
	rl.KeyF10,
	rl.KeyF10,
}

func IsKeyPressed(key int) bool {
	bit := key - 1
	if GlobalFlightData == nil {
		return rl.IsKeyPressed(mapping[bit])
	} else if GlobalFlightData.Input.FunctionKeyMask&(1<<bit) != 0 && !ConnectionInputProcessed {
		ConnectionInputProcessed = true
		return true
	} else {
		return rl.IsKeyPressed(mapping[bit])
	}
}
