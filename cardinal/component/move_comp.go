package component

type Move struct {
	Move int
}

func (Move) Name() string {
	return "Move"
}
