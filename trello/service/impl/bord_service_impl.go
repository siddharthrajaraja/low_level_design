package impl

import (
	"github.com/siddharthrajaraja/low_level_design/trello/dtos"
	"github.com/siddharthrajaraja/low_level_design/trello/service"
	"log"
)

type BoardServiceImpl struct {
	BoardMap    map[string]*dtos.Board
	userService service.IUserService
	listService service.IListService
}

func (b BoardServiceImpl) CreateBoard(name string) {

	board := dtos.NewBoard(name)
	b.BoardMap[board.GetBoardId()] = board
}

func (b BoardServiceImpl) AddUserToMembersList(userId string, boardId string) {

	if !b.checkBoardPresent(boardId) {
		log.Printf("Board not present. Please create a board with id %s\n", boardId)
		return
	}

	if !b.userService.CheckUserExists(userId) {
		log.Printf("User not present.Please create a user with id %s\n", userId)
	}

	members := b.BoardMap[boardId].GetMembersList()
	members = append(members, userId)
	b.BoardMap[boardId].SetMembersList(members)
}

func (b BoardServiceImpl) RemoveUserFromMembersList(userId string, boardId string) {
	if !b.checkBoardPresent(boardId) {
		log.Printf("Board not present. Please create a board with id %s\n", boardId)
		return
	}

	if !b.userService.CheckUserExists(userId) {
		log.Printf("User not present.Please create a user with id %s\n", userId)
	}

	var newMembers []string

	members := b.BoardMap[boardId].GetMembersList()
	members = append(members, userId)

	for _, everyMember := range members {
		if everyMember != userId {
			newMembers = append(newMembers, everyMember)
		}
	}

	b.BoardMap[boardId].SetMembersList(newMembers)
}

func (b BoardServiceImpl) UpdateBoardAttributes(boardId string, newBoard *dtos.Board) {

	if !b.checkBoardPresent(boardId) {
		log.Printf("Board not present. Please create a board with id %s\n", boardId)
		return
	}

	board := b.GetBoard(boardId)

	if board != nil {

		for _, everyListId := range newBoard.GetList() {
			if b.listService.CheckListExists(everyListId) {
				board.SetList(everyListId)
			}

		}

		for _, everyListId := range newBoard.GetMembersList() {
			if b.userService.CheckUserExists(everyListId) {
				board.SetMember(everyListId)
			}

		}

		board.SetPrivacy(newBoard.GetPrivacy())
	}
}

func (b BoardServiceImpl) DeleteBoard(id string) {

	if !b.checkBoardPresent(id) {
		log.Printf("Board not present. Please create a board with id %s\n", id)
		return
	}

	listIds := b.BoardMap[id].GetList()

	for _, everyListId := range listIds {
		b.listService.DeleteList(everyListId)
	}

	delete(b.BoardMap, id)
}

func (b BoardServiceImpl) GetAllBoards() []*dtos.Board {
	var allBoards []*dtos.Board

	for _, boardObj := range b.BoardMap {
		allBoards = append(allBoards, boardObj)
	}

	return allBoards
}

func (b BoardServiceImpl) GetBoard(id string) *dtos.Board {

	return b.BoardMap[id]
}

func (b BoardServiceImpl) checkBoardPresent(id string) bool {

	_, isPresent := b.BoardMap[id]
	return isPresent
}

func NewBoardServiceImpl(userService service.IUserService, listService service.IListService) service.IBoardService {

	boardMap := make(map[string]*dtos.Board)
	return &BoardServiceImpl{
		BoardMap:    boardMap,
		listService: listService,
		userService: userService,
	}
}
