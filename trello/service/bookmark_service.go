package service

type IBookmarkService interface {
	AddCardToUser(userId, cardId string)
	AddListToUser(userId, listId string)
	AddBoardToUser(userId, boardId string)
}
