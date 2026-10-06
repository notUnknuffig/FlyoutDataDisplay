package indicators

import (
	"math"

	"example.com/MFDTest/internal/navigation"
	"example.com/MFDTest/internal/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func DrawFlap(flapAngle float64, slatAngle float64, size float64, anchorX, anchorY int32) {
	var (
		wingCord      int32 = state.Scale(100)
		wingThickness       = state.Scale(20)
		wingRoundness       = state.ScaleF(7)
		flapLength          = state.Scale(50)
		padding             = state.Scale(10)
		slatRoundness       = state.ScaleF(7)
		slatLength          = state.Scale(20)
	)

	x := anchorX + padding + slatLength
	y := anchorY + padding + int32(float32(wingThickness)*0.5)

	// Wing
	rl.DrawRectangle(x, anchorY+padding, wingCord, wingThickness, state.COLOR_TEXT_SELECT)
	rl.DrawEllipse(x+int32(float64(wingCord)*0.5), anchorY+padding, float32(wingCord)*0.5, wingRoundness, state.COLOR_TEXT_SELECT)

	// Slat
	sinSlat := math.Sin(navigation.DegToRad(slatAngle))
	cosSlat := math.Cos(navigation.DegToRad(slatAngle))

	rl.DrawCircle(x, y, float32(wingThickness)*0.7, rl.Black)
	rl.DrawCircle(x, y, float32(wingThickness)*0.5, state.COLOR_TEXT_SELECT)

	rl.DrawCircle(x-int32(float64(slatLength)*cosSlat), y+int32(float64(slatLength)*sinSlat), slatRoundness, state.COLOR_TEXT_SELECT)
	e := rl.Vector2{X: float32(x) - float32(float64(wingThickness)*0.5*sinSlat), Y: float32(y) - float32(float64(wingThickness)*0.5*cosSlat)}
	f := rl.Vector2{X: float32(x) + float32(float64(wingThickness)*0.5*sinSlat), Y: float32(y) + float32(float64(wingThickness)*0.5*cosSlat)}
	g := rl.Vector2{
		X: float32(x) - float32(float64(slatLength)*cosSlat) - float32(float64(slatRoundness)*sinSlat),
		Y: float32(y) + float32(float64(slatLength)*sinSlat) - float32(float64(slatRoundness)*cosSlat),
	}
	h := rl.Vector2{
		X: float32(x) - float32(float64(slatLength)*cosSlat) + float32(float64(slatRoundness)*sinSlat),
		Y: float32(y) + float32(float64(slatLength)*sinSlat) + float32(float64(slatRoundness)*cosSlat),
	}
	rl.DrawTriangle(e, g, f, state.COLOR_TEXT_SELECT)
	rl.DrawTriangle(f, g, h, state.COLOR_TEXT_SELECT)

	// Flap

	sinFlap := math.Sin(navigation.DegToRad(flapAngle))
	cosFlap := math.Cos(navigation.DegToRad(flapAngle))

	rl.DrawCircle(x+wingCord, y, float32(wingThickness)*0.7, rl.Black)
	rl.DrawCircleSector(rl.Vector2{X: float32(x + wingCord), Y: float32(y)}, float32(wingThickness)*0.5, float32(270+flapAngle), float32(90+flapAngle), 0, state.COLOR_TEXT_SELECT)

	a := rl.Vector2{X: float32(x+wingCord) - float32(float64(wingThickness)*0.5*sinFlap), Y: float32(y) + float32(float64(wingThickness)*0.5*cosFlap)}
	b := rl.Vector2{X: float32(x+wingCord) + float32(float64(wingThickness)*0.5*sinFlap), Y: float32(y) - float32(float64(wingThickness)*0.5*cosFlap)}
	c := rl.Vector2{X: float32(x+wingCord) + float32(float64(flapLength)*cosFlap), Y: float32(y) + float32(float64(flapLength)*sinFlap)}
	rl.DrawTriangle(c, b, a, state.COLOR_TEXT_SELECT)

}
