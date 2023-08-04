package service

import "github.com/siddharthrajaraja/low_level_design/trello/dtos"

type IListService interface {
	CreateList(name string) string
	GetList(id string) *dtos.List
	UpdateListAttributes(listId string, card *dtos.List)
	// delete all cards inside it
	DeleteList(id string)
	CheckListExists(id string) bool
}
