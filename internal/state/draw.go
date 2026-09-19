package state

import (
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawNoData() {
	var messageWidth = Scale(156)
	var messageHeight = Scale(52)
	var messageMargin = Scale(12)
	var centerBoxX = (int32(rl.GetRenderWidth()) - messageWidth) / 2
	var centerBoxY = (int32(rl.GetRenderHeight()) - messageHeight) / 2
	rl.DrawRectangle(centerBoxX, centerBoxY, int32(messageWidth), int32(messageHeight), COLOR_UNSELECT)
	rl.DrawText("No Data", centerBoxX+int32(messageMargin), centerBoxY+int32(messageMargin), Scale(32), COLOR_SELECT)
}

func DrawButton(text string, index int) {
	buttonWidth := Scale(MENU_BUTTON_WIDTH)
	buttonHeight := Scale(MENU_BUTTON_HEIGHT)
	buttonFont := Scale(MENU_BUTTON_FONT_SIZE)
	x, y := GetButtonAnchor(index)

	if (index < 10 && index >= 5) || (index < 15 && index >= 10) {
		rl.DrawRectangle(x, y, buttonHeight, buttonWidth, COLOR_UNSELECT)
		symbols := strings.Split(text, "")
		centerX := x + (buttonHeight / 2)
		for i := 0; i < len(symbols); i++ {
			offsetPercent := (-float32(len(symbols)+1))/2 + float32(i+1)
			centerY := y + (Scale(MENU_BUTTON_WIDTH) / 2) + int32((float32(buttonFont)*0.8)*offsetPercent)
			rl.DrawText(symbols[i], centerX-(rl.MeasureText(symbols[i], buttonFont))/2, centerY-(buttonFont/2), buttonFont, COLOR_SELECT)
		}
	} else {
		rl.DrawRectangle(x, y, buttonWidth, buttonHeight, COLOR_UNSELECT)
		centerX, centerY := x+(buttonWidth/2), y+(buttonHeight/2)
		rl.DrawText(text, centerX-(rl.MeasureText(text, buttonFont))/2, centerY-(buttonFont/2), buttonFont, COLOR_SELECT)
	}
}

// Returns Achor Y coord below up button.
func DrawArrowButtonsHorizontal(i int) (int32, int32) {
	var buttonSize = Scale(MENU_BUTTON_HEIGHT) // Square
	var baseX = int(Scale(SCREEN_MARGIN))
	var baseY = int(Scale(SCREEN_MARGIN)) + ((rl.GetRenderHeight()-int(2*Scale(SCREEN_MARGIN)))/(BUTTON_LENGTH+1))*(i+1) - int(buttonSize/2)
	var anchorY = int32(baseY) + buttonSize
	var diff = int32(0)
	d := rl.Vector2{X: float32(baseX + int(buttonSize)), Y: float32(baseY + int(buttonSize))}
	e := rl.Vector2{X: float32(baseX), Y: float32(baseY + int(buttonSize))}
	f := rl.Vector2{X: float32(baseX + int(buttonSize/2)), Y: float32(baseY)}
	rl.DrawTriangle(f, e, d, COLOR_SELECT)
	rl.DrawTriangle(rl.Vector2Add(f, rl.Vector2{X: 0, Y: ScaleF(7)}), rl.Vector2Add(e, rl.Vector2{X: ScaleF(6), Y: ScaleF(-3)}), rl.Vector2Add(d, rl.Vector2{X: ScaleF(-6), Y: ScaleF(-3)}), rl.Black)

	baseY = int(Scale(SCREEN_MARGIN)) + ((rl.GetRenderHeight()-(2*int(Scale(SCREEN_MARGIN))))/(BUTTON_LENGTH+1))*(i+1+1) - int(buttonSize/2)
	diff = int32(baseY) - anchorY + buttonSize
	a := rl.Vector2{X: float32(baseX), Y: float32(baseY + int(buttonSize))}
	b := rl.Vector2{X: float32(baseX + int(buttonSize)), Y: float32(baseY + int(buttonSize))}
	c := rl.Vector2{X: float32(baseX + int(buttonSize/2)), Y: float32(baseY + int(buttonSize*2))}
	rl.DrawTriangle(c, b, a, COLOR_SELECT)
	rl.DrawTriangle(rl.Vector2Add(c, rl.Vector2{X: 0, Y: ScaleF(-7)}), rl.Vector2Add(b, rl.Vector2{X: ScaleF(-6), Y: ScaleF(3)}), rl.Vector2Add(a, rl.Vector2{X: ScaleF(6), Y: ScaleF(3)}), rl.Black)
	return anchorY, diff
}

// Returns Achor X coord Between the buttons.
func DrawArrowButtonsVertical(i int) (int32, int32) {
	var buttonSize = Scale(MENU_BUTTON_HEIGHT) // Square
	baseX, baseY := GetButtonAnchor(i)
	var anchorX = int32(baseX) + buttonSize
	var diff = int32(0)

	d := rl.Vector2{X: float32(baseX + buttonSize), Y: float32(baseY + buttonSize)}
	e := rl.Vector2{X: float32(baseX), Y: float32(baseY + (buttonSize / 2))}
	f := rl.Vector2{X: float32(baseX + buttonSize), Y: float32(baseY)}
	rl.DrawTriangle(f, e, d, COLOR_SELECT)
	rl.DrawTriangle(rl.Vector2Add(f, rl.Vector2{X: ScaleF(-3), Y: ScaleF(6)}), rl.Vector2Add(e, rl.Vector2{X: ScaleF(7), Y: 0}), rl.Vector2Add(d, rl.Vector2{X: ScaleF(-3), Y: ScaleF(-6)}), rl.Black)

	baseX, baseY = GetButtonAnchor(i + 1)
	diff = int32(baseX) - anchorX + buttonSize
	a := rl.Vector2{X: float32(baseX + buttonSize), Y: float32(baseY + buttonSize)}
	b := rl.Vector2{X: float32(baseX + buttonSize), Y: float32(baseY)}
	c := rl.Vector2{X: float32(baseX + (buttonSize * 2)), Y: float32(baseY + (buttonSize / 2))}
	rl.DrawTriangle(c, b, a, COLOR_SELECT)
	rl.DrawTriangle(rl.Vector2Add(c, rl.Vector2{X: ScaleF(-7), Y: 0}), rl.Vector2Add(b, rl.Vector2{X: ScaleF(3), Y: ScaleF(6)}), rl.Vector2Add(a, rl.Vector2{X: ScaleF(3), Y: ScaleF(-6)}), rl.Black)
	return anchorX, diff
}

func DrawMarginBox(drawMargin bool) {
	if drawMargin {
		screenMargin := Scale(SCREEN_MARGIN)
		cornerMargin := Scale(CORNER_MARGIN)

		rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), screenMargin, rl.Yellow)
		rl.DrawRectangle(0, 0, screenMargin, int32(rl.GetScreenHeight()), rl.Yellow)
		rl.DrawRectangle(int32(rl.GetScreenWidth())-screenMargin, 0, screenMargin, int32(rl.GetScreenHeight()), rl.Yellow)
		rl.DrawRectangle(0, int32(rl.GetScreenHeight())-screenMargin, int32(rl.GetScreenWidth()), screenMargin, rl.Yellow)

		rl.DrawRectangle(screenMargin, screenMargin, cornerMargin, cornerMargin, rl.Green)
		rl.DrawRectangle(screenMargin, int32(rl.GetScreenHeight())-screenMargin-cornerMargin, cornerMargin, cornerMargin, rl.Green)
		rl.DrawRectangle(int32(rl.GetScreenWidth())-screenMargin-cornerMargin, screenMargin, cornerMargin, cornerMargin, rl.Green)
		rl.DrawRectangle(
			int32(rl.GetScreenHeight())-screenMargin-cornerMargin,
			int32(rl.GetScreenWidth())-screenMargin-cornerMargin,
			cornerMargin, cornerMargin, rl.Green)
	}

	outLine := Scale(2)
	anchorX := Scale(MENU_BUTTON_HEIGHT) + Scale(SCREEN_MARGIN)
	anchorY := Scale(MENU_BUTTON_HEIGHT) + Scale(SCREEN_MARGIN)
	width := GetDisplayAreaWidth()
	rl.DrawRectangle(anchorX, anchorY, int32(width), outLine, COLOR_SELECT)
	rl.DrawRectangle(anchorX, anchorY, outLine, int32(width), COLOR_SELECT)
	rl.DrawRectangle(anchorX, anchorY+int32(width)-outLine, int32(width), outLine, COLOR_SELECT)
	rl.DrawRectangle(anchorX+int32(width)-outLine, anchorY, outLine, int32(width), COLOR_SELECT)
}
