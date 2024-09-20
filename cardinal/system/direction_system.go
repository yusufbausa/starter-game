package system

import (
	"fmt"

	// "pkg.world.dev/world-engine/cardinal/search/filter"
	// "pkg.world.dev/world-engine/cardinal/types"

	"pkg.world.dev/world-engine/cardinal"
	// "pkg.world.dev/world-engine/cardinal/search/filter"
	// "pkg.world.dev/world-engine/cardinal/types"

	comp "starter-game/component"
	"starter-game/msg"

	// "strings"
)

type PlayerDirectionRequest struct {
	// W string
	// A string
	// S string
	// D string
	InputKeys string
}

type PlayerDirectionResponse struct {
	Direction int
}



// func DirectionSystem(world cardinal.WorldContext, req *PlayerDirectionRequest) (*PlayerDirectionResponse, error) {
// 	// var playerDirection *comp.Move
// 	var playerPosition *comp.Position
// 	var keys *comp.Move
// 	var err error
// 	searchErr := cardinal.NewSearch().Entity(
// 		filter.Exact(filter.Component[comp.Player](), filter.Component[comp.Direction](), filter.Component[comp.Position]())).
// 		Each(world, func(id types.EntityID) bool {
			
// 			// var playerPosition *comp.Position
// 			// keys, err = cardinal.GetComponent[comp.Move](world, id)
// 			// if err != nil {
// 			// 	return false
// 			// }

// 			// Terminates the search if the player is found
// 			if keys.W == req.InputKeys {
// 				playerPosition.Y++
// 			} else if keys.A == req.InputKeys {
// 				playerPosition.X--
// 			} else if keys.S == req.InputKeys {
// 				playerPosition.Y--
// 			} else if keys.D == req.InputKeys {
// 				playerPosition.X++
// 			} /*else {
// 				return nil, fmt.Errorf("Keys %s does not exist", req.Direction)
// 			}*/


// 			// if req.InputKeys == "W" {
// 			// 	playerPosition.Y++
// 			// } else if req.InputKeys == "A" {
// 			// 	playerPosition.X--
// 			// } else if req.InputKeys == "S"{
// 			// 	playerPosition.Y--
// 			// } else if req.InputKeys == "D"{
// 			// 	playerPosition.X++
// 			// }

// 			// Continue searching if the player is not the target player
// 			return true
// 		})
// 	if searchErr != nil {
// 		return nil, searchErr
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	if keys == nil {
// 		return nil, fmt.Errorf("player %s does not exist", req.InputKeys)
// 	}

// 	return &PlayerDirectionResponse{Direction: playerPosition.X}, nil
// }

func DirectionSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.PlayerDirectionMsg, msg.PlayerDirectionMsgReply](
		world,
		func(direction cardinal.TxData[msg.PlayerDirectionMsg]) (msg.PlayerDirectionMsgReply, error) {
			// playerID, playerDirection, err := queryTargetPlayerDirection(world, direction.Msg.TargetNickname)
			// if err != nil {
			// 	return msg.PlayerDirectionMsgReply{}, fmt.Errorf("failed to found : %w", err)
			// }
			
			playerID, playerDirection, err := queryTargetPlayerDirection(world, direction.Msg.TargetNickname)
			if err != nil {
				return msg.PlayerDirectionMsgReply{}, fmt.Errorf("failed to found user : %w", err)
			}

			if err := cardinal.SetComponent[comp.Direction](world, playerID, playerDirection); err != nil {
				return msg.PlayerDirectionMsgReply{}, fmt.Errorf("failed to update position: %w", err)
			}

			// return msg.PlayerPositionMsgReply{Position: AttackDamage}, nil
			return msg.PlayerDirectionMsgReply{}, fmt.Errorf("failed to inflict damage: %w", err)
			// return nil
		})

	}

