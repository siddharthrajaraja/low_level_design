package service

import "github.com/siddharthrajaraja/low_level_design/trello/dtos"

type IBoardService interface {
	CreateBoard(name string) string
	AddUserToMembersList(userId string, boardId string)
	RemoveUserFromMembersList(userId string, boardId string)
	UpdateBoardAttributes(boardId string, newBoard *dtos.Board)
	// Delete all list inside it and delete each card inside it
	DeleteBoard(id string)
	GetAllBoards() []*dtos.Board
	GetBoard(id string) *dtos.Board
}
