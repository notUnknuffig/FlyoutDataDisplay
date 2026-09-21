package weapons

import (
	"math"

	"example.com/MFDTest/internal/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type _state struct {
	missileThickness int
	missileInfo      *state.Missile
}

func Init() _state {
	return _state{
		missileThickness: 12,
		missileInfo:      nil,
	}
}

func (s _state) Input() state.State {
	return s
}

func (s _state) Draw() {
	if state.GlobalFlightData == nil {
		state.DrawNoData()
		return
	}

	s.DrawWingShape()
}

func (s _state) DrawWingShape() {
	anchorX := int32(rl.GetRenderWidth() / 2)
	anchorY := int32(rl.GetRenderWidth()/2) - state.Scale(100)
	fuselageWidth := state.Scale(80)
	wingSweep := state.Scale(80)
	wingWidth := state.Scale(275)

	rl.SetLineWidth(state.ScaleF(3))

	s.missileInfo = nil
	s.drawPayload(anchorX, anchorY)
	s.drawPayloadInfo()

	// Draw Generic Plane
	rl.DrawLine(anchorX-wingWidth, anchorY+wingSweep, anchorX-fuselageWidth/2, anchorY, state.COLOR_SELECT)
	rl.DrawLine(anchorX-fuselageWidth/2, anchorY-wingSweep/2, anchorX-fuselageWidth/2, anchorY+wingSweep, state.COLOR_SELECT)
	rl.DrawLine(anchorX+wingWidth, anchorY+wingSweep, anchorX+fuselageWidth/2, anchorY, state.COLOR_SELECT)
	rl.DrawLine(anchorX+fuselageWidth/2, anchorY-wingSweep/2, anchorX+fuselageWidth/2, anchorY+wingSweep, state.COLOR_SELECT)
}

func (s _state) drawPayload(anchorX, anchorY int32) {
	missile := len(state.GlobalFlightData.Missiles)
	if missile > 5 {
		missile = 5
	}
	for i := 0; i < missile; i++ {
		s.drawMissile(anchorX-state.Scale(75+(50*i)), anchorX+state.Scale(75+(50*i)), anchorY+state.Scale(12+(17*i)), state.GlobalFlightData.Missiles[i])
	}
}

func (s _state) drawWeaponStats(anchorX, anchorY int32, count int) {
	missileCount := count
	for i := 0; i < missileCount; i++ {
		row := int(math.Floor(float64(i) / 2.0))
		offset := -1
		if i%2 == 0 {
			offset = 1
		}
		rl.DrawCircle(anchorX+state.Scale(6*offset), anchorY+state.Scale(12*row), state.ScaleF(4), state.COLOR_TEXT_SELECT)
	}
}

func (s _state) drawMissile(anchorXRight, anchorXLeft, anchorY int32, missile state.Missile) {
	countLeft := int(math.Ceil(float64(missile.Count) / 2))
	countRight := int(math.Floor(float64(missile.Count) / 2))
	switch missile.Type {
	case "Infrared Missile":
		s.drawIRMissile(anchorXLeft, anchorY, true)
		s.drawIRMissile(anchorXRight, anchorY, false)
	case "Unguided Missile":
		s.drawBomb(anchorXLeft, anchorY, true)
		s.drawBomb(anchorXRight, anchorY, false)
	default:
		s.drawRadarMissile(anchorXLeft, anchorY, true)
		s.drawRadarMissile(anchorXRight, anchorY, false)
	}
	s.drawWeaponStats(anchorXLeft, anchorY+state.Scale(20), countLeft)
	s.drawWeaponStats(anchorXRight, anchorY+state.Scale(20), countRight)
	if missile.Name == state.GlobalFlightData.ActiveMissile {
		s.missileInfo = &missile
		yOffset := int(math.Ceil(float64(missile.Count) / 4))
		s.drawActiveMark(anchorXLeft, anchorXRight, anchorY+state.Scale(20*yOffset))
	}
}

func (s _state) drawActiveMark(anchorXRight, anchorXLeft, anchorY int32) {
	rl.DrawRectangle(anchorXLeft-state.Scale(15), anchorY, state.Scale(30), state.Scale(4), state.COLOR_SELECT)
	rl.DrawRectangle(anchorXRight-state.Scale(15), anchorY, state.Scale(30), state.Scale(4), state.COLOR_SELECT)
}

