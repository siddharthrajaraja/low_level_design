package dtos

import uuid "github.com/satori/go.uuid"

type List struct {
	Id       string
	CardIds  []string
	ListType ListType
	Name     string
}

func NewList(name string, listType ListType) *List {
	uuidValue, _ := uuid.NewV4()

	return &List{
		Id:       uuidValue.String(),
		Name:     name,
		ListType: listType,
	}
}

func (l *List) GetListId() string {

	return l.Id
}

func (l *List) GetListType() ListType {

	return l.ListType
}

func (l *List) GetCardIds() []string {

	return l.CardIds
}

func (l *List) SetCardId(cardId string) {

	l.CardIds = append(l.CardIds, cardId)
}

func (l *List) SetListType(listType ListType) {

	l.ListType = listType
}
