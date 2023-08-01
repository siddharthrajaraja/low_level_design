package dtos

type MoveType int

// Enum to get all move types
const (
	IntersectionDiagonalMove MoveType = 2
	FirstDiagonalMove        MoveType = 0
	SecondDiagonalMove       MoveType = 1
	NoDiagonalMove           MoveType = -1
)