// Sharp Missile Head
// Backplaced Fins
func (s _state) drawRadarMissile(anchorX, anchorY int32, mirrored bool) {
	// rl.DrawLine(anchorX+(s.missileThickness/2), anchorY-state.Scale(2), anchorX-(s.missileThickness/2), anchorY+state.Scale(2), state.COLOR_SELECT)
	missileFinLength := state.Scale(8)
	mirroredInv := int32(1)
	if mirrored {
		mirroredInv = -1
	}

	rl.DrawLine(anchorX-(state.Scale(s.missileThickness)/2), anchorY+state.Scale(2)*mirroredInv, anchorX-(state.Scale(s.missileThickness)/2), anchorY-state.Scale(40), state.COLOR_SELECT)
	rl.DrawLine(anchorX+(state.Scale(s.missileThickness)/2), anchorY-state.Scale(2)*mirroredInv, anchorX+(state.Scale(s.missileThickness)/2), anchorY-state.Scale(40), state.COLOR_SELECT)
	rl.DrawTriangleLines(
		rl.Vector2{X: float32(anchorX-(state.Scale(s.missileThickness)/2)) - state.ScaleF(1), Y: float32(anchorY - state.Scale(40))},
		rl.Vector2{X: float32(anchorX+(state.Scale(s.missileThickness)/2)) + state.ScaleF(1), Y: float32(anchorY - state.Scale(40))},
		rl.Vector2{X: float32(anchorX), Y: float32(anchorY - state.Scale(60))},
		state.COLOR_SELECT,
	)

	// Fins
	rl.DrawTriangleLines(
		rl.Vector2{X: float32(anchorX-(state.Scale(s.missileThickness)/2)) - state.ScaleF(1), Y: float32(anchorY - state.Scale(10))},
		rl.Vector2{X: float32(anchorX-(state.Scale(s.missileThickness)/2)) - state.ScaleF(1), Y: float32(anchorY - state.Scale(20))},
		rl.Vector2{X: float32(anchorX-(state.Scale(s.missileThickness)/2)-missileFinLength) - state.ScaleF(1), Y: float32(anchorY - state.Scale(10))},
		state.COLOR_SELECT,
	)
	rl.DrawTriangleLines(
		rl.Vector2{X: float32(anchorX+(state.Scale(s.missileThickness)/2)) + state.ScaleF(1), Y: float32(anchorY - state.Scale(20))},
		rl.Vector2{X: float32(anchorX+(state.Scale(s.missileThickness)/2)) + state.ScaleF(1), Y: float32(anchorY - state.Scale(10))},
		rl.Vector2{X: float32(anchorX+(state.Scale(s.missileThickness)/2)+missileFinLength) + state.ScaleF(1), Y: float32(anchorY - state.Scale(10))},
		state.COLOR_SELECT,
	)
}

// Thick Body
// No Fins
func (s _state) drawBomb(anchorX, anchorY int32, mirrored bool) {
	// rl.DrawLine(anchorX+(state.Scale(s.missileThickness)/2), anchorY-state.Scale(2), anchorX-(state.Scale(s.missileThickness)/2), anchorY+state.Scale(2), state.COLOR_SELECT)
	missileFinLength := state.Scale(8)
	mirroredInv := int32(1)
	if mirrored {
		mirroredInv = -1
	}

	rl.DrawLine(anchorX-(state.Scale(s.missileThickness)/2), anchorY+state.Scale(2)*mirroredInv, anchorX-(state.Scale(s.missileThickness)/2), anchorY-state.Scale(40), state.COLOR_SELECT)
	rl.DrawLine(anchorX+(state.Scale(s.missileThickness)/2), anchorY-state.Scale(2)*mirroredInv, anchorX+(state.Scale(s.missileThickness)/2), anchorY-state.Scale(40), state.COLOR_SELECT)
	rl.DrawTriangleLines(
		rl.Vector2{X: float32(anchorX-(state.Scale(s.missileThickness)/2)) - state.ScaleF(1), Y: float32(anchorY - state.Scale(40))},
		rl.Vector2{X: float32(anchorX+(state.Scale(s.missileThickness)/2)) + state.ScaleF(1), Y: float32(anchorY - state.Scale(40))},
		rl.Vector2{X: float32(anchorX), Y: float32(anchorY - state.Scale(60))},
		state.COLOR_SELECT,
	)

	// Fins
	rl.DrawTriangleLines(
		rl.Vector2{X: float32(anchorX-(state.Scale(s.missileThickness)/2)) - state.ScaleF(1), Y: float32(anchorY - state.Scale(10))},
		rl.Vector2{X: float32(anchorX-(state.Scale(s.missileThickness)/2)) - state.ScaleF(1), Y: float32(anchorY - state.Scale(20))},
		rl.Vector2{X: float32(anchorX-(state.Scale(s.missileThickness)/2)-missileFinLength) - state.ScaleF(1), Y: float32(anchorY - state.Scale(10))},
		state.COLOR_SELECT,
	)
	rl.DrawTriangleLines(
		rl.Vector2{X: float32(anchorX+(state.Scale(s.missileThickness)/2)) + state.ScaleF(1), Y: float32(anchorY - state.Scale(20))},
		rl.Vector2{X: float32(anchorX+(state.Scale(s.missileThickness)/2)) + state.ScaleF(1), Y: float32(anchorY - state.Scale(10))},
		rl.Vector2{X: float32(anchorX+(state.Scale(s.missileThickness)/2)+missileFinLength) + state.ScaleF(1), Y: float32(anchorY - state.Scale(10))},
		state.COLOR_SELECT,
	)
}

