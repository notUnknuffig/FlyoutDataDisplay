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
	OPTION_UNITS        = 3
)

var aspectRatios []string = []string{
	"1x1",
	"4x5",
	"3x4",
	"9x16",
}

var scales []float32 = []float32{
	0.45,
	0.75,
	0.9,
	1,
	1.25,
	1.5,
}

var resolutions [][][]int = [][][]int{
	{{300, 300}, {512, 512}, {720, 720}, {800, 800}, {910, 910}, {1080, 1080}},    // 1x1  1
	{{300, 375}, {512, 640}, {720, 900}, {800, 1000}, {910, 1138}, {1080, 1350}},  // 4x5  1.25
	{{300, 400}, {512, 683}, {720, 960}, {800, 1067}, {910, 1213}, {1080, 1440}},  // 3x4  1.333333333
	{{300, 533}, {512, 910}, {720, 1280}, {800, 1422}, {910, 1618}, {1080, 1920}}, // 9x16 1.777777778
}

const (
	UNITS_SI    = 0
	UNITS_AM    = 1
	UNITS_SI_AV = 2
	UNITS_AM_AV = 3
	UNITS_RU    = 4
)

var units []state.Units = []state.Units{
	{ // SI
		DistanceUnit:       "km",
		DistanceConversion: 1,
		SpeedUnit:          "m/s",
		SpeedConversion:    1,
		HightUnit:          "m",
		HightConversion:    1,
		ClimbUnit:          "m/s",
		ClimbConversion:    1,
		HeatUnit:           "k",
		HeatConversion:     1,
		HeatFreezingPoint:  0,
		WeightUnit:         "kg",
		WeightConversion:   1,
		TonUnit:            "t",
		TonConversion:      1_000,
		FlowRateUnit:       "kg/s",
		FlowRateConversion: 1,
	},
	{ // Imperial
		DistanceUnit:       "mi",
		DistanceConversion: 0.6213712,
		SpeedUnit:          "mph",
		SpeedConversion:    2.236936,
		HightUnit:          "ft",
		HightConversion:    3.280839895,
		ClimbUnit:          "ft/s",
		ClimbConversion:    3.280839895,
		HeatUnit:           "°F",
		HeatConversion:     1.8,
		HeatFreezingPoint:  -459.67,
		WeightUnit:         "lbs",
		WeightConversion:   2.204623,
		TonUnit:            "t",
		TonConversion:      907.18,
		FlowRateUnit:       "lbs/s",
		FlowRateConversion: 2.204623,
	},
	{ // SI Aviation
		DistanceUnit:       "nm",
		DistanceConversion: 0.5399568,
		SpeedUnit:          "kt",
		SpeedConversion:    1.943844,
		HightUnit:          "m",
		HightConversion:    1,
		ClimbUnit:          "m/s",
		ClimbConversion:    1,
		HeatUnit:           "°C",
		HeatConversion:     1,
		HeatFreezingPoint:  -273.15,
		WeightUnit:         "kg",
		WeightConversion:   1,
		TonUnit:            "t",
		TonConversion:      1_000,
		FlowRateUnit:       "kg/s",
		FlowRateConversion: 1,
	},
	{ // Imperial Aviation
		DistanceUnit:       "nm",
		DistanceConversion: 0.5399568,
		SpeedUnit:          "kt",
		SpeedConversion:    1.943844,
		HightUnit:          "ft",
		HightConversion:    3.280839895,
		ClimbUnit:          "ft/s",
		ClimbConversion:    3.280839895,
		HeatUnit:           "°F",
		HeatConversion:     1.8,
		HeatFreezingPoint:  -459.67,
		WeightUnit:         "lbs",
		WeightConversion:   0.453592,
		TonUnit:            "t",
		TonConversion:      907.18,
		FlowRateUnit:       "lbs/s",
		FlowRateConversion: 0.453592,
	},
	{ // Russian (SI)
		DistanceUnit:       "km",
		DistanceConversion: 1,
		SpeedUnit:          "km/h",
		SpeedConversion:    3.6,
		HightUnit:          "m",
		HightConversion:    1,
		ClimbUnit:          "m/s",
		ClimbConversion:    1,
		HeatUnit:           "°C",
		HeatConversion:     1,
		HeatFreezingPoint:  -273.15,
		WeightUnit:         "kg",
		WeightConversion:   1,
		TonUnit:            "t",
		TonConversion:      1_000,
		FlowRateUnit:       "kg/s",
		FlowRateConversion: 1,
	},
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
	unitIndex               int
}

