package impl

import (
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/dtos"
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/service"
)

type FirstDiagonalValidator struct {
}

func (r *FirstDiagonalValidator) Validate(board *dtos.Board, x, y int, choice string) bool {

	i := x
	j := y
	for i >= 0 && j >= 0 {
		if board.Matrix[i][j] != choice {
			return false
		}
		i--
		j--
	}

	i = x
	j = y

	for i < board.N && j < board.N {

		if board.Matrix[i][j] != choice {
			return false
		}
		i++
		j++
	}
	return true
}

func NewFirstDiagonalValidator() service.IMoveValidator {

	return &FirstDiagonalValidator{}
}