// Round Seeker Head
// Frontplaced Fins
func (s _state) drawIRMissile(anchorX, anchorY int32, mirrored bool) {
	// rl.DrawLine(anchorX+(state.Scale(s.missileThickness)/2), anchorY-state.Scale(2), anchorX-(state.Scale(s.missileThickness)/2), anchorY+state.Scale(2), state.COLOR_SELECT)
	missileFinLength := state.Scale(8)
	mirroredInv := int32(1)
	if mirrored {
		mirroredInv = -1
	}

	rl.DrawLine(anchorX-(state.Scale(s.missileThickness)/2), anchorY+state.Scale(2)*mirroredInv, anchorX-(state.Scale(s.missileThickness)/2), anchorY-state.Scale(40), state.COLOR_SELECT)
	rl.DrawLine(anchorX+(state.Scale(s.missileThickness)/2), anchorY-state.Scale(2)*mirroredInv, anchorX+(state.Scale(s.missileThickness)/2), anchorY-state.Scale(40), state.COLOR_SELECT)
	rl.DrawRing(
		rl.Vector2{X: float32(anchorX), Y: float32(anchorY - state.Scale(40))},
		float32(state.Scale(s.missileThickness)/2)-state.ScaleF(1.5),
		float32(state.Scale(s.missileThickness)/2)+state.ScaleF(1.5),
		180,
		360,
		0,
		state.COLOR_SELECT,
	)

	// Fins
	rl.DrawTriangleLines(
		rl.Vector2{X: float32(anchorX-(state.Scale(s.missileThickness)/2)) - state.ScaleF(1), Y: float32(anchorY - state.Scale(20))},
		rl.Vector2{X: float32(anchorX-(state.Scale(s.missileThickness)/2)) - state.ScaleF(1), Y: float32(anchorY - state.Scale(30))},
		rl.Vector2{X: float32(anchorX-(state.Scale(s.missileThickness)/2)-missileFinLength) - state.ScaleF(1), Y: float32(anchorY - state.Scale(20))},
		state.COLOR_SELECT,
	)
	rl.DrawTriangleLines(
		rl.Vector2{X: float32(anchorX+(state.Scale(s.missileThickness)/2)) + state.ScaleF(1), Y: float32(anchorY - state.Scale(30))},
		rl.Vector2{X: float32(anchorX+(state.Scale(s.missileThickness)/2)) + state.ScaleF(1), Y: float32(anchorY - state.Scale(20))},
		rl.Vector2{X: float32(anchorX+(state.Scale(s.missileThickness)/2)+missileFinLength) + state.ScaleF(1), Y: float32(anchorY - state.Scale(20))},
		state.COLOR_SELECT,
	)
	rl.DrawLine(anchorX-(state.Scale(s.missileThickness)/2), anchorY-state.Scale(36), anchorX+(state.Scale(s.missileThickness)/2), anchorY-state.Scale(36), state.COLOR_SELECT)
}

func (s _state) drawPayloadInfo() {

}
