package dtos

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/siddharthrajaraja/low_level_design/trello/utils"
)

type Board struct {
	id      string
	Name    string
	Privacy BoardPrivacy
	Url     string
	Members []string
	Lists   []string
}

func NewBoard(name string) *Board {

	uuidValue, _ := uuid.NewV4()

	return &Board{
		id:      uuidValue.String(),
		Name:    name,
		Privacy: Public,
		Url:     fmt.Sprintf(utils.BoardUrlPrefix, uuidValue),
	}
}

func (b *Board) GetBoardId() string {

	return b.id
}

func (b *Board) GetMembersList() []string {

	return b.Members
}

func (b *Board) SetMember(member string) {

	b.Members = append(b.Members, member)
}

func (b *Board) SetMembersList(member []string) {

	b.Members = member
}

func (b *Board) GetUrl() string {

	return b.Url
}

func (b *Board) SetUrl(url string) {

	b.Url = url
}

func (b *Board) GetPrivacy() BoardPrivacy {

	return b.Privacy
}

func (b *Board) SetPrivacy(privacy BoardPrivacy) {

	b.Privacy = privacy
}

func (b *Board) GetList() []string {

	return b.Lists
}

func (b *Board) SetList(listsId string) {

	b.Lists = append(b.Lists, listsId)
}