func Init(cfg *state.Options) OptionState {
	cfg.Units = &units[0]
	return OptionState{
		options:          cfg,
		selection:        0,
		resolutionIndex:  2,
		aspectRatioIndex: 0,
	}
}

func (s OptionState) Input() state.State {
	if rl.IsKeyPressed(state.KEY_F7) && s.selection < OPTION_LENGTH-1 {
		s.selection = s.selection + 1
	} else if rl.IsKeyPressed(state.KEY_F6) && s.selection > 0 {
		s.selection = s.selection - 1
	}

	switch s.selection {
	case OPTION_RESOLUTION:
		if s.options.Fullscreen {

		} else if rl.IsKeyPressed(state.KEY_F11) && s.resolutionIndex > 0 {
			s.resolutionIndex -= 1
			s.options.ResolutionX = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.ResolutionY = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]
		} else if rl.IsKeyPressed(state.KEY_F12) && s.resolutionIndex < len(resolutions[0])-1 {
			s.resolutionIndex += 1
			s.options.ResolutionX = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.ResolutionY = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]
		}
	case OPTION_ASPECT_RATIO:
		if s.options.Fullscreen {

		} else if rl.IsKeyPressed(state.KEY_F11) && s.aspectRatioIndex > 0 {
			s.aspectRatioIndex -= 1
			s.options.AspectRatio = aspectRatios[s.aspectRatioIndex]
			s.options.ResolutionX = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.ResolutionY = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]
		} else if rl.IsKeyPressed(state.KEY_F12) && s.aspectRatioIndex < len(aspectRatios)-1 {
			s.aspectRatioIndex += 1
			s.options.AspectRatio = aspectRatios[s.aspectRatioIndex]
			s.options.ResolutionX = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.ResolutionY = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]
		}
	case OPTION_FULLSCREEN:
		if rl.IsKeyPressed(state.KEY_F12) {
			s.options.Fullscreen = true
			monitor := rl.GetCurrentMonitor()
			s.options.ResolutionX = rl.GetMonitorWidth(monitor)
			s.options.ResolutionY = rl.GetMonitorHeight(monitor)

		} else if rl.IsKeyPressed(state.KEY_F11) {
			s.options.Fullscreen = false
			s.options.ResolutionX = resolutions[s.aspectRatioIndex][s.resolutionIndex][0]
			s.options.ResolutionY = resolutions[s.aspectRatioIndex][s.resolutionIndex][1]

		}
	case OPTION_UNITS:
		if rl.IsKeyPressed(state.KEY_F11) && s.unitIndex > 0 {
			s.unitIndex -= 1
			s.options.Units = &units[s.unitIndex]
		} else if rl.IsKeyPressed(state.KEY_F12) && s.unitIndex < len(units)-1 {
			s.unitIndex += 1
			s.options.Units = &units[s.unitIndex]
		}
	}
	if rl.IsKeyPressed(state.KEY_F8) {
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
	return s
}

func (s OptionState) Draw() {
	s.drawOptions()
	s.drawButtons()
}

func (s OptionState) drawButtons() {
	state.DrawArrowButtonsHorizontal(0)
	state.DrawArrowButtonsVertical(11 - 1)
	state.DrawButton("ENT", 8-1)
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
		case OPTION_UNITS:
			switch s.unitIndex {
			case UNITS_SI:
				rl.DrawText("Units: Scientific", centerBoxX+buttonMargin, centerBoxY+gap*int32(i)+buttonMargin, fontSize, textColor)
			case UNITS_AM:
				rl.DrawText("Units: Imperial", centerBoxX+buttonMargin, centerBoxY+gap*int32(i)+buttonMargin, fontSize, textColor)
			case UNITS_SI_AV:
				rl.DrawText("Units: Aviation (SI)", centerBoxX+buttonMargin, centerBoxY+gap*int32(i)+buttonMargin, fontSize, textColor)
			case UNITS_AM_AV:
				rl.DrawText("Units: Aviation (Imp)", centerBoxX+buttonMargin, centerBoxY+gap*int32(i)+buttonMargin, fontSize, textColor)
			case UNITS_RU:
				rl.DrawText("Units: European", centerBoxX+buttonMargin, centerBoxY+gap*int32(i)+buttonMargin, fontSize, textColor)
			}
		}
	}
}
