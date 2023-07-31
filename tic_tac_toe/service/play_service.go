package service

import (
	"fmt"
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/dtos"
	"github.com/siddharthrajaraja/low_level_design/tic_tac_toe/utils"
	"math/rand"
)

type IPlayService interface {
	Play() string
	toss()
}

type PlayServiceImpl struct {
	// true -> player1 and false -> player2
	startingPlayer bool

	player1      *dtos.Player
	player2      *dtos.Player
	boardService IBoardService
}

func (p *PlayServiceImpl) Play() string {

	var winningPlayer *dtos.Player

	p.toss()
	for !p.boardService.CheckBoardEmpty() {
		var xMove int
		var yMove int
		currentPlayer := p.player1

		if !p.startingPlayer {
			currentPlayer = p.player2
		}

		for {
			fmt.Printf(utils.PlayerTurnToPick, currentPlayer.Name, currentPlayer.GetChoice())
			fmt.Scanf("%d", &xMove)
			fmt.Scanf("%d", &yMove)
			if !p.boardService.IsMoveValid(xMove, yMove) || p.boardService.IsBoardFilledAtPosition(xMove, yMove) {
				fmt.Printf(utils.InvalidMove, currentPlayer.Name)
			} else {
				break
			}
		}

		p.boardService.MarkBoard(xMove, yMove, currentPlayer.GetChoice())
		p.boardService.PrintBoard("Printing board")

		if p.boardService.Validate(xMove, yMove, currentPlayer.GetChoice()) {
			winningPlayer = currentPlayer
			break
		}

		p.startingPlayer = !p.startingPlayer
	}

	if winningPlayer != nil {
		return fmt.Sprintf(utils.GameWinningString, winningPlayer.Name)
	}
	return utils.Draw
}

func (p *PlayServiceImpl) toss() {

	var tossWinningPlayer *dtos.Player
	var tossLosingPlayer *dtos.Player
	var choiceOfWinningPlayer string
	var choiceOfLosingPlayer string

	if rand.Intn(1000)%2 == 0 {
		p.startingPlayer = true
		tossWinningPlayer = p.player1
		tossLosingPlayer = p.player2
	} else {
		p.startingPlayer = false
		tossWinningPlayer = p.player2
		tossLosingPlayer = p.player1
	}

	fmt.Printf(utils.TossWinningString, tossWinningPlayer.Name, utils.X, utils.O)
	fmt.Scanln(&choiceOfWinningPlayer)

	if choiceOfWinningPlayer == utils.X {
		choiceOfLosingPlayer = utils.O
	} else if choiceOfWinningPlayer == utils.O {
		choiceOfLosingPlayer = utils.X
	} else {
		choiceOfWinningPlayer = utils.X
		choiceOfLosingPlayer = utils.O
		fmt.Printf(utils.AutoAssigningChoicesFallback, tossWinningPlayer.Name, choiceOfWinningPlayer)
	}

	fmt.Printf(utils.AssignOtherChoiceString, tossLosingPlayer.Name, choiceOfLosingPlayer)

	tossWinningPlayer.AssignChoice(choiceOfWinningPlayer)
	tossLosingPlayer.AssignChoice(choiceOfLosingPlayer)
}

func NewPlayServiceImpl(player1 *dtos.Player, player2 *dtos.Player, boardService IBoardService) IPlayService {

	return &PlayServiceImpl{
		player1:      player1,
		player2:      player2,
		boardService: boardService,
	}
}
