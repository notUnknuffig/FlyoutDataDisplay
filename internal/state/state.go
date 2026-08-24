package state

type State interface {
	Input() State
	Draw()
}
