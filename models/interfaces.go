package models

type Elevatorer interface {
	Move
	Move([]int)
}
type Move interface {
	moveDown(int)
	moveUp(int)
}
