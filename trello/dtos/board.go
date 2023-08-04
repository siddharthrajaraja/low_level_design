package dtos

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/siddharthrajaraja/low_level_design/trello/utils"
)

type Board struct {
	id      string
	name    string
	privacy BoardPrivacy
	url     string
	members []string
	lists   []string
}

func NewBoard(name string) *Board {

	uuidValue, _ := uuid.NewV4()

	return &Board{
		id:      uuidValue.String(),
		name:    name,
		privacy: Public,
		url:     fmt.Sprintf(utils.BoardUrlPrefix, uuidValue),
	}
}

func (b *Board) GetBoardId() string {

	return b.id
}

func (b *Board) GetMembersList() []string {

	return b.members
}

func (b *Board) SetMember(member string) {

	b.members = append(b.members, member)
}

func (b *Board) SetMembersList(member []string) {

	b.members = member
}

func (b *Board) GetUrl() string {

	return b.url
}

func (b *Board) SetUrl(url string) {

	b.url = url
}

func (b *Board) GetPrivacy() BoardPrivacy {

	return b.privacy
}

func (b *Board) SetPrivacy(privacy BoardPrivacy) {

	b.privacy = privacy
}

func (b *Board) GetList() []string {

	return b.lists
}

func (b *Board) SetList(listsId string) {

	b.lists = append(b.lists, listsId)
}
