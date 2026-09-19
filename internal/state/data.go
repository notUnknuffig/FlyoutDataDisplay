package state

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var GlobalFlightData *FlightData
var SmoothedGlobalFlightData *FlightData
var GlobalOptions *Options

func SmoothData(ms int) {
	//
}

// Returns the top right corner of a button around the display.
func GetButtonAnchor(index int) (int32, int32) {
	var x, y int32
	corneredWidth := int32(rl.GetScreenWidth() - int(2*Scale(SCREEN_MARGIN+CORNER_MARGIN)))
	corneredHeight := int32(rl.GetScreenHeight() - int(2*Scale(SCREEN_MARGIN+CORNER_MARGIN)))
	if index < 5 { // Bottom Buttons
		y = int32(rl.GetRenderHeight()) - Scale(SCREEN_MARGIN+MENU_BUTTON_HEIGHT)
		x = Scale(CORNER_MARGIN+SCREEN_MARGIN) +
			((corneredWidth-Scale(MENU_BUTTON_WIDTH))/
				(BUTTON_LENGTH-1))*int32(index%5)
	} else if index < 10 { // Left Buttons
		y = Scale(CORNER_MARGIN+SCREEN_MARGIN) +
			((corneredHeight-Scale(MENU_BUTTON_WIDTH))/
				(BUTTON_LENGTH-1))*int32(index%5)
		x = Scale(SCREEN_MARGIN)
	} else if index < 15 { // Top Buttons
		y = Scale(SCREEN_MARGIN)
		x = Scale(CORNER_MARGIN+SCREEN_MARGIN) +
			((corneredWidth-Scale(MENU_BUTTON_WIDTH))/
				(BUTTON_LENGTH-1))*int32(index%5)
	} else { // Left Buttons
		y = Scale(CORNER_MARGIN+SCREEN_MARGIN) +
			((corneredHeight-Scale(MENU_BUTTON_WIDTH))/
				(BUTTON_LENGTH-1))*int32(index%5)
		x = int32(rl.GetRenderWidth()) - Scale(SCREEN_MARGIN+MENU_BUTTON_HEIGHT)
	}

	return x, y
}

func GetDisplayAreaWidth() int32 {
	return int32(rl.GetScreenWidth() - int(4*Scale(SCREEN_MARGIN)+2*Scale(MENU_BUTTON_HEIGHT)))
}

func GetDisplayAreaHeight() int32 {
	return int32(rl.GetScreenHeight() - int(4*Scale(SCREEN_MARGIN)+2*Scale(MENU_BUTTON_HEIGHT)))
}

func Scale(a int) int32 {
	return int32(float32(a) * GlobalOptions.Scale)
}

func ScaleF(a float32) float32 {
	return a * GlobalOptions.Scale
}
