package stateNavigation

import (
	"math"
	"strconv"

	"example.com/MFDTest/internal/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type _state struct {
	scale          int
	maxScale       int
	minScale       int
	navMan         *NavManager
	useHaversine   bool
	selectedObject int
	selectedType   ObjectType
	isSelecting    bool
}

func Init(nav *NavManager) _state {
	return _state{
		scale:          30,
		maxScale:       150,
		minScale:       10,
		navMan:         nav,
		useHaversine:   false,
		selectedObject: -1,
		selectedType:   NAV_POINT,
		isSelecting:    false,
	}
}

const (
	KEY_UP    = rl.KeyUp
	KEY_DOWN  = rl.KeyDown
	KEY_LEFT  = rl.KeyLeft
	KEY_RIGHT = rl.KeyRight
	KEY_F6    = rl.KeyF6
	KEY_F7    = rl.KeyF7
	KEY_F8    = rl.KeyF8
	KEY_F9    = rl.KeyF9
)

func (s _state) Input() state.State {
	if state.GlobalFlightData == nil {
		return s
	}
	if rl.IsKeyPressed(KEY_F6) {
		s.useHaversine = !s.useHaversine
	}
	if rl.IsKeyPressed(KEY_F7) {
		s.isSelecting = !s.isSelecting
	}
	if s.isSelecting && rl.IsKeyPressed(KEY_RIGHT) {
		if (s.selectedType == AIRFIELD) && s.selectedObject < len(s.navMan.Airfields)-1 {
			s.selectedObject += 1
		} else if (s.selectedType == NAV_POINT) && s.selectedObject < len(s.navMan.NavPoints)-1 {
			s.selectedObject += 1
		}
	} else if s.isSelecting && rl.IsKeyPressed(KEY_LEFT) && s.selectedObject >= 0 {
		s.selectedObject -= 1
	}
	if s.isSelecting && rl.IsKeyPressed(KEY_F8) {
		switch s.selectedType {
		case AIRFIELD:
			s.selectedType = NAV_POINT
			s.selectedObject = -1
		case NAV_POINT:
			s.selectedType = AIRFIELD
			s.selectedObject = 0
		}
	}
	if rl.IsKeyPressed(KEY_UP) && s.scale < s.maxScale {
		s.scale = s.scale + 20
	} else if rl.IsKeyPressed(KEY_DOWN) && s.scale > s.minScale {
		s.scale = s.scale - 20
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

	anchorY, diff := state.DrawArrowButtons(0)
	scaleFontSize := state.Scale(24)
	scaleStr := strconv.FormatInt(int64(s.scale), 10)
	scaleCenter := (state.Scale(state.MENU_BUTTON_SIZE)-rl.MeasureText(scaleStr, scaleFontSize))/2 + state.Scale(state.SCREEN_MARGIN)
	rl.DrawRectangle(scaleCenter-state.Scale(8), anchorY, rl.MeasureText(scaleStr, scaleFontSize)+state.Scale(16), diff, rl.Black)
	rl.DrawText(scaleStr, scaleCenter, anchorY+(diff/2)-scaleFontSize+state.Scale(4), scaleFontSize, state.COLOR_SELECT)
	rl.DrawText("km", state.Scale(state.SCREEN_MARGIN)+state.Scale(8), anchorY+(diff/2)+state.Scale(4), scaleFontSize, state.COLOR_SELECT)

	fontSize := state.Scale(20)
	color := state.COLOR_SELECT
	bColor := state.COLOR_UNSELECT
	if s.isSelecting {
		color = state.COLOR_UNSELECT
		bColor = state.COLOR_SELECT
	}

	rl.DrawRectangle(int32(rl.GetRenderWidth())-state.Scale(state.SCREEN_MARGIN)-state.Scale(state.MENU_BUTTON_SIZE), state.GetButtonAnchorByIndex(0), state.Scale(state.MENU_BUTTON_SIZE), state.Scale(60), bColor)
	rl.DrawText("S", int32(rl.GetRenderWidth())-state.Scale(state.SCREEN_MARGIN)-state.Scale(state.MENU_BUTTON_SIZE)+state.Scale(10), state.GetButtonAnchorByIndex(0)+state.Scale(4), fontSize, color)
	rl.DrawText("E", int32(rl.GetRenderWidth())-state.Scale(state.SCREEN_MARGIN)-state.Scale(state.MENU_BUTTON_SIZE)+state.Scale(10), state.GetButtonAnchorByIndex(0)+state.Scale(0)+fontSize, fontSize, color)
	rl.DrawText("L", int32(rl.GetRenderWidth())-state.Scale(state.SCREEN_MARGIN)-state.Scale(state.MENU_BUTTON_SIZE)+state.Scale(10), state.GetButtonAnchorByIndex(0)+state.Scale(-4)+fontSize*2, fontSize, color)

	rl.DrawRectangle(int32(rl.GetRenderWidth())-state.Scale(state.SCREEN_MARGIN)-state.Scale(state.MENU_BUTTON_SIZE), state.GetButtonAnchorByIndex(1), state.Scale(state.MENU_BUTTON_SIZE), state.Scale(60), state.COLOR_UNSELECT)
	rl.DrawText("T", int32(rl.GetRenderWidth())-state.Scale(state.SCREEN_MARGIN)-state.Scale(state.MENU_BUTTON_SIZE)+state.Scale(10), state.GetButtonAnchorByIndex(1)+state.Scale(4), fontSize, state.COLOR_SELECT)
	rl.DrawText("Y", int32(rl.GetRenderWidth())-state.Scale(state.SCREEN_MARGIN)-state.Scale(state.MENU_BUTTON_SIZE)+state.Scale(10), state.GetButtonAnchorByIndex(1)+state.Scale(0)+fontSize, fontSize, state.COLOR_SELECT)
	rl.DrawText("P", int32(rl.GetRenderWidth())-state.Scale(state.SCREEN_MARGIN)-state.Scale(state.MENU_BUTTON_SIZE)+state.Scale(10), state.GetButtonAnchorByIndex(1)+state.Scale(-4)+fontSize*2, fontSize, state.COLOR_SELECT)

	rl.DrawRectangle(state.Scale(state.SCREEN_MARGIN), state.GetButtonAnchorByIndex(2), state.Scale(state.MENU_BUTTON_SIZE), state.Scale(60), state.COLOR_UNSELECT)
	if s.useHaversine {
		rl.DrawText("H", state.Scale(state.SCREEN_MARGIN)+state.Scale(10), state.GetButtonAnchorByIndex(2)+state.Scale(4), fontSize, state.COLOR_SELECT)
		rl.DrawText("A", state.Scale(state.SCREEN_MARGIN)+state.Scale(10), state.GetButtonAnchorByIndex(2)+state.Scale(0)+fontSize, fontSize, state.COLOR_SELECT)
		rl.DrawText("V", state.Scale(state.SCREEN_MARGIN)+state.Scale(10), state.GetButtonAnchorByIndex(2)+state.Scale(-4)+fontSize*2, fontSize, state.COLOR_SELECT)
	} else {
		rl.DrawText("T", state.Scale(state.SCREEN_MARGIN)+state.Scale(10), state.GetButtonAnchorByIndex(2)+state.Scale(4), fontSize, state.COLOR_SELECT)
		rl.DrawText("R", state.Scale(state.SCREEN_MARGIN)+state.Scale(10), state.GetButtonAnchorByIndex(2)+state.Scale(0)+fontSize, fontSize, state.COLOR_SELECT)
		rl.DrawText("I", state.Scale(state.SCREEN_MARGIN)+state.Scale(10), state.GetButtonAnchorByIndex(2)+state.Scale(-4)+fontSize*2, fontSize, state.COLOR_SELECT)
	}

	headingFontSize := state.Scale(20)
	headingStr := strconv.FormatFloat(math.Round(float64(state.GlobalFlightData.Heading*100))/100, 'f', 2, 64) + "°"
	headingCenter := state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_SIZE) + (state.GetDisplayAreaWidth()-rl.MeasureText(headingStr, headingFontSize))/2
	rl.DrawRectangle(state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_SIZE)+(state.GetDisplayAreaWidth()-state.Scale(84))/2, state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_SIZE)-state.Scale(1), state.Scale(84), headingFontSize+state.Scale(2), rl.Black)
	rl.DrawText(headingStr, headingCenter, state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_SIZE), headingFontSize, state.COLOR_SELECT)
}

