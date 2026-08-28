package stateSystems

import "example.com/MFDTest/internal/state"

type _state struct {
}

func Init() _state {
	return _state{}
}

func (s _state) Input() state.State {
	return s
}

func (s _state) Draw() {
	if state.GlobalFlightData == nil {
		state.DrawNoData()
		return
	}
}
