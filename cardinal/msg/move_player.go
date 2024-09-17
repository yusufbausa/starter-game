package msg

type MovePlayerDirectionMsg struct {
	TargetDirection string `json:"direction"`
}

type MovePlayerDirectionMsgReply struct {
	Direction string `json:"direction"`
}
