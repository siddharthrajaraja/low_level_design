package impl

import (
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/dtos"
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/service"
)

type MoveValidatorFactory struct {
	factory map[dtos.ValidatorType]service.IMoveValidator
}

func NewMoveValidatorFactory() service.IMoveValidatorFactory {

	return &MoveValidatorFactory{
		factory: map[dtos.ValidatorType]service.IMoveValidator{
			dtos.RowValidator:            NewRowValidator(),
			dtos.ColumnValidator:         NewColumnValidator(),
			dtos.FirstDiagonalValidator:  NewFirstDiagonalValidator(),
			dtos.SecondDiagonalValidator: NewSecondDiagonalValidator(),
		},
	}
}

func (m *MoveValidatorFactory) GetMoveValidator(moveType dtos.ValidatorType) service.IMoveValidator {

	return m.factory[moveType]
}
