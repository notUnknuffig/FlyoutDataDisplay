package stateEngine

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
	KEY_APPLY = rl.KeyEnter
)

type _state struct {
	offset    int
	maxOffset int
}

func Init() _state {
	return _state{
		offset: 0,
	}
}

func (s _state) Input() state.State {
	s.maxOffset = int(math.Ceil(float64(len(state.GlobalFlightData.JetEngines)+len(state.GlobalFlightData.PistonEngines))/4) - 1)
	if rl.IsKeyPressed(KEY_DOWN) && s.offset < s.maxOffset {
		s.offset = s.offset + 1
	} else if rl.IsKeyPressed(KEY_UP) && s.offset > 0 {
		s.offset = s.offset - 1
	}
	return s
}

func (s _state) Draw() {
	if state.GlobalFlightData == nil {
		state.DrawNoData()
		return
	}

	var length = 0
	if len(state.GlobalFlightData.JetEngines)-s.offset*4 < 4 {
		length = len(state.GlobalFlightData.JetEngines) - s.offset*4
	} else {
		length = 4
	}
	for i := 0; i < length; i++ {
		s.drawEngineStatistics(state.Scale(60)+state.Scale(300)*int32(i%2), state.Scale(60)+state.Scale(300)*int32(math.Floor(float64(i/2))), i+(s.offset*4), true)
		if i < 2 && (length > 2 || len(state.GlobalFlightData.PistonEngines) >= 2) {
			rl.DrawRectangle(state.Scale(60)+state.Scale(300)*int32((i)%2)+state.Scale(20), state.Scale(60)+state.Scale(300)*int32(math.Floor(float64(i/2)))+state.Scale(300), state.Scale(260), 2, state.COLOR_SELECT)
		}
		if i%2 == 0 && (length > 1 || len(state.GlobalFlightData.PistonEngines) > 0) {
			rl.DrawRectangle(state.Scale(60)+state.Scale(300)*int32((i+1)%2)-1, state.Scale(60)+state.Scale(300)*int32(math.Floor(float64(i/2)))+state.Scale(20), 2, state.Scale(260), state.COLOR_SELECT)
		}
	}

	// var o = len(state.GlobalFlightData.PistonEngines) % 4
	var o = 0
	if len(state.GlobalFlightData.JetEngines)-s.offset*4 < 4 {

		if len(state.GlobalFlightData.JetEngines)-s.offset*4 > 0 {
			o = len(state.GlobalFlightData.JetEngines) % 4
		}
		if len(state.GlobalFlightData.PistonEngines)+len(state.GlobalFlightData.JetEngines)-s.offset*4 < 4 {
			length = len(state.GlobalFlightData.PistonEngines) + len(state.GlobalFlightData.JetEngines) - s.offset*4 - o
		} else {
			length = 4 - o
		}
		for i := 0; i < length; i++ {
			s.drawEngineStatistics(state.Scale(60)+state.Scale(300)*int32((i+o)%2), state.Scale(60)+state.Scale(300)*int32(math.Floor(float64((i+o)/2))), i+o+(s.offset*4)-len(state.GlobalFlightData.JetEngines), false)
			if i+o < 2 && length > 2 {
				rl.DrawRectangle(state.Scale(60)+state.Scale(300)*int32((i+o)%2)+state.Scale(20), state.Scale(60)+state.Scale(300)*int32(math.Floor(float64((i+o)/2)))+state.Scale(300), state.Scale(260), 2, state.COLOR_SELECT)
			}
			if (i+o)%2 == 0 && length > 1 {
				rl.DrawRectangle(state.Scale(60)+state.Scale(300)*int32(((i+o)+1)%2)-1, state.Scale(60)+state.Scale(300)*int32(math.Floor(float64((i+o)/2)))+state.Scale(20), 2, state.Scale(260), state.COLOR_SELECT)
			}
		}
	}
	anchorY, diff := state.DrawArrowButtons(0)
	str := strconv.FormatInt(int64(s.offset), 10) + " - " + strconv.FormatInt(int64(s.maxOffset), 10)
	center := (state.Scale(state.MENU_BUTTON_SIZE)-rl.MeasureText(str, state.Scale(12)))/2 + state.SCREEN_MARGIN
	rl.DrawText(str, center, anchorY+(diff-state.Scale(12))/2, state.Scale(12), state.COLOR_SELECT)
}