func (s _state) drawNavigation() {
	width := state.GetDisplayAreaWidth()
	anchorX := state.Scale(state.SCREEN_MARGIN*2) + state.Scale(state.MENU_BUTTON_SIZE) + (width / 2)
	anchorY := state.Scale(state.SCREEN_MARGIN*2) + state.Scale(state.MENU_BUTTON_SIZE) + (width * 3 / 4)

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
	if s.selectedObject < 0 {
		return
	}
	var obj MappedObjects
	switch s.selectedType {
	case NAV_POINT:
		obj = s.navMan.NavPoints[s.selectedObject]
	case AIRFIELD:
		obj = s.navMan.Airfields[s.selectedObject]
	}

	recAnchorX := int32(rl.GetRenderWidth()) - state.Scale(state.SCREEN_MARGIN*2) - state.Scale(state.MENU_BUTTON_SIZE) - state.Scale(200)
	recAnchorY := state.Scale(state.SCREEN_MARGIN*2) + state.Scale(state.MENU_BUTTON_SIZE)
	outline := state.Scale(4)
	fontSize := state.Scale(18)
	rl.DrawRectangle(recAnchorX, recAnchorY, state.Scale(200), state.Scale(100), state.COLOR_UNSELECT)
	rl.DrawRectangle(recAnchorX+outline/2, recAnchorY+outline/2, state.Scale(200)-outline, state.Scale(100)-outline, rl.Black)
	rl.DrawText(obj.Name, recAnchorX+state.Scale(200/2)-rl.MeasureText(obj.Name, fontSize)/2, recAnchorY+outline*2, fontSize, state.COLOR_SELECT)

	rl.DrawText(strconv.FormatFloat(float64(obj.Latitude), 'f', 2, 64)+"°", recAnchorX+outline*2, recAnchorY+outline*3+fontSize, fontSize, state.COLOR_SELECT)
	longStr := strconv.FormatFloat(float64(obj.Longitude), 'f', 2, 64) + "°"
	rl.DrawText(longStr, recAnchorX+state.Scale(200)-outline*2-rl.MeasureText(longStr, fontSize), recAnchorY+outline*3+fontSize, fontSize, state.COLOR_SELECT)

	rl.DrawText("Dist:", recAnchorX+outline*2, recAnchorY+outline*4+fontSize*2, fontSize, state.COLOR_SELECT)
	distStr := strconv.FormatFloat(obj.dist, 'f', 3, 64) + "km"
	rl.DrawText(distStr, recAnchorX+state.Scale(200)-outline*2-rl.MeasureText(distStr, fontSize), recAnchorY+outline*4+fontSize*2, fontSize, state.COLOR_SELECT)

	rl.DrawText("Bear:", recAnchorX+outline*2, recAnchorY+outline*5+fontSize*3, fontSize, state.COLOR_SELECT)
	bearStr := strconv.FormatFloat(math.Mod(360+RadToDeg(obj.bearing), 360), 'f', 2, 64) + "°"
	rl.DrawText(bearStr, recAnchorX+state.Scale(200)-outline*2-rl.MeasureText(bearStr, fontSize), recAnchorY+outline*5+fontSize*3, fontSize, state.COLOR_SELECT)

	if obj.dist/float64(s.scale) > 0.75 {
		xOff := float32(math.Sin(math.Pi-obj.bearing+DegToRad(float64(state.GlobalFlightData.Heading))) * float64(state.ScaleF(150)))
		yOff := float32(math.Cos(math.Pi-obj.bearing+DegToRad(float64(state.GlobalFlightData.Heading))) * float64(state.ScaleF(150)))
		rectOut := rl.Rectangle{X: float32(anchorX) + xOff, Y: float32(anchorY) + yOff, Width: state.ScaleF(18), Height: state.ScaleF(18)}
		rectIn := rl.Rectangle{X: float32(anchorX) + xOff, Y: float32(anchorY) + yOff, Width: state.ScaleF(12), Height: state.ScaleF(12)}
		rl.DrawRectanglePro(rectOut, rl.Vector2{X: state.ScaleF(9), Y: state.ScaleF(9)}, float32(45+RadToDeg(obj.bearing)-float64(state.GlobalFlightData.Heading)), state.COLOR_SELECT)
		rl.DrawRectanglePro(rectIn, rl.Vector2{X: state.ScaleF(6), Y: state.ScaleF(6)}, float32(45+RadToDeg(obj.bearing)-float64(state.GlobalFlightData.Heading)), rl.Black)
	}
}

