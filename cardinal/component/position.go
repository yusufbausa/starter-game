package component

// type Position struct {
// 	PositionPlayer int
// }

type Position struct {
	X int
	Y int
}

func (Position) Name() string {
	return "Position"
}
