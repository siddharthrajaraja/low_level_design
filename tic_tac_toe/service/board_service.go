package service

import (
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/dtos"
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/utils"
)

type IBoardService interface {
	Validate(int, int, string) bool
	getMoveType(int, int) int
	MarkBoard(int, int, string)
	CheckBoardEmpty() bool
	IsBoardFilledAtPosition(int, int) bool
	PrintBoard(string)
	IsMoveValid(int, int) bool
}

type BoardServiceImpl struct {
	board *dtos.Board
}

func NewBoardServiceImpl(board *dtos.Board) IBoardService {

	return &BoardServiceImpl{
		board: board,
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

	if b.validateColumn(lastXMove, lastYMove, choice) || b.validateRow(lastXMove, lastYMove, choice) {
		return true
	}

	moveType := b.getMoveType(lastXMove, lastYMove)
	switch moveType {
	case utils.IntersectionDiagonalMove:
		return b.validateFirstDiagonalMove(lastXMove, lastYMove, choice) ||
			b.validateSecondDiagonalMove(lastXMove, lastYMove, choice)
	case utils.FirstDiagonalMove:
		return b.validateFirstDiagonalMove(lastXMove, lastYMove, choice)
	case utils.SecondDiagonalMove:
		return b.validateSecondDiagonalMove(lastXMove, lastYMove, choice)
	}

	return false
}

func (b *BoardServiceImpl) getMoveType(x, y int) int {

	if x == y {
		if b.board.N%2 != 0 && x == b.board.N/2 {
			return utils.IntersectionDiagonalMove
		}
		return utils.FirstDiagonalMove
	} else if x+y+1 == b.board.N {
		return utils.SecondDiagonalMove
	}
	return utils.NoDiagonalMove
}

func (b *BoardServiceImpl) validateFirstDiagonalMove(x, y int, choice string) bool {

	i := x
	j := y
	for i >= 0 && j >= 0 {
		if b.board.Matrix[i][j] != choice {
			return false
		}
		i--
		j--
	}

	i = x
	j = y

	for i < b.board.N && j < b.board.N {

		if b.board.Matrix[i][j] != choice {
			return false
		}
		i++
		j++
	}
	return true
}

func (b *BoardServiceImpl) validateSecondDiagonalMove(x, y int, choice string) bool {

	i := x
	j := y
	for i >= 0 && j < b.board.N {
		if b.board.Matrix[i][j] != choice {
			return false
		}
		i--
		j++
	}

	i = x
	j = y

	for i < b.board.N && j >= 0 {

		if b.board.Matrix[i][j] != choice {
			return false
		}
		i++
		j--
	}
	return true
}

func (b *BoardServiceImpl) validateRow(x, y int, choice string) bool {

	for j := 0; j < b.board.N; j++ {
		if b.board.Matrix[x][j] != choice {
			return false
		}
	}
	return true
}

func (b *BoardServiceImpl) validateColumn(x, y int, choice string) bool {

	for i := 0; i < b.board.N; i++ {
		if b.board.Matrix[i][y] != choice {
			return false
		}
	}
	return true
}
