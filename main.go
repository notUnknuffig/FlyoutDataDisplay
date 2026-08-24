package main

import (
	"example.com/MFDTest/internal/menu"
	"example.com/MFDTest/internal/options"
	"example.com/MFDTest/internal/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	initWindow()
}

func initWindow() {
	rl.InitWindow(1024, 1024, "Multi-Function Display")
	defer rl.CloseWindow()
	rl.SetTargetFPS(30)

	var state state.State
	menu := menu.Menu{}
	config := options.Options{
		Resolution_x: 1024,
		Resolution_y: 1024,
		AspectRatio:  "1x1",
		Fullscreen:   false,
	}
	state = options.Init(config)

	for !rl.WindowShouldClose() {
		// State Input
		menu.Input()
		state = state.Input()

		// Visuals
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// Draw State
		menu.Draw()
		state.Draw()

		rl.EndDrawing()
	}
}
