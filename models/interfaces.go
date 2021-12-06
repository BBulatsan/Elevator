package models

type Elevators interface {
	Move
	Move([]int)
}
type Move interface {
	moveDown(int)
	moveUp(int)
}
