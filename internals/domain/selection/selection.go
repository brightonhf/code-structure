package selection

type selection struct {
}

type Core interface {
	CountSelection(meals int) int
}

func NewCore() Core {
	return &selection{}
}

func (s selection) CountSelection(meals int) int {
	return meals * 2
}
