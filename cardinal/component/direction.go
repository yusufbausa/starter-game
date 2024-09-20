package component

type Direction struct {
	W string
	A string
	S string
	D string
	// Direction string
}

func (Direction) Name() string {
	return "Direction"
}
