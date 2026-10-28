package state

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	if GlobalInputBuffer == 0 {
		return rl.IsKeyPressed(mapping[bit])
	} else if GlobalInputBuffer&(1<<bit) != 0 {
		GlobalInputBuffer ^= (1 << bit)
		return true
	} else {
		return rl.IsKeyPressed(mapping[bit])
	}
}
