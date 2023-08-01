package main

import (
	"fmt"
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/dtos"
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/service"
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/service/impl"
)

func main() {

	var player1Name, player2Name string
	var numberOfRows int

	fmt.Print("Enter number of rows in matrix :")
	fmt.Scanln(&numberOfRows)

	fmt.Print("Enter player1Name Name: ")
	fmt.Scanln(&player1Name)

	fmt.Print("Enter player2Name Name: ")
	fmt.Scanln(&player2Name)

	player1 := dtos.NewPlayer(player1Name)
	player2 := dtos.NewPlayer(player2Name)

	board := dtos.NewBord(numberOfRows)
	board.InitialiseBoard()
	board.PrintBoard("Printing board post initialisation")
	moveValidatorFactory := impl.NewMoveValidatorFactory()
	boardService := service.NewBoardServiceImpl(board, moveValidatorFactory)

	playService := service.NewPlayServiceImpl(player1, player2, boardService)
	resp := playService.Play()
	fmt.Printf("\n------GAME ENDED-----\n")
	fmt.Println(resp)
}