var COLOR_AFTERBURNER_BACKGROUND = rl.Color{R: 153, G: 102, B: 0, A: 255}
var COLOR_AFTERBURNER_FORGROUND = rl.Color{R: 255, G: 153, B: 0, A: 255}

func (s _state) drawEngineStatistics(anchorX, anchorY int32, engine int, turbine bool) {
	var boxSize = state.Scale(300)
	var margin = state.Scale(20)
	var ringWidth = state.Scale(20)
	var ringOffset = state.Scale(8)
	var ringGap = int32(0)
	var centerX = int(anchorX + (boxSize / 2))
	var centerY = int(anchorY + (boxSize / 2))
	var fontSize = state.Scale(20)

	var throttle = float32(0.0)
	var rpm = float32(0.0)
	if turbine {
		throttle = state.GlobalFlightData.JetEngines[engine].Throttle
		// rpm = int(state.GlobalFlightData.JetEngines[engine].RPM)
	} else {
		throttle = state.GlobalFlightData.PistonEngines[engine].Throttle
		rpm = state.GlobalFlightData.PistonEngines[engine].RPM
	}

	// rl.DrawRectangle(int32(anchorX), int32(anchorY), int32(boxSize), int32(boxSize), state.COLOR_TEXT_UNSELECT)
	rl.DrawText(strconv.FormatInt(int64(engine+1), 10), anchorX+margin, anchorY+margin, fontSize, state.COLOR_SELECT)
	if turbine {
		if state.GlobalFlightData.JetEngines[engine].HasAfterburner {
			ringGap = state.Scale(25)
			rl.DrawRing(
				rl.Vector2{X: float32(centerX), Y: float32(centerY)},
				float32(boxSize/2-margin-ringWidth),
				float32(boxSize/2-margin),
				306-1.8, 360,
				0,
				COLOR_AFTERBURNER_BACKGROUND,
			)
			rl.DrawRectangle(anchorX+int32(boxSize-margin-ringWidth), int32(centerY), int32(ringWidth), ringOffset/2, COLOR_AFTERBURNER_BACKGROUND)
			rl.DrawRing(
				rl.Vector2{X: float32(centerX), Y: float32(centerY)},
				float32(boxSize/2-margin+(ringOffset/2)-ringWidth),
				float32(boxSize/2-margin-(ringOffset/2)),
				306,
				306+54*state.GlobalFlightData.JetEngines[engine].AfterburnerThrottle,
				0,
				COLOR_AFTERBURNER_FORGROUND,
			)
		}

		textFuelFlow := strconv.FormatFloat(math.Round(float64(state.GlobalFlightData.JetEngines[engine].FuelFlow*100))/100, 'f', 1, 64) + "kg/s"
		rl.DrawText("Fuel Flow", anchorX+int32(margin), int32(centerY)+margin, fontSize, state.COLOR_SELECT)
		rl.DrawText(textFuelFlow, anchorX+boxSize-int32(margin)-rl.MeasureText(textFuelFlow, fontSize), int32(centerY)+margin, fontSize, state.COLOR_SELECT)

		textAltPwr := strconv.FormatFloat(math.Round(float64(state.GlobalFlightData.JetEngines[engine].AlternatorPower/100))/10, 'f', 1, 64) + "kw/h"
		rl.DrawText("Alt Pwr", anchorX+int32(margin), int32(centerY)+margin+fontSize*2, fontSize, state.COLOR_SELECT)
		rl.DrawText(textAltPwr, anchorX+boxSize-int32(margin)-rl.MeasureText(textAltPwr, fontSize), int32(centerY)+margin+fontSize, fontSize, state.COLOR_SELECT)

		textHydrPwr := strconv.FormatFloat(math.Round(float64(state.GlobalFlightData.JetEngines[engine].HydraulicPower/100))/10, 'f', 1, 64) + "kw/h"
		rl.DrawText("Hydr Pwr", anchorX+int32(margin), int32(centerY)+margin+fontSize, fontSize, state.COLOR_SELECT)
		rl.DrawText(textHydrPwr, anchorX+boxSize-int32(margin)-rl.MeasureText(textHydrPwr, fontSize), int32(centerY)+margin+fontSize*2, fontSize, state.COLOR_SELECT)
	} else {
		ringGap = state.Scale(25)
		rl.DrawRectangle(anchorX+int32(margin), int32(centerY), int32(ringWidth), ringOffset/2, COLOR_AFTERBURNER_BACKGROUND)
		rl.DrawRing(
			rl.Vector2{X: float32(centerX), Y: float32(centerY)},
			float32(boxSize/2-margin-ringWidth),
			float32(boxSize/2-margin),
			180, 360,
			0,
			COLOR_AFTERBURNER_BACKGROUND,
		)
		rl.DrawRectangle(anchorX+int32(boxSize-margin-ringWidth), int32(centerY), int32(ringWidth), ringOffset/2, COLOR_AFTERBURNER_BACKGROUND)
		rl.DrawRing(
			rl.Vector2{X: float32(centerX), Y: float32(centerY)},
			float32(boxSize/2-margin+(ringOffset/2)-ringWidth),
			float32(boxSize/2-margin-(ringOffset/2)),
			180,
			180+180*float32(rpm/6000),
			0,
			COLOR_AFTERBURNER_FORGROUND,
		)

		textFuelFlow := strconv.FormatFloat(math.Round(float64(state.GlobalFlightData.PistonEngines[engine].FuelFlow*100))/100, 'f', 1, 64) + "kg/s"
		rl.DrawText("Fuel Flow", anchorX+int32(margin), int32(centerY)+margin, fontSize, state.COLOR_SELECT)
		rl.DrawText(textFuelFlow, anchorX+boxSize-int32(margin)-rl.MeasureText(textFuelFlow, fontSize), int32(centerY)+margin, fontSize, state.COLOR_SELECT)

		textPower := strconv.FormatFloat(math.Round(float64(state.GlobalFlightData.PistonEngines[engine].Power/100))/10, 'f', 1, 64) + "kw/h"
		rl.DrawText("Power", anchorX+int32(margin), int32(centerY)+margin+fontSize, fontSize, state.COLOR_SELECT)
		rl.DrawText(textPower, anchorX+boxSize-int32(margin)-rl.MeasureText(textPower, fontSize), int32(centerY)+margin+fontSize, fontSize, state.COLOR_SELECT)

		textTemp := strconv.FormatFloat(math.Round(float64(state.GlobalFlightData.PistonEngines[engine].Temperature*100))/100, 'f', 1, 64) + "°k"
		rl.DrawText("Temp", anchorX+int32(margin), int32(centerY)+margin+fontSize*2, fontSize, state.COLOR_SELECT)
		rl.DrawText(textTemp, anchorX+boxSize-int32(margin)-rl.MeasureText(textTemp, fontSize), int32(centerY)+margin+fontSize*2, fontSize, state.COLOR_SELECT)
	}
	rl.DrawRing(
		rl.Vector2{X: float32(centerX), Y: float32(centerY)},
		float32(boxSize/2-margin-ringWidth-ringGap),
		float32(boxSize/2-margin-ringGap),
		180,
		360,
		0,
		state.COLOR_UNSELECT,
	)
	rl.DrawRectangle(anchorX+int32(margin+ringGap), int32(centerY), int32(ringWidth), ringOffset/2, state.COLOR_UNSELECT)
	rl.DrawRectangle(anchorX+int32(boxSize-margin-ringWidth-ringGap), int32(centerY), int32(ringWidth), ringOffset/2, state.COLOR_UNSELECT)
	rl.DrawRing(
		rl.Vector2{X: float32(centerX), Y: float32(centerY)},
		float32(boxSize/2-margin+(ringOffset/2)-ringWidth-ringGap),
		float32(boxSize/2-margin-(ringOffset/2)-ringGap),
		180,
		180+180*throttle,
		0,
		state.COLOR_SELECT,
	)
}
