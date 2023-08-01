package service

import (
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/dtos"
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/utils"
)

type IBoardService interface {
	Validate(int, int, string) bool
	getMoveType(int, int) dtos.MoveType
	MarkBoard(int, int, string)
	CheckBoardEmpty() bool
	IsBoardFilledAtPosition(int, int) bool
	PrintBoard(string)
	IsMoveValid(int, int) bool
}

type BoardServiceImpl struct {
	board         *dtos.Board
	moveValidator IMoveValidatorFactory
}

func NewBoardServiceImpl(board *dtos.Board, moveValidator IMoveValidatorFactory) IBoardService {

	return &BoardServiceImpl{
		board:         board,
		moveValidator: moveValidator,
	}
}

func (b *BoardServiceImpl) IsMoveValid(x, y int) bool {

	return x >= 0 && x < b.board.N && y >= 0 && y < b.board.N
}

func (b *BoardServiceImpl) PrintBoard(msg string) {

	b.board.PrintBoard(msg)
}

func (b *BoardServiceImpl) IsBoardFilledAtPosition(x, y int) bool {

	return b.board.Matrix[x][y] != utils.HashTag
}

func (b *BoardServiceImpl) CheckBoardEmpty() bool {

	return b.board.IsBoardEmpty()
}

func (b *BoardServiceImpl) MarkBoard(xMove, yMove int, choice string) {

	b.board.Matrix[xMove][yMove] = choice
	b.board.DecrementEmptyCells()
}

func (b *BoardServiceImpl) Validate(lastXMove, lastYMove int, choice string) bool {

	if b.moveValidator.GetMoveValidator(dtos.RowValidator).Validate(b.board, lastXMove, lastYMove, choice) ||
		b.moveValidator.GetMoveValidator(dtos.ColumnValidator).Validate(b.board, lastXMove, lastYMove, choice) {
		return true
	}

	moveType := b.getMoveType(lastXMove, lastYMove)

	switch moveType {
	case dtos.IntersectionDiagonalMove:
		return b.moveValidator.GetMoveValidator(dtos.FirstDiagonalValidator).Validate(b.board, lastXMove, lastYMove, choice) ||
			b.moveValidator.GetMoveValidator(dtos.SecondDiagonalValidator).Validate(b.board, lastXMove, lastYMove, choice)
	case dtos.FirstDiagonalMove:
		return b.moveValidator.GetMoveValidator(dtos.FirstDiagonalValidator).Validate(b.board, lastXMove, lastYMove,
			choice)
	case dtos.SecondDiagonalMove:
		return b.moveValidator.GetMoveValidator(dtos.SecondDiagonalValidator).Validate(b.board, lastXMove, lastYMove,
			choice)
	}

	return false
}

func (b *BoardServiceImpl) getMoveType(x, y int) dtos.MoveType {

	if x == y {
		if b.board.N%2 != 0 && x == b.board.N/2 {
			return dtos.IntersectionDiagonalMove
		}
		return dtos.FirstDiagonalMove
	} else if x+y+1 == b.board.N {
		return dtos.SecondDiagonalMove
	}
	return dtos.NoDiagonalMove
}
