package impl

import (
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/dtos"
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/service"
)

type RowValidator struct {
}

func (r *RowValidator) Validate(board *dtos.Board, x, y int, choice string) bool {

	for j := 0; j < board.N; j++ {
		if board.Matrix[x][j] != choice {
			return false
		}
	}
	return true
}

func NewRowValidator() service.IMoveValidator {

	return &RowValidator{}
}
