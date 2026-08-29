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
	COLOR_NAVPOINT = rl.DarkGreen
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
	case AIRCRAFT:
	case NAV_POINT:
		color = COLOR_NAV_POINT
		rl.DrawRectangle(x-state.Scale(8), y-state.Scale(8), state.Scale(16), state.Scale(16), color)
		rl.DrawRectangle(x-state.Scale(4), y-state.Scale(4), state.Scale(8), state.Scale(8), rl.Black)
	case TURNING_POINT:
		color = COLOR_NAV_POINT
		rl.DrawRing(rl.Vector2{X: float32(x), Y: float32(y)}, float32(state.Scale(2)), float32(state.Scale(4)), 0, 360, -1, color)
	}
}
