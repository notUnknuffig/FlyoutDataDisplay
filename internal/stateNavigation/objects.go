package stateNavigation

import (
	"example.com/MFDTest/internal/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ObjectType int

const (
	AIRFIELD ObjectType = iota
	AIRCRAFT
	NAV_POINT
	TURNING_POINT
)

var (
	COLOR_NAVPOINT = rl.White
	COLOR_ALLIED   = rl.White
	COLOR_ENEMY    = rl.Orange
	COLOR_NAV_LINE = rl.White
)

type MappedObjects struct {
	Latitude  float32
	Longitude float32
	Heading   float32
	Allied    bool
	Type      ObjectType
}

func (n MappedObjects) DrawNavObject(x, y int32) {
	var color rl.Color
	switch n.Type {
	case AIRFIELD:
		var color rl.Color
		if n.Allied {
			color = COLOR_ALLIED
		} else {
			color = COLOR_ENEMY
		}
		height := state.ScaleF(32)
		width := state.ScaleF(12)
		rl.DrawRectanglePro(
			rl.NewRectangle(float32(x), float32(y), width, height),
			rl.Vector2{X: width / 2, Y: height / 2},
			n.Heading-state.GlobalFlightData.Heading,
			color,
		)
		height = state.ScaleF(32 - 4)
		width = state.ScaleF(12 - 4)
		rl.DrawRectanglePro(
			rl.NewRectangle(float32(x), float32(y), width, height),
			rl.Vector2{X: (width) / 2, Y: (height) / 2},
			n.Heading-state.GlobalFlightData.Heading,
			rl.Black,
		)
		// rl.DrawRing(rl.Vector2{float32(x), float32(y)}, 8, 12, 0, 360, 0, state.COLOR_SELECT)
	case AIRCRAFT:

	case NAV_POINT:
		color = COLOR_NAVPOINT
		rl.DrawRectangle(x-state.Scale(8), y-state.Scale(8), state.Scale(16), state.Scale(16), color)
		rl.DrawRectangle(x-state.Scale(5), y-state.Scale(5), state.Scale(10), state.Scale(10), rl.Black)
	case TURNING_POINT:
		color = COLOR_NAVPOINT
		rl.DrawRing(rl.Vector2{X: float32(x), Y: float32(y)}, state.ScaleF(4), state.ScaleF(6), 0, 360, -1, color)
	}
}
