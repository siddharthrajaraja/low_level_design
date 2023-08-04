package impl

import (
	"github.com/siddharthrajaraja/low_level_design/trello/dtos"
	"github.com/siddharthrajaraja/low_level_design/trello/service"
	"log"
)

type CardServiceImpl struct {
	userService service.IUserService
	CardMap     map[string]*dtos.Cards
}

func (c CardServiceImpl) DeleteCard(id string) {

	if !c.CheckCardExists(id) {
		delete(c.CardMap, id)
	}
}

func (c CardServiceImpl) CreateCard(name, description string) string {

	newCard := dtos.NewCards(name, description)
	c.CardMap[newCard.GetId()] = newCard
	return newCard.GetId()
}

func (c CardServiceImpl) AssignCardToUser(userId string, cardId string) {

	if !c.CheckCardExists(cardId) {
		log.Printf("Card not present. Please create a card with id %s\n", cardId)
		return
	}
	if !c.userService.CheckUserExists(userId) {
		log.Printf("User not present. Please create a User with id %s\n", userId)
	}

	c.CardMap[cardId].SetAssignedUser(userId)
}

func (c CardServiceImpl) UpdateCardAttributes(cardId string, newCard *dtos.Cards) {
	if !c.CheckCardExists(cardId) {
		log.Printf("Card not present. Please create a card with id %s\n", cardId)
		return
	}

	c.CardMap[cardId].SetStatus(newCard.GetStatus())
	c.CardMap[cardId].SetEta(newCard.GetEta())
	c.CardMap[cardId].SetCreatedAt(newCard.GetCreatedAt())
	c.CardMap[cardId].SetFinishedAt(newCard.GetFinishedAt())
	c.CardMap[cardId].SetAssignedUser(newCard.GetAssignedUser())
	c.CardMap[cardId].SetDescription(newCard.GetDescription())
	c.CardMap[cardId].SetName(newCard.GetName())
	c.CardMap[cardId].SetPriority(newCard.GetPriority())
}

func (c CardServiceImpl) ChangeListType(listType dtos.ListType, cardId string) {

	if !c.CheckCardExists(cardId) {
		log.Printf("Card not present. Please create a card with id %s\n", cardId)
		return
	}
	c.CardMap[cardId].SetStatus(listType)
}

func (c CardServiceImpl) GetCard(id string) *dtos.Cards {

	if !c.CheckCardExists(id) {
		log.Printf("Card not present. Please create a card with id %s\n", id)
		return nil
	}
	return c.CardMap[id]
}

func (c CardServiceImpl) CheckCardExists(id string) bool {

	_, isPresent := c.CardMap[id]
	return isPresent
}

func NewCardServiceImpl(userService service.IUserService) service.ICardService {

	cardMap := make(map[string]*dtos.Cards)
	return &CardServiceImpl{
		userService: userService,
		CardMap:     cardMap,
	}
}
