package models

// Definir el tipo de estado del auto
type AutoState int

const (
	StateEntering AutoState = iota
	StateParked
	StateExiting
)

type Auto struct {
	PosX       float64
	PosY       float64
	Dir        float64
	Cajon      int
	EnTransito bool
	State      AutoState
}

// Cambiar el estado del auto
func (a *Auto) ChangeState(state AutoState) {
	a.State = state
}
