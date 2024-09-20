package msg

import (
	comp "starter-game/component"
)

type PlayerPositionMsg struct {
	TargetNickname string `json:"target"`
	Direction string `json:"Enter W/A/S/D to proceed"`
}

type PlayerPositionMsgReply struct {
	Position comp.Position `json:"position"`
}
