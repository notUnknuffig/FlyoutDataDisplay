package state

import rl "github.com/gen2brain/raylib-go/raylib"

var COLOR_UNSELECT = rl.DarkGreen
var COLOR_SELECT = rl.Green
var COLOR_TEXT_SELECT = rl.White
var COLOR_TEXT_UNSELECT = rl.Gray

var (
	// Bottom Buttons
	F1 = rl.KeyF1
	F2 = rl.KeyF2
	F3 = rl.KeyF3
	F4 = rl.KeyF4
	F5 = rl.KeyF5

	// Side Buttons
	F6  = rl.KeyUp
	F7  = rl.KeyDown
	F8  = rl.KeyF6
	F9  = rl.KeyF7
	F10 = rl.KeyF8

	// Top Buttons
	F11 = rl.KeyF9
	F12 = rl.KeyF10
	F13 = rl.KeyF12
	F14 = rl.KeyF10
	F15 = rl.KeyF10
)

const MENU_BUTTON_SIZE = 36

const BUTTON_LENGTH = 5

const SCREEN_MARGIN = 16

const FONT_SIZE = 30
