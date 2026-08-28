package app

import (
	"example.com/MFDTest/internal/state"
	"example.com/MFDTest/internal/stateAttitude"
	"example.com/MFDTest/internal/stateEngine"
	"example.com/MFDTest/internal/stateNavigation"
	"example.com/MFDTest/internal/stateOptions"
	"example.com/MFDTest/internal/stateSystems"
	"example.com/MFDTest/internal/stateWeapons"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	KEY_F1  = rl.KeyF1
	KEY_F2  = rl.KeyF2
	KEY_F3  = rl.KeyF3
	KEY_F4  = rl.KeyF4
	KEY_F5  = rl.KeyF5
	KEY_F6  = rl.KeyF6
	KEY_F7  = rl.KeyF7
	KEY_F8  = rl.KeyF8
	KEY_F9  = rl.KeyF9
	KEY_F10 = rl.KeyF10
	KEY_F11 = rl.KeyF11
	KEY_F12 = rl.KeyF12
	KEY_F13 = rl.KeyF12 + 1
	KEY_F14 = rl.KeyF12 + 2
	KEY_F15 = rl.KeyF12 + 3
	KEY_F16 = rl.KeyF12 + 4
	KEY_F17 = rl.KeyF12 + 5
	KEY_F18 = rl.KeyF12 + 6
	KEY_F19 = rl.KeyF12 + 7
	KEY_F20 = rl.KeyF12 + 8
)

const (
	Attidude   = 0
	Engine     = 1
	Navigation = 2
	Systems    = 3
	Weapons    = 4
	Options    = 5
)

type App struct {
	state           state.State
	availableStates []state.State
}

func (a App) Init() {
	rl.InitWindow(720, 720, "Multi-Function Display")
	defer rl.CloseWindow()
	rl.SetTargetFPS(30)

	config := state.Options{
		Resolution_x: 720,
		Resolution_y: 720,
		AspectRatio:  "1x1",
		Fullscreen:   false,
		Scale:        1,
	}
	state.GlobalOptions = &config
	a.availableStates = []state.State{
		stateAttitude.Init(),
		stateEngine.Init(),
		stateNavigation.Init(),
		stateSystems.Init(),
		stateWeapons.Init(),
		stateOptions.Init(&config),
	}

	a.state = a.availableStates[Options]

	for !rl.WindowShouldClose() {
		// State Input
		a.state = a.Input()
		a.state = a.state.Input()

		// Visuals
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// Draw State
		a.Draw()
		a.state.Draw()

		rl.EndDrawing()
	}
}

func (a App) Input() state.State {
	switch rl.GetKeyPressed() {
	case KEY_F1:
		return a.availableStates[Attidude]
	case KEY_F2:
		return a.availableStates[Engine]
	case KEY_F3:
		return a.availableStates[Navigation]
	case KEY_F4:
		return a.availableStates[Systems]
	case KEY_F5:
		return a.availableStates[Weapons]
	}
	return a.state
}

func (a App) Draw() {
	var buttonHeight = state.Scale(36) // Square
	var buttonWidth = state.Scale(64)
	var fontSize = state.Scale(24)
	var buttonMarginX = state.Scale(8)
	var buttonMarginY = (buttonHeight - fontSize) / 2
	for i := 0; i < state.BUTTON_LENGTH; i++ {
		var colorBackground = state.COLOR_UNSELECT
		var colorText = state.COLOR_TEXT_UNSELECT
		var baseX = state.SCREEN_MARGIN + ((rl.GetRenderWidth()-(2*state.SCREEN_MARGIN))/(state.BUTTON_LENGTH+1))*(i+1) - int(buttonWidth/2)
		var baseY = rl.GetRenderHeight() - state.SCREEN_MARGIN - int(buttonHeight)
		rl.DrawRectangle(int32(baseX), int32(baseY), buttonWidth, buttonHeight, colorBackground)
		switch i {
		case Attidude:
			rl.DrawText("Att", int32(baseX)+buttonMarginX, int32(baseY)+buttonMarginY, fontSize, colorText)
		case Engine:
			rl.DrawText("Eng", int32(baseX)+buttonMarginX, int32(baseY)+buttonMarginY, fontSize, colorText)
		case Navigation:
			rl.DrawText("Nav", int32(baseX)+buttonMarginX, int32(baseY)+buttonMarginY, fontSize, colorText)
		case Systems:
			rl.DrawText("Sys", int32(baseX)+buttonMarginX, int32(baseY)+buttonMarginY, fontSize, colorText)
		case Weapons:
			rl.DrawText("Wpn", int32(baseX)+buttonMarginX, int32(baseY)+buttonMarginY, fontSize, colorText)
		}
	}
}
