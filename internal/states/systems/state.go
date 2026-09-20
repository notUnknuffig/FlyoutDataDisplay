package systems

import (
	"example.com/MFDTest/internal/state"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type _state struct {
	optionState state.State
}

func Init(optState state.State) _state {
	return _state{
		optionState: optState,
	}
}

func (s _state) Input() state.State {
	if rl.IsKeyPressed(state.KEY_F8) {
		return s.optionState
	}
	return s
}

func (s _state) Draw() {
	state.DrawButton("OPT", 8-1)
	if state.GlobalFlightData == nil {
		state.DrawNoData()
		return
	}
}
