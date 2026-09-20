package navigator

import (
	"math"
	"strconv"

	"example.com/MFDTest/internal/navigation"
	"example.com/MFDTest/internal/state"
	"example.com/MFDTest/internal/states/navPointEntry"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type _state struct {
	scale    int
	maxScale int
	minScale int
	navMan   *navigation.NavManager
}

func Init(nav *navigation.NavManager) _state {
	s := _state{
		scale:    30,
		maxScale: 150,
		minScale: 10,
		navMan:   nav,
	}
	return s
}

func (s _state) Input() state.State {
	if state.GlobalFlightData == nil {
		return s
	}
	if rl.IsKeyPressed(state.KEY_F8) {
		s.navMan.UseHaversine = !s.navMan.UseHaversine
	}
	if rl.IsKeyPressed(state.KEY_F12) {
		if (s.navMan.SelectedType == navigation.AIRFIELD) && s.navMan.SelectedObject < len(s.navMan.Airfields)-1 {
			s.navMan.SelectedObject += 1
		} else if (s.navMan.SelectedType == navigation.NAV_POINT) && s.navMan.SelectedObject < len(s.navMan.NavPoints)-1 {
			s.navMan.SelectedObject += 1
		}
	} else if rl.IsKeyPressed(state.KEY_F11) && s.navMan.SelectedObject >= 0 {
		s.navMan.SelectedObject -= 1
	}
	if rl.IsKeyPressed(state.KEY_F9) {
		switch s.navMan.SelectedType {
		case navigation.AIRFIELD:
			s.navMan.SelectedType = navigation.NAV_POINT
			s.navMan.SelectedObject = -1
		case navigation.NAV_POINT:
			s.navMan.SelectedType = navigation.AIRFIELD
			s.navMan.SelectedObject = 0
		}
	}
	if rl.IsKeyPressed(state.KEY_F6) && s.scale < s.maxScale {
		s.scale = s.scale + 20
	} else if rl.IsKeyPressed(state.KEY_F7) && s.scale > s.minScale {
		s.scale = s.scale - 20
	}
	if s.navMan.SelectedType == navigation.NAV_POINT && rl.IsKeyPressed(state.KEY_F10) {
		return navPointEntry.Init(s.navMan, s)
	}
	return s
}

var COLOR_PLANE_SELF = rl.Color{R: 68, G: 190, B: 220, A: 255}

func (s _state) Draw() {
	if state.GlobalFlightData == nil {
		state.DrawNoData()
		return
	}

	// Draw Nav Screen
	s.drawNavigation()

	anchorY, diffY := state.DrawArrowButtonsHorizontal(0)
	scaleFontSize := state.Scale(24)
	scaleStr := strconv.FormatInt(int64(s.scale), 10)
	scaleCenter := (state.Scale(state.MENU_BUTTON_HEIGHT)-rl.MeasureText(scaleStr, scaleFontSize))/2 + state.Scale(state.SCREEN_MARGIN)
	rl.DrawRectangle(state.Scale(state.SCREEN_MARGIN), anchorY, state.Scale(state.MENU_BUTTON_HEIGHT), diffY, rl.Black)
	rl.DrawText(scaleStr, scaleCenter, anchorY+(diffY/2)-scaleFontSize+state.Scale(4), scaleFontSize, state.COLOR_SELECT)
	rl.DrawText(state.GlobalOptions.Units.DistanceUnit, state.Scale(state.SCREEN_MARGIN)+state.Scale(state.MENU_BUTTON_HEIGHT)/2-rl.MeasureText(state.GlobalOptions.Units.DistanceUnit, scaleFontSize)/2, anchorY+(diffY/2)+state.Scale(4), scaleFontSize, state.COLOR_SELECT)

	if s.navMan.UseHaversine {
		state.DrawButton("HAV", 8-1)
	} else {
		state.DrawButton("TRI", 8-1)
	}

	state.DrawButton("TYP", 9-1)
	anchorX, diffX := state.DrawArrowButtonsVertical(11 - 1)
	rl.DrawRectangle(anchorX, state.Scale(state.SCREEN_MARGIN), diffX, state.Scale(state.MENU_BUTTON_HEIGHT), rl.Black)
	targetString := "Nothing"
	switch s.navMan.SelectedType {
	case navigation.NAV_POINT:
		state.DrawButton("EDT", 10-1)
		if len(s.navMan.NavPoints) == 0 {
			targetString = "No NP"
		} else if s.navMan.SelectedObject == -1 {

		} else {

			targetString = strconv.FormatInt(int64(s.navMan.SelectedObject+1), 10) + " / " + strconv.FormatInt(int64(len(s.navMan.NavPoints)), 10)
		}
	case navigation.AIRFIELD:
		if len(s.navMan.Airfields) == 0 {
			targetString = "No AF"
		} else if s.navMan.SelectedObject == -1 {

		} else {
			targetString = strconv.FormatInt(int64(s.navMan.SelectedObject+1), 10) + " / " + strconv.FormatInt(int64(len(s.navMan.Airfields)), 10)
		}
	}
	rl.DrawText(
		targetString,
		anchorX+diffX/2-rl.MeasureText(targetString, scaleFontSize)/2,
		state.Scale(state.SCREEN_MARGIN)+(state.Scale(state.MENU_BUTTON_HEIGHT)/2)-(scaleFontSize/2),
		scaleFontSize,
		state.COLOR_SELECT,
	)

	headingFontSize := state.Scale(20)
	headingStr := strconv.FormatFloat(math.Round(float64(state.GlobalFlightData.Heading*100))/100, 'f', 2, 64) + "°"
	headingCenter := state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_HEIGHT) + (state.GetDisplayAreaWidth()-rl.MeasureText(headingStr, headingFontSize))/2
	rl.DrawRectangle(state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_HEIGHT)+(state.GetDisplayAreaWidth()-state.Scale(84))/2, state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_HEIGHT)-state.Scale(1), state.Scale(84), headingFontSize+state.Scale(2), rl.Black)
	rl.DrawText(headingStr, headingCenter, state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_HEIGHT), headingFontSize, state.COLOR_SELECT)
}

