package app

import (
	"fmt"
	"math"

	"example.com/MFDTest/internal/navigation"
	"example.com/MFDTest/internal/state"
	"example.com/MFDTest/internal/states/attitude"
	"example.com/MFDTest/internal/states/engine"
	"example.com/MFDTest/internal/states/fuel"
	"example.com/MFDTest/internal/states/navigator"
	"example.com/MFDTest/internal/states/stateOptions"
	"example.com/MFDTest/internal/states/systems"
	"example.com/MFDTest/internal/states/weapons"
	rl "github.com/gen2brain/raylib-go/raylib"
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
	rl.SetExitKey(-1)
	icon := rl.LoadImage("icon.png")
	rl.SetWindowIcon(*icon)

	config := a.readSettings()

	var airFields []navigation.MappedObject
	// var objects []navigation.MappedObject
	var err error
	airFields, _, err = readAreaData()
	if err != nil {
		fmt.Print("WARNING: Unable to display custom airfields. Only loading default airfield coordinates.")
		airFields = []navigation.MappedObject{
			{
				Name:      "Default Airfield",
				Latitude:  -35.604,
				Longitude: -51.8061,
				Heading:   0,
				Allied:    true,
				Type:      navigation.AIRFIELD,
			},
			{
				Name:      "Desert Airfield",
				Latitude:  11.6763,
				Longitude: -63.3045,
				Heading:   309,
				Allied:    true,
				Type:      navigation.AIRFIELD,
			},
		}
	}
	nav := navigation.NavManager{
		NavPoints:      []navigation.MappedObject{},
		Airfields:      airFields,
		UseHaversine:   false,
		SelectedObject: -1,
		SelectedType:   navigation.NAV_POINT,
	}
	state.GlobalOptions = &config
	var fuelState state.State = fuel.Init()
	a.availableStates = []state.State{
		attitude.Init(&nav),
		engine.Init(fuelState),
		navigator.Init(&nav),
		systems.Init(),
		weapons.Init(),
		stateOptions.Init(&config),
	}

	a.state = a.availableStates[Options]

	// Startup Animation
	var sideLength = float32(state.Scale(200))
	var cos = float32(math.Cos(1.333333333333333*math.Pi)) * sideLength
	var sin = float32(math.Sin(1.333333333333333*math.Pi)) * sideLength

	var center = rl.Vector2{X: float32(rl.GetRenderHeight() / 2), Y: float32(rl.GetRenderWidth() / 2)}
	var centerHigh = rl.Vector2Add(center, rl.Vector2{X: 0, Y: -sideLength})
	var leftLow = rl.Vector2Add(center, rl.Vector2{X: sin, Y: -cos})
	var rightLow = rl.Vector2Add(center, rl.Vector2{X: -sin, Y: -cos})

	var i = 1
	for !rl.WindowShouldClose() {
		// State Input
		a.state = a.Input()
		a.state = a.state.Input()

		// Visuals
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		if i < 20 {
			var alpha = float32(math.Min(float64(i)/15, 1))
			rl.DrawTriangle(center, centerHigh, leftLow, rl.Fade(rl.Red, alpha))
			rl.DrawTriangle(leftLow, centerHigh, rl.Vector2Add(leftLow, rl.Vector2{X: 0, Y: -sideLength}), rl.Fade(rl.Red, alpha))

			rl.DrawTriangle(centerHigh, center, rightLow, rl.Fade(rl.Blue, alpha))
			rl.DrawTriangle(centerHigh, rightLow, rl.Vector2Add(rightLow, rl.Vector2{X: 0, Y: -sideLength}), rl.Fade(rl.Blue, alpha))

			rl.DrawTriangle(rightLow, center, leftLow, rl.Fade(rl.Green, alpha))
			rl.DrawTriangle(rightLow, leftLow, rl.Vector2Add(center, rl.Vector2{X: 0, Y: sideLength}), rl.Fade(rl.Green, alpha))
		} else if i < 30 {

		} else {
			// Draw State
			a.state.Draw()
			a.Draw()
		}

		rl.EndDrawing()
		i++
	}

	a.writeSettings()
}

func (a App) Input() state.State {
	switch rl.GetKeyPressed() {
	case state.KEY_F1:
		return a.availableStates[Attidude]
	case state.KEY_F2:
		return a.availableStates[Engine]
	case state.KEY_F3:
		return a.availableStates[Navigation]
	case state.KEY_F4:
		return a.availableStates[Systems]
	case state.KEY_F5:
		return a.availableStates[Weapons]
	}
	return a.state
}

func (a App) Draw() {
	// state.DrawMarginBox(true)
	// for i := 0; i < 20; i++ {
	// 	state.DrawButton(strconv.FormatInt(int64(i+1), 10), i)
	// }
	state.DrawButton("ATT", Attidude)
	state.DrawButton("ENG", Engine)
	state.DrawButton("NAV", Navigation)
	state.DrawButton("SYS", Systems)
	state.DrawButton("WPN", Weapons)
}
