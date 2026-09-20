package app

import "example.com/MFDTest/internal/state"

func (a App) readSettings() state.Options {
	return state.Options{
		ResolutionX: 800,
		ResolutionY: 800,
		AspectRatio: "1x1",
		Fullscreen:  false,
		Scale:       1,
	}
}

func (a App) writeSettings() {

}
