package dtos

import uuid "github.com/satori/go.uuid"

type List struct {
	id       string
	cardIds  []string
	listType ListType
	name     string
}

func NewList(name string, listType ListType) *List {
	uuidValue, _ := uuid.NewV4()

	return &List{
		id:       uuidValue.String(),
		name:     name,
		listType: listType,
	}
}

func (l *List) GetListId() string {

	return l.id
}

func (l *List) GetListType() ListType {

	return l.listType
}

func (l *List) GetCardIds() []string {

	return l.cardIds
}

func (l *List) SetCardId(cardId string) {

	l.cardIds = append(l.cardIds, cardId)
}

func (l *List) SetListType(listType ListType) {

	l.listType = listType
}
