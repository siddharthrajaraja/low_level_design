package impl

import (
	"github.com/siddharthrajaraja/low_level_design/trello/dtos"
	"github.com/siddharthrajaraja/low_level_design/trello/service"
	"log"
)

type ListServiceImpl struct {
	ListMap     map[string]*dtos.List
	cardService service.ICardService
}

func (l ListServiceImpl) CreateList(name string) string {

	list := dtos.NewList(name, dtos.New)
	l.ListMap[list.GetListId()] = list
	return list.GetListId()
}

func (l ListServiceImpl) GetList(id string) *dtos.List {

	var list *dtos.List
	if !l.CheckListExists(id) {

		log.Printf("List not present. Please create a list with id %s\n", id)
		return list
	}
	return l.ListMap[id]
}

func (l ListServiceImpl) UpdateListAttributes(listId string, updatedListObj *dtos.List) {

	list := l.GetList(listId)

	if list != nil {

		for _, everyCardId := range updatedListObj.GetCardIds() {
			if l.cardService.CheckCardExists(everyCardId) {
				list.SetCardId(everyCardId)
			}
		}
		list.SetListType(updatedListObj.GetListType())
	}
}

func (l ListServiceImpl) DeleteList(id string) {

	if !l.CheckListExists(id) {
		log.Printf("List not present. Please create a list with id %s\n", id)
		return
	}
	cardIdList := l.ListMap[id].GetCardIds()

	for _, everyCardId := range cardIdList {
		l.cardService.DeleteCard(everyCardId)
	}

	delete(l.ListMap, id)
}

func (l ListServiceImpl) CheckListExists(id string) bool {

	_, isPresent := l.ListMap[id]
	return isPresent
}

func NewListServiceImpl(cardService service.ICardService) service.IListService {

	listMap := make(map[string]*dtos.List)
	return &ListServiceImpl{
		ListMap:     listMap,
		cardService: cardService,
	}
}
