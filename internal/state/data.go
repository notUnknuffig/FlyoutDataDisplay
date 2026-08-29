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

func DrawNoData() {
	var messageWidth = Scale(156)
	var messageHeight = Scale(52)
	var messageMargin = Scale(12)
	var centerBoxX = (int32(rl.GetRenderWidth()) - messageWidth) / 2
	var centerBoxY = (int32(rl.GetRenderHeight()) - messageHeight) / 2
	rl.DrawRectangle(centerBoxX, centerBoxY, int32(messageWidth), int32(messageHeight), COLOR_UNSELECT)
	rl.DrawText("No Data", centerBoxX+int32(messageMargin), centerBoxY+int32(messageMargin), Scale(32), COLOR_SELECT)

}

// Returns Achor Y coord below up button.
func DrawArrowButtons(i int) (int32, int32) {
	var buttonSize = Scale(MENU_BUTTON_SIZE) // Square
	var baseX = SCREEN_MARGIN
	var baseY = SCREEN_MARGIN + ((rl.GetRenderHeight()-(2*SCREEN_MARGIN))/(BUTTON_LENGTH+1))*(i+1) - int(buttonSize/2)
	var anchorY = int32(baseY) + buttonSize
	var diff = int32(0)
	d := rl.Vector2{X: float32(baseX + int(buttonSize)), Y: float32(baseY + int(buttonSize))}
	e := rl.Vector2{X: float32(baseX), Y: float32(baseY + int(buttonSize))}
	f := rl.Vector2{X: float32(baseX + int(buttonSize/2)), Y: float32(baseY)}
	rl.DrawTriangle(f, e, d, COLOR_SELECT)
	rl.DrawTriangle(rl.Vector2Add(f, rl.Vector2{0, ScaleF(7)}), rl.Vector2Add(e, rl.Vector2{ScaleF(6), ScaleF(-3)}), rl.Vector2Add(d, rl.Vector2{ScaleF(-6), ScaleF(-3)}), rl.Black)

	baseX = SCREEN_MARGIN
	baseY = SCREEN_MARGIN + ((rl.GetRenderHeight()-(2*SCREEN_MARGIN))/(BUTTON_LENGTH+1))*(i+1+1) - int(buttonSize/2)
	diff = int32(baseY) - anchorY
	a := rl.Vector2{X: float32(baseX), Y: float32(baseY)}
	b := rl.Vector2{X: float32(baseX + int(buttonSize)), Y: float32(baseY)}
	c := rl.Vector2{X: float32(baseX + int(buttonSize/2)), Y: float32(baseY + int(buttonSize))}
	rl.DrawTriangle(c, b, a, COLOR_SELECT)
	rl.DrawTriangle(rl.Vector2Add(c, rl.Vector2{0, ScaleF(-7)}), rl.Vector2Add(b, rl.Vector2{ScaleF(-6), ScaleF(3)}), rl.Vector2Add(a, rl.Vector2{ScaleF(6), ScaleF(3)}), rl.Black)
	return anchorY, diff
}

func Scale(a int) int32 {
	return int32(float32(a) * GlobalOptions.Scale)
}

func ScaleF(a float32) float32 {
	return a * GlobalOptions.Scale
}
