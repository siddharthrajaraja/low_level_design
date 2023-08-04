package service

import (
	"github.com/siddharthrajaraja/low_level_design/trello/dtos"
)

type ICardService interface {
	CreateCard(name, description string) string
	AssignCardToUser(userId string, cardId string)
	UpdateCardAttributes(cadId string, card *dtos.Cards)
	ChangeListType(listType dtos.ListType, cardId string)
	GetCard(id string) *dtos.Cards
	DeleteCard(id string)
	CheckCardExists(id string) bool
}
