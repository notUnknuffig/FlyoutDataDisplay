package stateFuel

import (
	"math"
	"strconv"

	"example.com/MFDTest/internal/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	KEY_UP    = rl.KeyUp
	KEY_DOWN  = rl.KeyDown
	KEY_LEFT  = rl.KeyLeft
	KEY_RIGHT = rl.KeyRight
	KEY_F6    = rl.KeyF6
	KEY_F5    = rl.KeyF6
	KEY_F7    = rl.KeyF6
	KEY_F8    = rl.KeyF6
	KEY_F9    = rl.KeyF6
)

var (
	COLOR_FUEL                = rl.Green
	COLOR_FUEL_BACKGROUND     = rl.DarkGreen
	COLOR_FUEL_LOW            = rl.Color{R: 255, G: 153, B: 0, A: 255}
	COLOR_FUEL_LOW_BACKGROUND = rl.Color{R: 119, G: 79, B: 0, A: 255}
)

type _state struct {
}

func Init() _state {
	return _state{}
}

func (s _state) Input() state.State {
	if state.GlobalFlightData == nil {
		return s
	}

	return s
}

func (s _state) Draw() {
	if state.GlobalFlightData == nil {
		state.DrawNoData()
		return
	}

	var margin = state.Scale(16)
	var offset = state.Scale(8)
	var width = state.Scale(32)
	var height = state.GetDisplayAreaWidth() - (margin * 2)
	var fuelHeight = int32(math.Round(float64(height) * float64(state.GlobalFlightData.FuelRatio)))
	x := state.Scale(state.SCREEN_MARGIN+state.MENU_BUTTON_SIZE) + margin
	y := state.Scale(state.SCREEN_MARGIN+state.MENU_BUTTON_SIZE) + margin
	rl.DrawRectangle(x, y, width, height, state.COLOR_UNSELECT)
	rl.DrawRectangle(
		x+offset/2,
		y+offset/2+(height-fuelHeight),
		width-offset,
		fuelHeight-offset,
		COLOR_FUEL,
	)

	var xOffset = width + margin
	for i := 0; i < len(state.GlobalFlightData.FuelTanks); i++ {
		if i+1 < len(state.GlobalFlightData.FuelTanks) {
			if state.GlobalFlightData.FuelTanks[i+1].Capacity == state.GlobalFlightData.FuelTanks[i].Capacity &&
				state.GlobalFlightData.FuelTanks[i+1].Priority == state.GlobalFlightData.FuelTanks[i].Priority {
				xOffset += drawMirroredFuelState(state.GlobalFlightData.FuelTanks[i+1], state.GlobalFlightData.FuelTanks[i], xOffset)
				i++
				continue
			}
		}
		xOffset += drawFuelState(state.GlobalFlightData.FuelTanks[i], xOffset)
	}
}

func drawFuelState(fuelTank state.FuelTank, xOffset int32) int32 {
	var color rl.Color
	var bColor rl.Color
	var margin = state.Scale(16)
	var offset = state.Scale(8)
	var width = state.Scale(32)
	var height = state.GetDisplayAreaWidth() - (margin * 2) - state.Scale(100)

	var fuelHeight int32
	if fuelTank.IsEmpty {
		fuelHeight = height
		color = COLOR_FUEL_LOW_BACKGROUND
		bColor = COLOR_FUEL_LOW
	} else {
		fuelHeight = int32(math.Round(float64(height) * float64(fuelTank.FuelPercent)))
		if fuelTank.FuelPercent < 0.25 {
			color = COLOR_FUEL_LOW
			bColor = COLOR_FUEL_LOW_BACKGROUND
		} else {
			color = COLOR_FUEL
			bColor = COLOR_FUEL_BACKGROUND
		}
	}
	x := state.Scale(state.SCREEN_MARGIN+state.MENU_BUTTON_SIZE) + margin + xOffset
	y := state.Scale(state.SCREEN_MARGIN+state.MENU_BUTTON_SIZE) + margin
	rl.DrawRectangle(x, y, width, height, bColor)
	rl.DrawRectangle(
		x+offset/2,
		y+offset/2+(height-fuelHeight),
		width-offset,
		fuelHeight-offset,
		color,
	)
	if fuelTank.IsEmpty {
		rl.DrawLineEx(rl.Vector2{X: float32(x + state.Scale(2)), Y: float32(y + state.Scale(2))}, rl.Vector2{X: float32(x + width - state.Scale(2)), Y: float32(y + height)}, state.ScaleF(4), bColor)
		rl.DrawLineEx(rl.Vector2{X: float32(x + width - state.Scale(2)), Y: float32(y + state.Scale(2))}, rl.Vector2{X: float32(x + state.Scale(2)), Y: float32(y + height)}, state.ScaleF(4), bColor)
	}

	for i := 0; i < 21; i++ {
		var scaleWidth = state.Scale(2)
		if i%5 == 0 {
			scaleWidth = 4
			str := strconv.FormatInt(int64(100-(5*i)), 10)
			rl.DrawText(str, x+width+(margin-rl.MeasureText(str, state.Scale(10)))/2, y+((height)/20)*int32(i)-state.Scale(5)+(offset/2), state.Scale(10), state.COLOR_TEXT_SELECT)
		}
		rl.DrawRectangle(x+width-(margin/2), y+((height)/20)*int32(i)-(scaleWidth/2)+(offset/2), margin/2, scaleWidth, state.COLOR_TEXT_SELECT)
	}
	return width + margin
}

