package state

import rl "github.com/gen2brain/raylib-go/raylib"

var COLOR_UNSELECT = rl.DarkGreen
var COLOR_SELECT = rl.Green
var COLOR_TEXT_SELECT = rl.White
var COLOR_TEXT_UNSELECT = rl.Gray

const (
	// Bottom Buttons
	KEY_F1 = rl.KeyF1
	KEY_F2 = rl.KeyF2
	KEY_F3 = rl.KeyF3
	KEY_F4 = rl.KeyF4
	KEY_F5 = rl.KeyF5

	// Left	 Side Buttons
	KEY_F6  = rl.KeyUp
	KEY_F7  = rl.KeyDown
	KEY_F8  = rl.KeyF6
	KEY_F9  = rl.KeyF7
	KEY_F10 = rl.KeyF8

	// Right Side Buttons
	KEY_F16 = rl.KeyF9
	KEY_F17 = rl.KeyF10
	KEY_F18 = rl.KeyF12
	KEY_F19 = rl.KeyF10
	KEY_F20 = rl.KeyF10

	// Top Buttons
	KEY_F11 = rl.KeyLeft
	KEY_F12 = rl.KeyRight
	KEY_F13 = rl.KeyF9
	KEY_F14 = rl.KeyF10
	KEY_F15 = rl.KeyF11
)

const MENU_BUTTON_HEIGHT = 36
const MENU_BUTTON_WIDTH = 64
const MENU_BUTTON_FONT_SIZE = 20

const CORNER_MARGIN = 80
const SCREEN_MARGIN = 16

const BUTTON_LENGTH = 5
