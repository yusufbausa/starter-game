package msg

type PlayerDirectionMsg struct {
	TargetNickname string `json:"target"`
	Direction string `json:"Enter W/A/S/D to proceed"`
}

type PlayerDirectionMsgReply struct {
	DirectionReply string `json:"Direction"`
}
