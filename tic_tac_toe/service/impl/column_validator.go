package impl

import (
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/dtos"
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/service"
)

type ColumnValidator struct {
}

func (r *ColumnValidator) Validate(board *dtos.Board, x, y int, choice string) bool {

	for i := 0; i < board.N; i++ {
		if board.Matrix[i][y] != choice {
			return false
		}
	}
	return true
}

func NewColumnValidator() service.IMoveValidator {

	return &ColumnValidator{}
}
