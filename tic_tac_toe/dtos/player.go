package dtos

type Player struct {
	Name   string
	choice string
}

func NewPlayer(name string) *Player {

	return &Player{
		Name: name,
	}
}

func (p *Player) AssignChoice(choice string) {

	p.choice = choice
}
func (p *Player) GetChoice() string {

	return p.choice
}