func (s _state) drawAirfields(anchorX, anchorY int32, width int32) {
	for i := 0; i < len(s.navMan.Airfields); i++ {
		var x, y, dist, bearing float64
		if s.useHaversine {
			x, y, dist, bearing = coordToDistanceHaversine(
				float64(state.GlobalFlightData.Latitude),
				float64(state.GlobalFlightData.Longitude),
				float64(s.navMan.Airfields[i].Latitude),
				float64(s.navMan.Airfields[i].Longitude),
				float64(state.GlobalFlightData.Heading),
			)
		} else {
			x, y, dist, bearing = coordToDistance(
				float64(state.GlobalFlightData.Latitude),
				float64(state.GlobalFlightData.Longitude),
				float64(s.navMan.Airfields[i].Latitude),
				float64(s.navMan.Airfields[i].Longitude),
				float64(state.GlobalFlightData.Heading),
			)
		}
		s.navMan.Airfields[i].dist = dist
		s.navMan.Airfields[i].bearing = bearing
		relX := (x / float64(s.scale)) * float64(width/4) * 3
		relY := (y / float64(s.scale)) * float64(width/4) * 3

		selected := false
		if s.selectedType == AIRFIELD && s.selectedObject == i {
			selected = true
		}
		s.navMan.Airfields[i].DrawNavObject(anchorX+int32(math.Round(relX)), anchorY+int32(math.Round(relY)), selected)
	}
}

