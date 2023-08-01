package impl

import (
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/dtos"
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/service"
)

type SecondDiagonalValidator struct {
}

func (r *SecondDiagonalValidator) Validate(board *dtos.Board, x, y int, choice string) bool {

	i := x
	j := y
	for i >= 0 && j < board.N {
		if board.Matrix[i][j] != choice {
			return false
		}
		i--
		j++
	}

	i = x
	j = y

	for i < board.N && j >= 0 {

		if board.Matrix[i][j] != choice {
			return false
		}
		i++
		j--
	}
	return true
}

func NewSecondDiagonalValidator() service.IMoveValidator {

	return &SecondDiagonalValidator{}
}
