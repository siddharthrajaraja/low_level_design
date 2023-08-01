package service

import "github.com/siddharthrajaraja/low_level_design/tic_tac_toe/dtos"

type IMoveValidatorFactory interface {
	GetMoveValidator(dtos.ValidatorType) IMoveValidator
}

type IMoveValidator interface {
	Validate(board *dtos.Board, x int, y int, choice string) bool
}
