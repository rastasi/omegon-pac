package types

type Player struct {
	Name string
}

type State struct {
	CurrentField Field
}

type Grid struct {
	Fields []Field
	State  State
}

type Field struct {
	Name string
	X    int
	Y    int
}
