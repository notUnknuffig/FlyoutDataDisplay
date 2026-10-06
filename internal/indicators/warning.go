package indicators

import (
	"example.com/MFDTest/internal/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	COLOR_INFO      = rl.White
	COLOR_WARNING   = rl.Orange
	COLOR_EMERGENCY = rl.Red
)

func DrawLandingGear(anchorX, anchorY int32) {
	var (
		width          = state.Scale(44)
		wheelSizeInner = state.ScaleF(4)
		wheelSize      = state.ScaleF(16)
		wheelOffset    = state.Scale(4)
		pistonWidth    = state.Scale(8)
		fontSize       = state.Scale(16)
	)
	// rl.DrawRectangle(anchorX, anchorY, 44, 44, rl.Blue)

	rl.DrawCircle(anchorX+(width/2), anchorY+(width/2)+wheelOffset, wheelSizeInner, COLOR_WARNING)
	rl.DrawRing(rl.Vector2{X: float32(anchorX) + float32(width)/2, Y: float32(anchorY) + float32(width)/2 + float32(wheelOffset)}, wheelSizeInner+state.ScaleF(4), wheelSize, 0, 360, 0, COLOR_WARNING)
	rl.DrawRectangle(anchorX+(width/2)-(pistonWidth/2), anchorY, pistonWidth, width-int32(wheelSize*2)-state.Scale(4), COLOR_WARNING)

	margin := rl.MeasureText("LG", fontSize)
	rl.DrawText("LG", anchorX+state.Scale(22)-(margin/2), anchorY+state.Scale(60)-fontSize, fontSize, COLOR_INFO)
}

// 16 *

func DrawAirbrake(anchorX, anchorY int32) {
	var (
		offsetX  = state.Scale(16)
		offsetY  = state.Scale(16)
		size     = state.Scale(6)
		fontSize = state.Scale(16)
	)

	if state.GlobalFlightData.Airbrake < 0.5 {
		rl.SetLineWidth(state.ScaleF(2))
		j := -1
		for i := 0; i < 9; i++ {
			if i%3 == 0 {
				j++
			}
			rl.DrawCircleLines(anchorX+offsetX*int32(i%3)+size, anchorY+offsetY*int32(j)+size, float32(size)-state.ScaleF(2), COLOR_INFO)
		}
	} else if state.GlobalFlightData.Airbrake < 0.9 {
		j := -1
		for i := 0; i < 9; i++ {
			if i%3 == 0 {
				j++
			}
			rl.DrawCircle(anchorX+offsetX*int32(i%3)+size, anchorY+offsetY*int32(j)+size, float32(size), COLOR_INFO)
		}
	} else {
		rl.SetLineWidth(state.ScaleF(4))

		rl.DrawRectangleLines(anchorX+state.Scale(2), anchorY+state.Scale(2), size*2+offsetX*2-state.Scale(4), size*2+offsetX*2-state.Scale(4), COLOR_INFO)

		for i := 0; i < 9; i++ {
			var a, b rl.Vector2
			if i <= 3 {
				a = rl.Vector2{X: float32(anchorX) + 10*float32(i+1) + state.ScaleF(4), Y: float32(anchorY) + state.ScaleF(4)}
				b = rl.Vector2{X: float32(anchorX) + state.ScaleF(4), Y: float32(anchorY) + state.ScaleF(4) + 10*float32(i+1)}
			} else {
				a = rl.Vector2{X: float32(anchorX) + state.ScaleF(4) + state.ScaleF(40) - 10*float32((i%4)), Y: float32(anchorY) + state.ScaleF(4) + state.ScaleF(40)}
				b = rl.Vector2{X: float32(anchorX) + state.ScaleF(4) + state.ScaleF(40), Y: float32(anchorY) + state.ScaleF(4) + state.ScaleF(40) - 10*float32((i%4))}
			}
			rl.DrawLineV(a, b, COLOR_INFO)
		}
	}
	margin := rl.MeasureText("SB", fontSize)
	rl.DrawText("SB", anchorX+state.Scale(22)-(margin/2), anchorY+state.Scale(60)-fontSize, fontSize, COLOR_INFO)
}

func DrawHook(anchorX, anchorY int32) {
	var (
		fontSize = state.Scale(16)
	)
	margin := rl.MeasureText("HK", fontSize)
	rl.DrawText("HK", anchorX+state.Scale(22)-(margin/2), anchorY+state.Scale(60)-fontSize, fontSize, COLOR_INFO)
}

func DrawApu(anchorX, anchorY int32) {
	var (
		fontSize = state.Scale(16)
	)
	margin := rl.MeasureText("APU", fontSize)
	rl.DrawText("APU", anchorX+state.Scale(22)-(margin/2), anchorY+state.Scale(60)-fontSize, fontSize, COLOR_INFO)
}
