package stateAttitude

import (
	"math"

	"example.com/MFDTest/internal/state"
	"example.com/MFDTest/internal/stateNavigation"
	rl "github.com/gen2brain/raylib-go/raylib"
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

	s.drawHorizon(20)
}

func (s _state) drawHorizon(maxDegree float32) {
	anchorX := int32(rl.GetRenderWidth() / 2)
	anchorY := int32(rl.GetRenderHeight() / 2)
	height := rl.GetRenderHeight() - int(state.Scale(state.MENU_BUTTON_SIZE+state.MENU_BUTTON_SIZE)*2)

	// Bank Angle Idicator
	innnerRadius := state.ScaleF(50)
	outerRadius := state.ScaleF(150)

	inX := math.Cos(stateNavigation.DegToRad(float64(state.GlobalFlightData.Roll))) * float64(innnerRadius)
	inY := math.Sin(stateNavigation.DegToRad(float64(state.GlobalFlightData.Roll))) * float64(innnerRadius)
	outX := math.Cos(stateNavigation.DegToRad(float64(state.GlobalFlightData.Roll))) * float64(outerRadius)
	outY := math.Sin(stateNavigation.DegToRad(float64(state.GlobalFlightData.Roll))) * float64(outerRadius)

	rl.DrawLine(anchorX-int32(inX), anchorY-int32(inY), anchorX-int32(outX), anchorY-int32(outY), state.COLOR_SELECT)
	rl.DrawLine(anchorX+int32(inX), anchorY+int32(inY), anchorX+int32(outX), anchorY+int32(outY), state.COLOR_SELECT)

	// Center
	rl.DrawRectangle(anchorX-state.Scale(4), anchorY-state.Scale(4), state.Scale(8), state.Scale(8), state.COLOR_SELECT)

	// Glide Slope
	glideSlopeOffset := float64((-state.GlobalFlightData.Alpha)/maxDegree) * float64(height)
	glideSlopeOffsetX := float32(math.Cos(stateNavigation.DegToRad(float64(state.GlobalFlightData.Roll))) * glideSlopeOffset)
	glideSlopeOffsetY := float32(math.Sin(stateNavigation.DegToRad(float64(state.GlobalFlightData.Roll))) * glideSlopeOffset)
	rl.DrawRing(rl.Vector2{X: float32(anchorX) - glideSlopeOffsetX, Y: float32(anchorY) - glideSlopeOffsetY}, state.ScaleF(8), state.ScaleF(10), 0, 360, 0, state.COLOR_SELECT)
}