func drawMirroredFuelState(aFuelTank state.FuelTank, bFuelTank state.FuelTank, xOffset int32) int32 {
	var color rl.Color
	var bColor rl.Color
	var margin = state.Scale(16)
	var offset = state.Scale(8)
	var width = state.Scale(32)
	var height = state.GetDisplayAreaWidth() - (margin * 2) - state.Scale(100)

	var fuelHeight int32
	if aFuelTank.IsEmpty {
		fuelHeight = height
		color = COLOR_FUEL_LOW_BACKGROUND
		bColor = COLOR_FUEL_LOW
	} else {
		fuelHeight = int32(math.Round(float64(height) * float64(aFuelTank.FuelPercent)))
		if aFuelTank.FuelPercent < 0.25 {
			color = COLOR_FUEL_LOW
			bColor = COLOR_FUEL_LOW_BACKGROUND
		} else {
			color = COLOR_FUEL
			bColor = COLOR_FUEL_BACKGROUND
		}
	}
	x := state.Scale(state.SCREEN_MARGIN+state.MENU_BUTTON_SIZE) + margin + xOffset
	y := state.Scale(state.SCREEN_MARGIN+state.MENU_BUTTON_SIZE) + margin
	rl.DrawRectangle(x, y, width, height, bColor)
	rl.DrawRectangle(
		x+offset/2,
		y+offset/2+(height-fuelHeight),
		width-offset,
		fuelHeight-offset,
		color,
	)
	if aFuelTank.IsEmpty {
		rl.DrawLineEx(rl.Vector2{X: float32(x + state.Scale(2)), Y: float32(y + state.Scale(2))}, rl.Vector2{X: float32(x + width - state.Scale(2)), Y: float32(y + height)}, state.ScaleF(4), bColor)
		rl.DrawLineEx(rl.Vector2{X: float32(x + width - state.Scale(2)), Y: float32(y + state.Scale(2))}, rl.Vector2{X: float32(x + state.Scale(2)), Y: float32(y + height)}, state.ScaleF(4), bColor)
	}

	if bFuelTank.IsEmpty {
		fuelHeight = height
		color = COLOR_FUEL_LOW_BACKGROUND
		bColor = COLOR_FUEL_LOW
	} else {
		fuelHeight = int32(math.Round(float64(height) * float64(aFuelTank.FuelPercent)))
		if bFuelTank.FuelPercent < 0.25 {
			color = COLOR_FUEL_LOW
			bColor = COLOR_FUEL_LOW_BACKGROUND
		} else {
			color = COLOR_FUEL
			bColor = COLOR_FUEL_BACKGROUND
		}
	}
	rl.DrawRectangle(x+width+margin, y, width, height, bColor)
	rl.DrawRectangle(
		x+offset/2+width+margin,
		y+offset/2+(height-fuelHeight),
		width-offset,
		fuelHeight-offset,
		color,
	)
	if bFuelTank.IsEmpty {
		rl.DrawLineEx(rl.Vector2{X: float32(x + width + margin + state.Scale(2)), Y: float32(y + state.Scale(2))}, rl.Vector2{X: float32(x + width*2 + margin - state.Scale(2)), Y: float32(y + height)}, state.ScaleF(4), bColor)
		rl.DrawLineEx(rl.Vector2{X: float32(x + width*2 + margin - state.Scale(2)), Y: float32(y + state.Scale(2))}, rl.Vector2{X: float32(x + width + margin + state.Scale(2)), Y: float32(y + height)}, state.ScaleF(4), bColor)
	}

	for i := 0; i < 21; i++ {
		var scaleWidth = state.Scale(2)
		if i%5 == 0 {
			scaleWidth = 4
			str := strconv.FormatInt(int64(100-(5*i)), 10)
			rl.DrawText(str, x+width+(margin-rl.MeasureText(str, state.Scale(10)))/2, y+((height)/20)*int32(i)-state.Scale(5)+(offset/2), state.Scale(10), state.COLOR_TEXT_SELECT)
		}
		rl.DrawRectangle(x+width-(margin/2), y+((height)/20)*int32(i)-(scaleWidth/2)+(offset/2), margin/2, scaleWidth, state.COLOR_TEXT_SELECT)
		rl.DrawRectangle(x+width+margin, y+((height)/20)*int32(i)-(scaleWidth/2)+(offset/2), margin/2, scaleWidth, state.COLOR_TEXT_SELECT)
	}
	return (width + margin) * 2
}