func (s _state) drawNavigation() {
	width := state.GetDisplayAreaWidth()
	anchorX := state.Scale(state.SCREEN_MARGIN*2) + state.Scale(state.MENU_BUTTON_HEIGHT) + (width / 2)
	anchorY := state.Scale(state.SCREEN_MARGIN*2) + state.Scale(state.MENU_BUTTON_HEIGHT) + (width * 3 / 4)

	// Draw Marking Rings
	d := (10 / float64(s.scale)) * float64(width/4) * 3
	switch s.scale {
	case 10:
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d/2)-state.ScaleF(0.5), float32(d/2)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d)-state.ScaleF(1.5), float32(d)+state.ScaleF(1.5), 0, 360, 0, state.COLOR_SELECT)
	case 30:
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d)-state.ScaleF(0.5), float32(d)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*2)-state.ScaleF(0.5), float32(d*2)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*3)-state.ScaleF(1.5), float32(d*3)+state.ScaleF(1.5), 0, 360, 0, state.COLOR_SELECT)
	case 50:
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d/2)*5-state.ScaleF(0.5), float32(d/2)*5+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*5)-state.ScaleF(1.5), float32(d*5)+state.ScaleF(1.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d/2)*15-state.ScaleF(0.5), float32(d/2)*15+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
	case 70:
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d)-state.ScaleF(0.5), float32(d)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*3)-state.ScaleF(0.5), float32(d*3)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*5)-state.ScaleF(0.5), float32(d*5)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*7)-state.ScaleF(1.5), float32(d*7)+state.ScaleF(1.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*9)-state.ScaleF(0.5), float32(d*9)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
	case 90:
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*3)-state.ScaleF(0.5), float32(d*3)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*6)-state.ScaleF(0.5), float32(d*6)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*9)-state.ScaleF(1.5), float32(d*9)+state.ScaleF(1.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*12)-state.ScaleF(0.5), float32(d*12)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
	case 110:
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d)-state.ScaleF(0.5), float32(d)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*4)-state.ScaleF(0.5), float32(d*4)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*8)-state.ScaleF(0.5), float32(d*8)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*11)-state.ScaleF(1.5), float32(d*11)+state.ScaleF(1.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*14)-state.ScaleF(0.5), float32(d*14)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
	case 130:
		// rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(a*2)-state.ScaleF(0.5), float32(a*2)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*5)-state.ScaleF(0.5), float32(d*5)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		// rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(a*8)-state.ScaleF(0.5), float32(a*8)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*13)-state.ScaleF(1.5), float32(d*13)+state.ScaleF(1.5), 0, 360, 0, state.COLOR_SELECT)
		// rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(a*14)-state.ScaleF(0.5), float32(a*14)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*18)-state.ScaleF(0.5), float32(d*18)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
	case 150:
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*5)-state.ScaleF(0.5), float32(d*5)+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d)*10-state.ScaleF(0.5), float32(d)*10+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d*15)-state.ScaleF(1.5), float32(d*15)+state.ScaleF(1.5), 0, 360, 0, state.COLOR_SELECT)
		rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, float32(d)*20-state.ScaleF(0.5), float32(d)*20+state.ScaleF(0.5), 0, 360, 0, state.COLOR_SELECT)
	}

	s.drawNavPoints(anchorX, anchorY, width)
	s.drawAirfields(anchorX, anchorY, width)
	s.drawNavInfo(anchorX, anchorY, width)

	planeSize := state.Scale(32)
	planeThickness := state.Scale(2)

	rl.DrawRectangle(anchorX-(planeThickness/2), anchorY, planeThickness, planeSize, COLOR_PLANE_SELF)
	rl.DrawRectangle(anchorX-planeSize/2, anchorY+(planeSize/3)-(planeThickness/2), planeSize, planeThickness, COLOR_PLANE_SELF)
	rl.DrawRectangle(anchorX-planeSize/4, anchorY+(planeSize*7/8)-2, planeSize/2, planeThickness, COLOR_PLANE_SELF)
}

