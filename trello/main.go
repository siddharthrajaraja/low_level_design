package main

import (
	"github.com/siddharthrajaraja/low_level_design/trello/dtos"
	"github.com/siddharthrajaraja/low_level_design/trello/service/impl"
)

func main() {

	userService := impl.NewUserServiceImpl()
	cardService := impl.NewCardServiceImpl(userService)

	card1 := cardService.CreateCard("card1", "card1 description")

	listService := impl.NewListServiceImpl(cardService)
	boardService := impl.NewBoardServiceImpl(userService, listService)

	boardId := boardService.CreateBoard("Siddharth Board")
	list1 := listService.CreateList("List 1")

	listService.UpdateListAttributes(list1, &dtos.List{
		ListType: dtos.New,
		CardIds:  []string{card1},
	})

	list2 := listService.CreateList("List 2")

	updatedBoard := &dtos.Board{
		Privacy: dtos.Private,
		Lists:   []string{list1, list2},
	}
	boardService.UpdateBoardAttributes(boardId, updatedBoard)

	userId1 := userService.CreateUser("siddharth", "siddharthraja9849@gmail.com")

	updatedBoard.Members = []string{userId1}
	boardService.UpdateBoardAttributes(boardId, updatedBoard)
	userService.CreateUser("siddharth", "siddharthraja9849@gmail.com")

}
