package stateAttitude

import (
	"math"
	"strconv"

	"example.com/MFDTest/internal/state"
	"example.com/MFDTest/internal/stateNavigation"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type _state struct {
	navMan *stateNavigation.NavManager
}

func Init(navMan *stateNavigation.NavManager) _state {
	return _state{
		navMan: navMan,
	}
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

	s.drawHorizon(30)
}

func (s _state) drawHorizon(maxDegree float32) {
	anchorX := int32(rl.GetRenderWidth() / 2)
	anchorY := int32(rl.GetRenderHeight() / 2)
	height := rl.GetRenderHeight() - int(state.Scale(state.MENU_BUTTON_SIZE+state.MENU_BUTTON_SIZE)*2)
	uiFontSize := state.Scale(20)

	// Bank Angle Idicator
	innnerRadius := state.ScaleF(50)
	outerRadius := state.ScaleF(100)

	sin := math.Sin(-stateNavigation.DegToRad(float64(state.GlobalFlightData.Roll)))
	cos := math.Cos(-stateNavigation.DegToRad(float64(state.GlobalFlightData.Roll)))

	inX := cos * float64(innnerRadius)
	inY := sin * float64(innnerRadius)
	outX := cos * float64(outerRadius)
	outY := sin * float64(outerRadius)

	rl.SetLineWidth(state.ScaleF(3))

	rl.DrawLine(anchorX-int32(inX), anchorY-int32(inY), anchorX-int32(outX), anchorY-int32(outY), state.COLOR_SELECT)
	rl.DrawLine(anchorX-int32(outX)-int32(float32(sin)*state.ScaleF(18)), anchorY-int32(outY)+int32(float32(cos)*state.ScaleF(18)), anchorX-int32(outX), anchorY-int32(outY), state.COLOR_SELECT)

	rl.DrawLine(anchorX+int32(inX), anchorY+int32(inY), anchorX+int32(outX), anchorY+int32(outY), state.COLOR_SELECT)
	rl.DrawLine(anchorX+int32(outX)-int32(float32(sin)*state.ScaleF(18)), anchorY+int32(outY)+int32(float32(cos)*state.ScaleF(18)), anchorX+int32(outX), anchorY+int32(outY), state.COLOR_SELECT)

	// Center
	rl.DrawRectangle(anchorX-state.Scale(4), anchorY-state.Scale(4), state.Scale(8), state.Scale(8), state.COLOR_SELECT)

	// Glide Slope
	alphaOffset := float64((-state.GlobalFlightData.Alpha)/maxDegree) * float64(height)
	betaOffset := float64((state.GlobalFlightData.Beta)/maxDegree) * float64(height)
	glideSlopeOffsetX := -float32(sin*alphaOffset) + float32(cos*betaOffset)
	glideSlopeOffsetY := float32(cos*alphaOffset) + float32(sin*betaOffset)
	rl.DrawRing(rl.Vector2{X: float32(anchorX) - glideSlopeOffsetX, Y: float32(anchorY) - glideSlopeOffsetY}, state.ScaleF(13), state.ScaleF(16), 0, 360, 0, state.COLOR_UNSELECT)
	rl.DrawLine(
		anchorX-int32(glideSlopeOffsetX),
		anchorY-int32(glideSlopeOffsetY)-state.Scale(14),
		anchorX-int32(glideSlopeOffsetX),
		anchorY-int32(glideSlopeOffsetY)-state.Scale(24),
		state.COLOR_UNSELECT,
	)
	rl.DrawLine(
		anchorX-int32(glideSlopeOffsetX)+state.Scale(14),
		anchorY-int32(glideSlopeOffsetY),
		anchorX-int32(glideSlopeOffsetX)+state.Scale(24),
		anchorY-int32(glideSlopeOffsetY),
		state.COLOR_UNSELECT,
	)
	rl.DrawLine(
		anchorX-int32(glideSlopeOffsetX)-state.Scale(14),
		anchorY-int32(glideSlopeOffsetY),
		anchorX-int32(glideSlopeOffsetX)-state.Scale(24),
		anchorY-int32(glideSlopeOffsetY),
		state.COLOR_UNSELECT,
	)

	// Heading
	hdgStr := strconv.FormatFloat(float64(state.GlobalFlightData.Heading), 'f', 1, 64)
	hdgWidth := rl.MeasureText(hdgStr, uiFontSize)
	for i := 0; i < 7; i++ {
		hdgCos := math.Cos(stateNavigation.DegToRad(35 - math.Mod(float64(state.GlobalFlightData.Heading)+float64(i*10), 70)))
		hdgSin := math.Sin(stateNavigation.DegToRad(35 - math.Mod(float64(state.GlobalFlightData.Heading)+float64(i*10), 70)))
		rl.DrawLine(
			anchorY+int32(hdgSin*float64(state.ScaleF(250))),
			anchorX+int32(hdgCos*float64(state.ScaleF(250))),
			anchorY+int32(hdgSin*float64(state.ScaleF(270))),
			anchorX+int32(hdgCos*float64(state.ScaleF(270))),
			state.COLOR_UNSELECT,
		)
	}
	rl.DrawTriangle(
		rl.Vector2{X: float32(anchorX) + state.ScaleF(12), Y: float32(anchorY + state.Scale(200) + uiFontSize*2)},
		rl.Vector2{X: float32(anchorX) - state.ScaleF(12), Y: float32(anchorY + state.Scale(200) + uiFontSize*2)},
		rl.Vector2{X: float32(anchorX), Y: float32(anchorY+state.Scale(200)+uiFontSize*2) + state.ScaleF(20)},
		state.COLOR_SELECT,
	)
	rl.DrawText(hdgStr, anchorX-hdgWidth/2, anchorY+state.Scale(200)+uiFontSize/2, uiFontSize, state.COLOR_SELECT)

	// Nav Heading
	if s.navMan.SelectedObject >= 0 {
		var obj stateNavigation.MappedObject
		switch s.navMan.SelectedType {
		case stateNavigation.NAV_POINT:
			obj = s.navMan.NavPoints[s.navMan.SelectedObject]
		case stateNavigation.AIRFIELD:
			obj = s.navMan.Airfields[s.navMan.SelectedObject]
		}
		_, _, _, bearing := s.navMan.CoordToDistance(
			float64(state.GlobalFlightData.Latitude),
			float64(state.GlobalFlightData.Longitude),
			float64(obj.Latitude),
			float64(obj.Longitude),
			float64(state.GlobalFlightData.Heading),
		)

		// TODO: Limit Angle between -20 and 20
		navAngleDiff := math.Mod(math.Mod(bearing, math.Pi*2)-math.Mod(stateNavigation.DegToRad(float64(state.GlobalFlightData.Heading)), math.Pi*2)+math.Pi*3, math.Pi*2) - math.Pi
		navAngleClamp := math.Min(math.Max(navAngleDiff, stateNavigation.DegToRad(-35)), stateNavigation.DegToRad(35))
		xOff := float32(math.Sin(navAngleClamp) * float64(state.ScaleF(260)))
		yOff := float32(math.Cos(navAngleClamp) * float64(state.ScaleF(260)))
		rectOut := rl.Rectangle{X: float32(anchorX) + xOff, Y: float32(anchorY) + yOff, Width: state.ScaleF(18), Height: state.ScaleF(18)}
		rectIn := rl.Rectangle{X: float32(anchorX) + xOff, Y: float32(anchorY) + yOff, Width: state.ScaleF(12), Height: state.ScaleF(12)}
		rl.DrawRectanglePro(rectOut, rl.Vector2{X: state.ScaleF(9), Y: state.ScaleF(9)}, float32(45+stateNavigation.RadToDeg(-navAngleClamp)), state.COLOR_SELECT)
		rl.DrawRectanglePro(rectIn, rl.Vector2{X: state.ScaleF(6), Y: state.ScaleF(6)}, float32(45+stateNavigation.RadToDeg(-navAngleClamp)), rl.Black)

	}

	// Spd Info
	strSpd := strconv.FormatInt(int64(math.Round(float64(state.GlobalFlightData.Airspeed))), 10) + "m/s"
	infoMargin := state.Scale(8)
	infoOutline := state.Scale(2)
	infoAnchorX := state.Scale(state.SCREEN_MARGIN*2 + state.MENU_BUTTON_SIZE)
	infoAnchorY := int32(rl.GetScreenHeight()/2) - infoMargin*2 - uiFontSize/2
	alphaStr := "a: " + strconv.FormatFloat(float64(state.GlobalFlightData.Alpha), 'f', 1, 64) + "°"
	gStr := "g: " + strconv.FormatFloat(float64(state.GlobalFlightData.G), 'f', 1, 64)

	wheelOffset := state.Scale(200)
	relSpeed := math.Mod(float64(state.GlobalFlightData.Airspeed), 10) / 10
	var spdWidth int32
	for i := 0; i < 20; i++ {
		spdWidth = state.Scale(20)
		rl.DrawRectangle(
			state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_SIZE),
			(anchorY-wheelOffset)+(int32(i)*state.Scale(20))+int32(math.Round(20*relSpeed)),
			spdWidth,
			state.Scale(2),
			state.COLOR_UNSELECT,
		)
	}
	rl.DrawRectangle(infoAnchorX+state.Scale(30), infoAnchorY, infoOutline*2+rl.MeasureText(strSpd, uiFontSize)+infoMargin*2, infoMargin*2+infoOutline*2+uiFontSize, state.COLOR_SELECT)
	rl.DrawRectangle(infoAnchorX+state.Scale(30)+infoOutline, infoAnchorY+infoOutline, infoMargin*2+rl.MeasureText(strSpd, uiFontSize), infoMargin*2+uiFontSize, state.COLOR_UNSELECT)
	rl.DrawTriangle(
		rl.Vector2{X: float32(infoAnchorX) + state.ScaleF(30), Y: float32(infoAnchorY + infoMargin + infoOutline)},
		rl.Vector2{X: float32(infoAnchorX) + state.ScaleF(10), Y: float32(infoAnchorY + infoMargin + infoOutline + uiFontSize/2)},
		rl.Vector2{X: float32(infoAnchorX) + state.ScaleF(30), Y: float32(infoAnchorY + infoMargin + infoOutline + uiFontSize)},
		state.COLOR_SELECT,
	)
	rl.DrawText(strSpd, infoAnchorX+state.Scale(30)+infoMargin+infoOutline, infoAnchorY+infoMargin+infoOutline, uiFontSize, state.COLOR_SELECT)
	rl.DrawText(alphaStr, infoAnchorX+state.Scale(30)+infoMargin+infoOutline, infoAnchorY+uiFontSize+infoMargin*4+infoOutline*2, state.Scale(18), state.COLOR_SELECT)
	rl.DrawText(gStr, infoAnchorX+state.Scale(30)+infoMargin+infoOutline, infoAnchorY+uiFontSize*2+infoMargin*5+infoOutline*2, state.Scale(18), state.COLOR_SELECT)

	// Alt Info
	strAlt := strconv.FormatInt(int64(math.Round(float64(state.GlobalFlightData.Altitude))), 10) + "m"
	infoAnchorX = int32(rl.GetRenderWidth()) - state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_SIZE)

	wheelOffset = state.Scale(200)
	relAlt := math.Mod(float64(state.GlobalFlightData.Altitude), 10) / 10
	// var altWidth int32
	for i := 0; i < 20; i++ {
		// spdWidth = state.Scale(20)
		rl.DrawRectangle(
			int32(rl.GetRenderWidth())-state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_SIZE)-state.Scale(20),
			(anchorY-wheelOffset)+(int32(i)*state.Scale(20))+int32(math.Round(20*relAlt)),
			state.Scale(20),
			state.Scale(2),
			state.COLOR_UNSELECT,
		)
	}
	rl.DrawRectangle(infoAnchorX-state.Scale(50)-rl.MeasureText(strAlt, uiFontSize), infoAnchorY, infoOutline*2+rl.MeasureText(strAlt, uiFontSize)+infoMargin*2, infoMargin*2+infoOutline*2+uiFontSize, state.COLOR_SELECT)
	rl.DrawRectangle(infoAnchorX-state.Scale(50)-rl.MeasureText(strAlt, uiFontSize)+infoOutline, infoAnchorY+infoOutline, infoMargin*2+rl.MeasureText(strAlt, uiFontSize), infoMargin*2+uiFontSize, state.COLOR_UNSELECT)
	rl.DrawTriangle(
		rl.Vector2{X: float32(infoAnchorX) - state.ScaleF(10), Y: float32(infoAnchorY + infoMargin + infoOutline + uiFontSize/2)},
		rl.Vector2{X: float32(infoAnchorX) - state.ScaleF(30), Y: float32(infoAnchorY + infoMargin + infoOutline)},
		rl.Vector2{X: float32(infoAnchorX) - state.ScaleF(30), Y: float32(infoAnchorY + infoMargin + infoOutline + uiFontSize)},
		state.COLOR_SELECT,
	)
	rl.DrawText(strAlt, infoAnchorX-state.Scale(50)-rl.MeasureText(strAlt, uiFontSize)+infoMargin+infoOutline, infoAnchorY+infoMargin+infoOutline, uiFontSize, state.COLOR_SELECT)

}