func (s _state) drawNavInfo(anchorX, anchorY int32, width int32) {
	if s.navMan.SelectedObject < 0 {
		return
	}
	var obj navigation.MappedObject
	switch s.navMan.SelectedType {
	case navigation.NAV_POINT:
		obj = s.navMan.NavPoints[s.navMan.SelectedObject]
	case navigation.AIRFIELD:
		obj = s.navMan.Airfields[s.navMan.SelectedObject]
	}

	recAnchorX := int32(rl.GetRenderWidth()) - state.Scale(state.SCREEN_MARGIN*2) - state.Scale(state.MENU_BUTTON_HEIGHT) - state.Scale(200)
	recAnchorY := state.Scale(state.SCREEN_MARGIN*2) + state.Scale(state.MENU_BUTTON_HEIGHT)
	outline := state.Scale(4)
	fontSize := state.Scale(18)
	rl.DrawRectangle(recAnchorX, recAnchorY, state.Scale(200), state.Scale(120), state.COLOR_UNSELECT)
	rl.DrawRectangle(recAnchorX+outline/2, recAnchorY+outline/2, state.Scale(200)-outline, state.Scale(120)-outline, rl.Black)
	rl.DrawText(obj.Name, recAnchorX+state.Scale(200/2)-rl.MeasureText(obj.Name, fontSize)/2, recAnchorY+outline*2, fontSize, state.COLOR_SELECT)

	rl.DrawText(strconv.FormatFloat(float64(obj.Latitude), 'f', 2, 64)+"°", recAnchorX+outline*2, recAnchorY+outline*3+fontSize, fontSize, state.COLOR_SELECT)
	longStr := strconv.FormatFloat(float64(obj.Longitude), 'f', 2, 64) + "°"
	rl.DrawText(longStr, recAnchorX+state.Scale(200)-outline*2-rl.MeasureText(longStr, fontSize), recAnchorY+outline*3+fontSize, fontSize, state.COLOR_SELECT)

	rl.DrawText("Alt:", recAnchorX+outline*2, recAnchorY+outline*4+fontSize*2, fontSize, state.COLOR_SELECT)
	altStr := strconv.FormatFloat(float64(obj.Altitude)*state.GlobalOptions.Units.HightConversion, 'f', 3, 64) + state.GlobalOptions.Units.HightUnit
	rl.DrawText(altStr, recAnchorX+state.Scale(200)-outline*2-rl.MeasureText(altStr, fontSize), recAnchorY+outline*4+fontSize*2, fontSize, state.COLOR_SELECT)

	rl.DrawText("Dist:", recAnchorX+outline*2, recAnchorY+outline*5+fontSize*3, fontSize, state.COLOR_SELECT)
	distStr := strconv.FormatFloat(obj.Dist*state.GlobalOptions.Units.DistanceConversion, 'f', 3, 64) + state.GlobalOptions.Units.DistanceUnit
	rl.DrawText(distStr, recAnchorX+state.Scale(200)-outline*2-rl.MeasureText(distStr, fontSize), recAnchorY+outline*5+fontSize*3, fontSize, state.COLOR_SELECT)

	rl.DrawText("Bear:", recAnchorX+outline*2, recAnchorY+outline*6+fontSize*4, fontSize, state.COLOR_SELECT)
	bearStr := strconv.FormatFloat(math.Mod(360+navigation.RadToDeg(obj.Bearing), 360), 'f', 2, 64) + "°"
	rl.DrawText(bearStr, recAnchorX+state.Scale(200)-outline*2-rl.MeasureText(bearStr, fontSize), recAnchorY+outline*6+fontSize*4, fontSize, state.COLOR_SELECT)

	if obj.Dist/float64(s.scale) > 0.75 {
		xOff := float32(math.Sin(math.Pi-obj.Bearing+navigation.DegToRad(float64(state.GlobalFlightData.Heading))) * float64(state.ScaleF(150)))
		yOff := float32(math.Cos(math.Pi-obj.Bearing+navigation.DegToRad(float64(state.GlobalFlightData.Heading))) * float64(state.ScaleF(150)))
		rectOut := rl.Rectangle{X: float32(anchorX) + xOff, Y: float32(anchorY) + yOff, Width: state.ScaleF(18), Height: state.ScaleF(18)}
		rectIn := rl.Rectangle{X: float32(anchorX) + xOff, Y: float32(anchorY) + yOff, Width: state.ScaleF(12), Height: state.ScaleF(12)}
		rl.DrawRectanglePro(rectOut, rl.Vector2{X: state.ScaleF(9), Y: state.ScaleF(9)}, float32(45+navigation.RadToDeg(obj.Bearing)-float64(state.GlobalFlightData.Heading)), state.COLOR_SELECT)
		rl.DrawRectanglePro(rectIn, rl.Vector2{X: state.ScaleF(6), Y: state.ScaleF(6)}, float32(45+navigation.RadToDeg(obj.Bearing)-float64(state.GlobalFlightData.Heading)), rl.Black)
	}
}

