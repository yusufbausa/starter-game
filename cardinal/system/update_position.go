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

	"strings"

)

// type PlayerDirectionRequest struct {
// 	// W string
// 	// A string
// 	// S string
// 	// D string
// 	InputKeys string
// }

// type PlayerDirectionResponse struct {
// 	Direction int
// }

func PositionSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.PlayerPositionMsg, msg.PlayerPositionMsgReply](
		world,
		func(position cardinal.TxData[msg.PlayerPositionMsg]) (msg.PlayerPositionMsgReply, error) {
			playerID, playerPosition, err := queryTargetPlayerPosition(world, position.Msg.TargetNickname)
			if err != nil {
				return msg.PlayerPositionMsgReply{}, fmt.Errorf("failed to move player: %w", err)
			}
			switch strings.ToUpper(position.Msg.Direction) {
			case "W":
				playerPosition.Y++
			case "A":
				playerPosition.X--
			case "S":
				playerPosition.Y--
			case "D":
				playerPosition.X++
			// default:
			// 	playerPosition
			}
			// playerPosition.Position -= AttackDamage
			// playerHealth.HP -= 
			if err := cardinal.SetComponent[comp.Position](world, playerID, playerPosition); err != nil {
				return msg.PlayerPositionMsgReply{}, fmt.Errorf("failed to update position: %w", err)
			}

			return msg.PlayerPositionMsgReply{Position: *playerPosition}, nil
			// return msg.PlayerPositionMsgReply{}, fmt.Errorf("failed to inflict damage: %w", err)
			// return nil
		})
}

// func DirectionSystem(world cardinal.WorldContext) error {
// 	return cardinal.NewSearch().Entity(
// 		filter.Exact(filter.Component[comp.Player](), filter.Component[comp.Position]())).
// 		Each(world, func(id types.EntityID) bool {
// 			position, err := cardinal.GetComponent[comp.Position](world, id)
// 			if err != nil {
// 				return true
// 			}
// 			if
// 			position.X++
// 			if direction.W == req.Keys {
// 				playerDirection, err = cardinal.SetComponent[comp.Position](world, id)
// 				if err != nil {
// 					return false
// 				}
// 				return false
// 			}


// 			health.HP++
// 			if err := cardinal.SetComponent[comp.Health](world, id, health); err != nil {
// 				return true
// 			}
// 			return true
// 		})
// }



// func DirectionSystem(world cardinal.WorldContext, req *PlayerDirectionRequest) (*PlayerDirectionResponse, error) {
// 	// var playerDirection *comp.Move
// 	var playerPosition *comp.Position
// 	// var keys /**comp.Move*/
// 	var err error
// 	searchErr := cardinal.NewSearch().Entity(
// 		filter.Exact(filter.Component[comp.Player](), /*filter.Component[comp.Move](),*/ filter.Component[comp.Position]())).
// 		Each(world, func(id types.EntityID) bool {
			
// 			// var playerPosition *comp.Position
// 			// keys, err = cardinal.GetComponent[comp.Move](world, id)
// 			// if err != nil {
// 			// 	return false
// 			// }

// 			// Terminates the search if the player is found
// 			// if keys.W == req.InputKeys {
// 			// 	playerPosition.Y++
// 			// } else if keys.A == req.InputKeys {
// 			// 	playerPosition.X--
// 			// } else if keys.S == req.InputKeys {
// 			// 	playerPosition.Y--
// 			// } else if keys.D == req.InputKeys {
// 			// 	playerPosition.X++
// 			// } /*else {
// 			// 	return nil, fmt.Errorf("Keys %s does not exist", req.Direction)
// 			// }*/


// 			if req.InputKeys == "W" {
// 				playerPosition.Y++
// 			} else if req.InputKeys == "A" {
// 				playerPosition.X--
// 			} else if req.InputKeys == "S"{
// 				playerPosition.Y--
// 			} else if req.InputKeys == "D"{
// 				playerPosition.X++
// 			}

// 			// Continue searching if the player is not the target player
// 			return true
// 		})
// 	if searchErr != nil {
// 		return nil, searchErr
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	// if keys == nil {
// 	// 	return nil, fmt.Errorf("player %s does not exist", req.InputKeys)
// 	// }

// 	return &PlayerDirectionResponse{Direction: playerPosition.X}, nil
// }



