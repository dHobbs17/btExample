package constants

type State int

// make these shared constants
const (
	ShowMenu State = iota
	ShowLogin
	ShowGame
)