func (s _state) drawNavPoints(anchorX, anchorY int32, width int32) {
	var lastX = 0.0
	var lastY = 0.0
	for i := 0; i < len(s.navMan.NavPoints); i++ {
		var x, y, dist, bearing float64
		if s.useHaversine {
			x, y, dist, bearing = coordToDistanceHaversine(
				float64(state.GlobalFlightData.Latitude),
				float64(state.GlobalFlightData.Longitude),
				float64(s.navMan.NavPoints[i].Latitude),
				float64(s.navMan.NavPoints[i].Longitude),
				float64(state.GlobalFlightData.Heading),
			)
		} else {
			x, y, dist, bearing = coordToDistance(
				float64(state.GlobalFlightData.Latitude),
				float64(state.GlobalFlightData.Longitude),
				float64(s.navMan.NavPoints[i].Latitude),
				float64(s.navMan.NavPoints[i].Longitude),
				float64(state.GlobalFlightData.Heading),
			)
		}
		s.navMan.NavPoints[i].dist = dist
		s.navMan.NavPoints[i].bearing = bearing
		relX := (x / float64(s.scale)) * float64(width/4) * 3
		relY := (y / float64(s.scale)) * float64(width/4) * 3
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
				COLOR_NAV_LINE,
			)
		}
		lastX = relX
		lastY = relY

		selected := false
		if s.selectedType == NAV_POINT && s.selectedObject == i {
			selected = true
		}
		s.navMan.NavPoints[i].DrawNavObject(anchorX+int32(math.Round(relX)), anchorY+int32(math.Round(relY)), selected)
	}
}
