package msg

type PositionPlayerMsg struct {
	MainPosition string `json:"position"`
}

type PositionPlayerMsgReply struct {
	MovePosition string `json:"moveposition"`
}
