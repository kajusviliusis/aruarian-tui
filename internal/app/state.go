package app

type AppState int

const (
	MenuState AppState = iota
	TodoState
	TimerState
)

