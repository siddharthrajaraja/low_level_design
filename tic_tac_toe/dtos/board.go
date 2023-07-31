package dtos

import (
	"fmt"
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/utils"
	"log"
)

type Board struct {
	N          int
	emptyCells int
	// singleton object
	Matrix [][]string
}

func NewBord(n int) *Board {

	return &Board{
		N:          n,
		emptyCells: n * n,
	}
}

func (b *Board) IsBoardEmpty() bool {

	return b.emptyCells == 0
}

func (b *Board) DecrementEmptyCells() {

	b.emptyCells--
}

func (b *Board) InitialiseBoard() {

	if b.Matrix != nil {
		log.Printf("Board already initialised")
		return
	}

	for i := 0; i < b.N; i++ {
		subMat := make([]string, 0)
		for j := 0; j < b.N; j++ {
			subMat = append(subMat, utils.HashTag)
		}
		b.Matrix = append(b.Matrix, subMat)
	}
}

func (b *Board) PrintBoard(msg string) {

	fmt.Println(msg)
	if b.Matrix != nil {
		for i := 0; i < b.N; i++ {
			for j := 0; j < b.N; j++ {
				fmt.Print(b.Matrix[i][j] + " ")
			}
			fmt.Println()
		}
	}
}
