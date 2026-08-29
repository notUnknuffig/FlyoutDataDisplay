package stateOptions

import (
	"fmt"

	"example.com/MFDTest/internal/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const OPTION_LENGTH = 4
const (
	OPTION_RESOLUTION   = 0
	OPTION_ASPECT_RATIO = 1
	OPTION_FULLSCREEN   = 2
	OPTION_APPLY        = 3
)
const (
	KEY_UP    = rl.KeyUp
	KEY_DOWN  = rl.KeyDown
	KEY_LEFT  = rl.KeyLeft
	KEY_RIGHT = rl.KeyRight
	KEY_APPLY = rl.KeyEnter
)

var aspectRatios []string = []string{
	"1x1",
	"3x4",
	"4x5",
	"9x16",
}

var scales []float32 = []float32{
	0.5,
	0.75,
	1,
	1.25,
	1.5,
}

var resolutions [][][]int = [][][]int{
	{{300, 300}, {512, 512}, {720, 720}, {910, 910}, {1080, 1080}},   // 1x1  1
	{{300, 400}, {512, 683}, {720, 960}, {910, 1213}, {1080, 1440}},  // 3x4  1.333333333
	{{300, 375}, {512, 640}, {720, 900}, {910, 1138}, {1080, 1350}},  // 4x5  1.25
	{{300, 533}, {512, 910}, {720, 1280}, {910, 1618}, {1080, 1920}}, // 9x16 1.777777778
}

/*
 * Options:
 * - Resolution
 * - Aspect Ratio
 * - Colorscheme
 * - Fullscreen / Windowed
 * - Apply
 */

type OptionState struct {
	options                 *state.Options
	selection               int // Option selected to Modify
	resolutionIndex         int // index
	aspectRatioIndex        int // index in reselutions
	previousFullscreenState bool
}

func Init(cfg *state.Options) OptionState {
	return OptionState{
		options:          cfg,
		selection:        0,
		resolutionIndex:  2,
		aspectRatioIndex: 0,
	}
}

func (s OptionState) Input() state.State {
	if rl.IsKeyPressed(KEY_DOWN) && s.selection < OPTION_LENGTH-1 {
		s.selection = s.selection + 1
	} else if rl.IsKeyPressed(KEY_UP) && s.selection > 0 {
		s.selection = s.selection - 1
	}

	switch s.selection {
	case OPTION_RESOLUTION:
		if s.options.Fullscreen {

		} else if rl.IsKeyPressed(KEY_LEFT) && s.resolutionIndex > 0 {
			s.resolutionIndex -= 1
			s.options.ResolutionX = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.ResolutionY = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]
		} else if rl.IsKeyPressed(KEY_RIGHT) && s.resolutionIndex < len(resolutions) {
			s.resolutionIndex += 1
			s.options.ResolutionX = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.ResolutionY = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]
		}
	case OPTION_ASPECT_RATIO:
		if s.options.Fullscreen {

		} else if rl.IsKeyPressed(KEY_LEFT) && s.aspectRatioIndex > 0 {
			s.aspectRatioIndex -= 1
			s.options.AspectRatio = aspectRatios[s.aspectRatioIndex]
			s.options.ResolutionX = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.ResolutionY = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]
		} else if rl.IsKeyPressed(KEY_RIGHT) && s.aspectRatioIndex < len(aspectRatios)-1 {
			s.aspectRatioIndex += 1
			s.options.AspectRatio = aspectRatios[s.aspectRatioIndex]
			s.options.ResolutionX = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.ResolutionY = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]
		}
	case OPTION_FULLSCREEN:
		if rl.IsKeyPressed(KEY_RIGHT) {
			s.options.Fullscreen = true
			monitor := rl.GetCurrentMonitor()
			s.options.ResolutionX = rl.GetMonitorWidth(monitor)
			s.options.ResolutionY = rl.GetMonitorHeight(monitor)

		} else if rl.IsKeyPressed(KEY_LEFT) {
			s.options.Fullscreen = false
			s.options.ResolutionX = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.ResolutionY = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]

		}
	case OPTION_APPLY:
		if rl.IsKeyPressed(KEY_APPLY) {
			if s.options.Fullscreen {
				rl.SetWindowSize(s.options.ResolutionX, s.options.ResolutionY)
				if !s.previousFullscreenState {
					rl.ToggleFullscreen()
				} // TODO: Figure out if this should be fullscreen
				s.previousFullscreenState = true
			} else {
				if s.previousFullscreenState {
					rl.ToggleFullscreen()
				} // TODO: Figure out if this should be fullscreen
				rl.SetWindowSize(s.options.ResolutionX, s.options.ResolutionY)
				s.options.Scale = scales[s.resolutionIndex]
				s.previousFullscreenState = false
			}
		}
	}
	return s
}

func (s OptionState) Draw() {
	s.drawOptions()
	s.drawButtons()
}

const BUTTON_LENGTH = 5

func (s OptionState) drawButtons() {
	state.DrawArrowButtons(0)
}

func (s OptionState) drawOptions() {
	var fontSize = state.Scale(20)
	var buttonHeight = state.Scale(32)
	var buttonWidth = state.Scale(240)
	var buttonMargin = (buttonHeight - fontSize) / 2
	var gap = state.Scale(48)
	var centerBoxX = (int32(rl.GetRenderWidth()) - buttonWidth) / 2
	var centerBoxY = (int32(rl.GetRenderHeight()) - buttonHeight - (OPTION_LENGTH * gap)) / 2
	for i := 0; i < OPTION_LENGTH; i++ {
		var textColor rl.Color
		var backgroundColor rl.Color
		if i == s.selection {
			textColor = state.COLOR_UNSELECT
			backgroundColor = state.COLOR_SELECT
		} else {
			textColor = state.COLOR_SELECT
			backgroundColor = state.COLOR_UNSELECT
		}
		rl.DrawRectangle(centerBoxX, centerBoxY+gap*int32(i), buttonWidth, buttonHeight, backgroundColor)
		switch i {
		case OPTION_RESOLUTION:
			rl.DrawText(fmt.Sprintf("Resolution: %dx%d", s.options.ResolutionX, s.options.ResolutionY), centerBoxX+buttonMargin, centerBoxY+gap*int32(i)+buttonMargin, fontSize, textColor)
		case OPTION_ASPECT_RATIO:
			rl.DrawText(fmt.Sprintf("Aspect Ratio %s", s.options.AspectRatio), centerBoxX+buttonMargin, centerBoxY+gap*int32(i)+buttonMargin, fontSize, textColor)
		case OPTION_FULLSCREEN:
			if s.options.Fullscreen {
				rl.DrawText("Fullscreen: on", centerBoxX+buttonMargin, centerBoxY+gap*int32(i)+buttonMargin, fontSize, textColor)
			} else {
				rl.DrawText("Fullscreen: off", centerBoxX+buttonMargin, centerBoxY+gap*int32(i)+buttonMargin, fontSize, textColor)
			}
		case OPTION_APPLY:
			rl.DrawText("Apply", centerBoxX+buttonMargin, centerBoxY+gap*int32(i)+buttonMargin, fontSize, textColor)
		}
	}
}
