package component

type Position struct {
	PositionPlayer int
}

func (Position) Name() string {
	return "Position"
}
