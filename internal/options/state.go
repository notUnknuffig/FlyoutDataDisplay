package options

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

var resolutions [][][]int = [][][]int{
	{{300, 300}, {512, 512}, {1024, 1024}, {1920, 1920}, {2440, 2440}}, // 1x1
	{{300, 400}, {540, 720}, {1080, 1440}, {1440, 1920}, {1920, 2560}}, // 3x4
	{{320, 400}, {576, 720}, {1080, 1350}, {1536, 1920}, {2048, 2560}}, // 4x5
	{{270, 480}, {405, 720}, {576, 1024}, {810, 1440}, {1080, 1920}},   // 9x16
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
	options                 Options
	selection               int // Option selected to Modify
	resolutionIndex         int // index
	aspectRatioIndex        int // index in reselutions
	previousFullscreenState bool
}

type Options struct {
	Resolution_x int
	Resolution_y int
	AspectRatio  string
	Fullscreen   bool
}

func Init(cfg Options) OptionState {
	return OptionState{
		options:          cfg,
		selection:        0,
		resolutionIndex:  0,
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
			s.options.Resolution_x = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.Resolution_y = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]
		} else if rl.IsKeyPressed(KEY_RIGHT) && s.resolutionIndex < len(resolutions)-1 {
			s.resolutionIndex += 1
			s.options.Resolution_x = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.Resolution_y = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]
		}
	case OPTION_ASPECT_RATIO:
		if s.options.Fullscreen {

		} else if rl.IsKeyPressed(KEY_LEFT) && s.aspectRatioIndex > 0 {
			s.aspectRatioIndex -= 1
			s.options.AspectRatio = aspectRatios[s.aspectRatioIndex]
			s.options.Resolution_x = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.Resolution_y = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]
		} else if rl.IsKeyPressed(KEY_RIGHT) && s.aspectRatioIndex < len(aspectRatios)-1 {
			s.aspectRatioIndex += 1
			s.options.AspectRatio = aspectRatios[s.aspectRatioIndex]
			s.options.Resolution_x = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.Resolution_y = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]
		}
	case OPTION_FULLSCREEN:
		if rl.IsKeyPressed(KEY_RIGHT) {
			s.options.Fullscreen = true
			monitor := rl.GetCurrentMonitor()
			s.options.Resolution_x = rl.GetMonitorWidth(monitor)
			s.options.Resolution_y = rl.GetMonitorHeight(monitor)

		} else if rl.IsKeyPressed(KEY_LEFT) {
			s.options.Fullscreen = false
			s.options.Resolution_x = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.Resolution_y = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]

		}
	case OPTION_APPLY:
		if rl.IsKeyPressed(KEY_APPLY) {
			if s.options.Fullscreen {
				rl.SetWindowSize(s.options.Resolution_x, s.options.Resolution_y)
				if !s.previousFullscreenState {
					rl.ToggleFullscreen()
				} // TODO: Figure out if this should be fullscreen
				s.previousFullscreenState = true
			} else {
				if s.previousFullscreenState {
					rl.ToggleFullscreen()
				} // TODO: Figure out if this should be fullscreen
				rl.SetWindowSize(s.options.Resolution_x, s.options.Resolution_y)
				s.previousFullscreenState = false
			}
		}
	}
	return s
}

var COLOR_UNSELECT = rl.DarkGreen
var COLOR_SELECT = rl.Green
var COLOR_TEXT_SELECT = rl.White
var COLOR_TEXT_UNSELECT = rl.Gray
var SCREEN_MARGIN = 20

func (s OptionState) Draw() {
	var fontSize = int32(25)
	s.drawOptions(fontSize)
	s.drawButtons(fontSize)
}

const BUTTON_LENGTH = 5

func (s OptionState) drawButtons(fontSize int32) {
	var buttonSize = int32(60) // Square
	// rl.DrawLine(0, 0, int32(rl.GetRenderHeight()), int32(rl.GetRenderWidth()), rl.White)
	// rl.DrawLine(int32(rl.GetRenderHeight()), 0, 0, int32(rl.GetRenderWidth()), rl.White)
	for i := 0; i < BUTTON_LENGTH; i++ {
		var baseX = SCREEN_MARGIN
		var baseY = SCREEN_MARGIN + ((rl.GetRenderHeight()-(2*SCREEN_MARGIN))/(BUTTON_LENGTH+1))*(i+1) - int(buttonSize/2)
		rl.DrawRectangle(int32(baseX)-5, int32(baseY)-5, buttonSize, buttonSize, rl.Gray)
		a := rl.Vector2{X: float32(baseX), Y: float32(baseY)}
		b := rl.Vector2{X: float32(baseX + int(buttonSize)), Y: float32(baseY)}
		c := rl.Vector2{X: float32(baseX + int(buttonSize/2)), Y: float32(baseY + int(buttonSize))}
		rl.DrawTriangle(a, b, c, COLOR_SELECT)
	}
}

func (s OptionState) drawOptions(fontSize int32) {
	var buttonWidth = int32(300)
	var buttonHeight = int32(40)
	var centerBoxX = (int32(rl.GetRenderWidth()) - buttonWidth) / 2
	var centerBoxY = (int32(rl.GetRenderHeight()) - buttonHeight - (OPTION_LENGTH * 50)) / 2
	for i := 0; i < OPTION_LENGTH; i++ {
		var textColor rl.Color
		var backgroundColor rl.Color
		if i == s.selection {
			textColor = COLOR_TEXT_SELECT
			backgroundColor = COLOR_SELECT
		} else {
			textColor = COLOR_TEXT_UNSELECT
			backgroundColor = COLOR_UNSELECT
		}
		rl.DrawRectangle(centerBoxX, centerBoxY+int32(50*i), buttonWidth, buttonHeight, backgroundColor)
		switch i {
		case OPTION_RESOLUTION:
			rl.DrawText(fmt.Sprintf("Resolution: %dx%d", s.options.Resolution_x, s.options.Resolution_y), centerBoxX+5, centerBoxY+int32(50*i)+5, fontSize, textColor)
		case OPTION_ASPECT_RATIO:
			rl.DrawText(fmt.Sprintf("Aspect Ratio %s", s.options.AspectRatio), centerBoxX+5, centerBoxY+int32(50*i)+5, fontSize, textColor)
		case OPTION_FULLSCREEN:
			if s.options.Fullscreen {
				rl.DrawText("Fullscreen: on", centerBoxX+5, centerBoxY+int32(50*i)+5, fontSize, textColor)
			} else {
				rl.DrawText("Fullscreen: off", centerBoxX+5, centerBoxY+int32(50*i)+5, fontSize, textColor)
			}
		case OPTION_APPLY:
			rl.DrawText("Apply", centerBoxX+5, centerBoxY+int32(50*i)+5, fontSize, textColor)
		}
	}
}
