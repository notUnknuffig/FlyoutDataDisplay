package stateEngine

import (
	"strconv"

	"example.com/MFDTest/internal/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type _state struct {
}

func Init() _state {
	return _state{}
}

func (s _state) Input() state.State {
	return s
}

func (s _state) Draw() {
	if state.GlobalFlightData == nil {
		state.DrawNoData()
		return
	}
	for i := 0; i < len(state.GlobalFlightData.JetEngines); i++ {
		s.drawEngineStatistics(100, 100, i, 1)
	}
	for i := 0; i < len(state.GlobalFlightData.PistonEngines); i++ {
		s.drawEngineStatistics(100, 100, i, 1)
	}
}

var COLOR_AFTERBURNER_BACKGROUND = rl.Color{R: 153, G: 102, B: 0, A: 255}
var COLOR_AFTERBURNER_FORGROUND = rl.Color{R: 255, G: 153, B: 0, A: 255}

func (s _state) drawEngineStatistics(anchorX, anchorY int, engine int, scale int) {
	var boxSize = 300
	var margin = 10
	var centerX = anchorX + (boxSize / 2)
	var centerY = anchorY + (boxSize / 2)

	// rl.DrawRectangle(int32(anchorX), int32(anchorY), int32(boxSize), int32(boxSize), state.COLOR_TEXT_UNSELECT)
	if state.GlobalFlightData.JetEngines[engine].HasAfterburner {
		rl.DrawRing(rl.Vector2{X: float32(centerX), Y: float32(centerY)}, float32(boxSize/2-margin-20), float32(boxSize/2-margin), 306-1.5, 360, 0, COLOR_AFTERBURNER_BACKGROUND)
		rl.DrawRectangle(int32(anchorX+boxSize-margin-20), int32(centerY), 20, 3, COLOR_AFTERBURNER_BACKGROUND)
		rl.DrawRing(rl.Vector2{X: float32(centerX), Y: float32(centerY)}, float32(boxSize/2-(margin+3)-14), float32(boxSize/2-(margin+3)), 306, 306+54*state.GlobalFlightData.JetEngines[engine].AfterburnerThrottle, 0, COLOR_AFTERBURNER_FORGROUND)
	}

	margin = 32
	rl.DrawRing(rl.Vector2{X: float32(centerX), Y: float32(centerY)}, float32(boxSize/2-margin-20), float32(boxSize/2-margin), 180, 360, 0, state.COLOR_UNSELECT)
	rl.DrawRectangle(int32(anchorX+margin), int32(centerY), 20, 3, state.COLOR_UNSELECT)
	rl.DrawRectangle(int32(anchorX+boxSize-margin-20), int32(centerY), 20, 3, state.COLOR_UNSELECT)
	rl.DrawRing(rl.Vector2{X: float32(centerX), Y: float32(centerY)}, float32(boxSize/2-(margin+3)-14), float32(boxSize/2-(margin+3)), 180, 180+180*state.GlobalFlightData.JetEngines[engine].Throttle, 0, state.COLOR_SELECT)

	rl.DrawText("Fuel Flow "+strconv.FormatFloat(float64(state.GlobalFlightData.JetEngines[engine].FuelFlow), 'f', -1, 64), int32(anchorX+margin), int32(centerY+20), 20, state.COLOR_SELECT)
	rl.DrawText("Alt Power "+strconv.FormatFloat(float64(state.GlobalFlightData.JetEngines[engine].AlternatorPower), 'f', -1, 64), int32(anchorX+margin), int32(centerY+40), 20, state.COLOR_SELECT)
	rl.DrawText("Hydr Pres "+strconv.FormatFloat(float64(state.GlobalFlightData.JetEngines[engine].HydraulicPower), 'f', -1, 64), int32(anchorX+margin), int32(centerY+60), 20, state.COLOR_SELECT)
}
