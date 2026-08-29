package stateNavigation

import (
	"fmt"
	"math"
	"strconv"

	"example.com/MFDTest/internal/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type _state struct {
	scale            int
	maxScale         int
	minScale         int
	navMan           *NavManager
	selectedNavPoint int
}

func Init(nav *NavManager) _state {
	return _state{
		scale:            30,
		maxScale:         130,
		minScale:         10,
		navMan:           nav,
		selectedNavPoint: 0,
	}
}

const (
	KEY_UP    = rl.KeyUp
	KEY_DOWN  = rl.KeyDown
	KEY_LEFT  = rl.KeyLeft
	KEY_RIGHT = rl.KeyRight
	KEY_APPLY = rl.KeyEnter
)

func (s _state) Input() state.State {
	if rl.IsKeyPressed(KEY_UP) && s.scale < s.maxScale {
		s.scale = s.scale + 20
	} else if rl.IsKeyPressed(KEY_DOWN) && s.scale > s.minScale {
		s.scale = s.scale - 20
	}
	return s
}

var COLOR_PLANE_SELF = rl.Color{R: 68, G: 190, B: 220, A: 255}
var COLOR_AIRPORT = rl.White
var COLOR_NAV_POINT = rl.Green
var COLOR_NAV_LINE = rl.Green

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

	headingFontSize := state.Scale(20)
	headingStr := strconv.FormatFloat(math.Round(float64(state.GlobalFlightData.Heading*100))/100, 'f', 2, 64) + "°"
	headingCenter := state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_SIZE) + (state.GetDisplayAreaWidth()-rl.MeasureText(headingStr, headingFontSize))/2
	rl.DrawRectangle(state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_SIZE)+(state.GetDisplayAreaWidth()-state.Scale(84))/2, state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_SIZE)-state.Scale(1), 84, headingFontSize+state.Scale(2), rl.Black)
	rl.DrawText(headingStr, headingCenter, state.Scale(state.SCREEN_MARGIN*2+state.MENU_BUTTON_SIZE), headingFontSize, state.COLOR_SELECT)
}

func (s _state) drawNavigation() {
	width := state.GetDisplayAreaWidth()
	anchorX := state.Scale(state.SCREEN_MARGIN*2) + state.Scale(state.MENU_BUTTON_SIZE) + (width / 2)
	anchorY := state.Scale(state.SCREEN_MARGIN*2) + state.Scale(state.MENU_BUTTON_SIZE) + (width * 3 / 4)

	s.drawNavPoints()

	planeSize := state.Scale(32)
	planeThickness := state.Scale(2)

	rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, state.ScaleF(150-1), state.ScaleF(150+1), 0, 360, 0, state.COLOR_SELECT)
	rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, state.ScaleF(150*2-1), state.ScaleF(150*2+1), 0, 360, 0, state.COLOR_SELECT)
	rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, state.ScaleF(150*3-1), state.ScaleF(150*3+1), 0, 360, 0, state.COLOR_SELECT)
	rl.DrawRing(rl.Vector2{X: float32(anchorX), Y: float32(anchorY)}, state.ScaleF(150*4-1), state.ScaleF(150*4+1), 0, 360, 0, state.COLOR_SELECT)

	rl.DrawRectangle(anchorX-(planeThickness/2), anchorY, planeThickness, planeSize, COLOR_PLANE_SELF)
	rl.DrawRectangle(anchorX-planeSize/2, anchorY+(planeSize/3)-(planeThickness/2), planeSize, planeThickness, COLOR_PLANE_SELF)
	rl.DrawRectangle(anchorX-planeSize/4, anchorY+(planeSize*7/8)-2, planeSize/2, planeThickness, COLOR_PLANE_SELF)
}

func (s _state) drawNavPoints() {
	for i := 0; i < len(s.navMan.NavPoints); i++ {
		x, y, dist, _ := coordToDistance(float64(state.GlobalFlightData.Latitude), float64(state.GlobalFlightData.Longitude), float64(s.navMan.NavPoints[i].Latitude), float64(s.navMan.NavPoints[i].Longitude), float64(state.GlobalFlightData.Heading))
		relX := x / 10_000 * float64(s.scale)
		relY := y / 10_000 * float64(s.scale)
		fmt.Printf("---------- Point (%f°, %f°) ----------\n", s.navMan.NavPoints[i].Latitude, s.navMan.NavPoints[i].Longitude)
		fmt.Printf("Distance %f (%fkm, %fkm)\n", dist, x, y)
		fmt.Printf("Pixel Space (%fpx, %fpx)\n", relX, relY)
		s.navMan.NavPoints[i].DrawNavObject(int32(math.Round(relX)), int32(math.Round(relY)))
	}
}