func (s _state) drawAirfields(anchorX, anchorY int32, width int32) {
	for i := 0; i < len(s.navMan.Airfields); i++ {
		x, y, dist, bearing := s.navMan.CoordToDistance(
			float64(state.GlobalFlightData.Latitude),
			float64(state.GlobalFlightData.Longitude),
			float64(s.navMan.Airfields[i].Latitude),
			float64(s.navMan.Airfields[i].Longitude),
			float64(state.GlobalFlightData.Heading),
		)
		s.navMan.Airfields[i].Dist = dist
		s.navMan.Airfields[i].Bearing = bearing
		relX := (x * state.GlobalOptions.Units.DistanceConversion / float64(s.scale)) * float64(width/4) * 3
		relY := (y * state.GlobalOptions.Units.DistanceConversion / float64(s.scale)) * float64(width/4) * 3

		selected := false
		if s.navMan.SelectedType == navigation.AIRFIELD && s.navMan.SelectedObject == i {
			selected = true
		}
		s.navMan.Airfields[i].DrawNavObject(anchorX+int32(math.Round(relX)), anchorY+int32(math.Round(relY)), selected)
	}
}

func (s _state) drawNavPoints(anchorX, anchorY int32, width int32) {
	var lastX = 0.0
	var lastY = 0.0
	for i := 0; i < len(s.navMan.NavPoints); i++ {
		x, y, dist, bearing := s.navMan.CoordToDistance(
			float64(state.GlobalFlightData.Latitude),
			float64(state.GlobalFlightData.Longitude),
			float64(s.navMan.NavPoints[i].Latitude),
			float64(s.navMan.NavPoints[i].Longitude),
			float64(state.GlobalFlightData.Heading),
		)
		s.navMan.NavPoints[i].Dist = dist
		s.navMan.NavPoints[i].Bearing = bearing
		// fmt.Printf("Distance X = %fkm, ")
		relX := ((x * state.GlobalOptions.Units.DistanceConversion) / float64(s.scale)) * float64(width/4) * 3
		relY := ((y * state.GlobalOptions.Units.DistanceConversion) / float64(s.scale)) * float64(width/4) * 3
		rl.SetLineWidth(state.ScaleF(3))
		if i > 0 {
			rl.DrawLineEx(
				rl.Vector2{
					X: float32(anchorX) + float32(relX),
					Y: float32(anchorY) + float32(relY),
				},
				rl.Vector2{
					X: float32(anchorX) + float32(lastX),
					Y: float32(anchorY) + float32(lastY),
				},
				2,
				navigation.COLOR_NAV_LINE,
			)
		}
		lastX = relX
		lastY = relY

		selected := false
		if s.navMan.SelectedType == navigation.NAV_POINT && s.navMan.SelectedObject == i {
			selected = true
		}
		s.navMan.NavPoints[i].DrawNavObject(anchorX+int32(math.Round(relX)), anchorY+int32(math.Round(relY)), selected)
	}
}
