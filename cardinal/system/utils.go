package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "starter-game/component"
)

// queryTargetPlayer queries for the target player's entity ID and health component.
func queryTargetPlayer(world cardinal.WorldContext, targetNickname string) (types.EntityID, *comp.Health, error) {
	var playerID types.EntityID
	var playerHealth *comp.Health
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Exact(filter.Component[comp.Player](), filter.Component[comp.Health](), filter.Component[comp.Position]())).Each(world,
		func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == targetNickname {
				playerID = id
				playerHealth, err = cardinal.GetComponent[comp.Health](world, id)
				if err != nil {
					return false
				}
				return false
			}

			// Continue searching if the player is not the target player
			return true
		})
	if searchErr != nil {
		return 0, nil, err
	}
	if err != nil {
		return 0, nil, err
	}
	if playerHealth == nil {
		return 0, nil, fmt.Errorf("player %q does not exist", targetNickname)
	}

	return playerID, playerHealth, err
}



//queryTargetPlayerPosition queries for the target player's entity ID and health component.
func queryTargetPlayerPosition(world cardinal.WorldContext, targetNickname string) (types.EntityID, *comp.Position, error) {
	var playerID types.EntityID
	var playerPosition *comp.Position
	// var playerHealth *comp.Health
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Exact(filter.Component[comp.Player](), filter.Component[comp.Health](), filter.Component[comp.Position]())).Each(world,
		func(id types.EntityID) bool {
			var player *comp.Player
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == targetNickname {
				playerID = id
				playerPosition, err = cardinal.GetComponent[comp.Position](world, id)
				// playerHealth, err = cardinal.GetComponent[comp.Health](world, id)
				if err != nil {
					return false
				}
				return false
			}

			// Continue searching if the player is not the target player
			return true
		})
	if searchErr != nil {
		return 0, nil, err
	}
	if err != nil {
		return 0, nil, err
	}
	if playerPosition == nil {
		return 0, nil, fmt.Errorf("player %q does not exist", targetNickname)
	}
	// if playerHealth == nil {
	// 	return 0, nil, fmt.Errorf("player %q does not exist", targetNickname)
	// }
	return playerID, playerPosition, err
}




// queryTargetPlayer queries for the target player's entity ID and health component.
// func queryTargetPlayerPosition(world cardinal.WorldContext, targetNickname string) (types.EntityID, *comp.Position, /**comp.Position,*/ error) {
// 	var playerID types.EntityID
// 	var playerPosition *comp.Position
// 	// var playerPosition *comp.Position  //new line
// 	// var playerDirection *comp.Move
// 	var err error
// 	searchErr := cardinal.NewSearch().Entity(
// 		filter.Exact(filter.Component[comp.Player](), filter.Component[comp.Position]()/*, filter.Component[comp.Position]()*/)).Each(world,
// 		func(id types.EntityID) bool {
// 			// var playerPosition *comp.Position
// 			var player *comp.Player
// 			player, err = cardinal.GetComponent[comp.Player](world, id)
// 			if err != nil {
// 				return false
// 			}

// 			// Terminates the search if the player is found
// 			if player.Nickname == targetNickname {
// 				playerID = id
// 				playerPosition, err = cardinal.GetComponent[comp.Position](world, id)
// 				if err != nil {
// 					return false
// 				}
// 				return false
// 			}

// 			// if player.Nickname == targetNickname {
// 			// 	playerID = id
// 			// 	playerPosition, err = cardinal.GetComponent[comp.Position](world, id)
// 			// 	if err != nil {
// 			// 		return false
// 			// 	}
// 			// 	return false
// 			// }

// 			// Continue searching if the player is not the target player
// 			return true
// 		})
// 	if searchErr != nil {
// 		return 0, nil, err
// 	}
// 	if err != nil {
// 		return 0, nil, err
// 	}
// 	if playerPosition == nil {
// 		return 0, nil, fmt.Errorf("player %q does not exist", targetNickname)
// 	}

// 	return playerID, playerPosition, /*playerPosition,*/ err
// }


// func queryTargetPlayerDirection(world cardinal.WorldContext, targetNickname string) (types.EntityID, *comp.Direction, /**comp.Position,*/ error) {
// 	var playerID types.EntityID
// 	var playerDirection *comp.Direction
// 	// var playerPosition *comp.Position  //new line
// 	// var playerDirection *comp.Move
// 	var err error
// 	searchErr := cardinal.NewSearch().Entity(
// 		filter.Exact(filter.Component[comp.Player](), filter.Component[comp.Direction]()/*, filter.Component[comp.Position]()*/)).Each(world,
// 		func(id types.EntityID) bool {
// 			// var playerPosition *comp.Position
// 			var player *comp.Player
// 			player, err = cardinal.GetComponent[comp.Player](world, id)
// 			if err != nil {
// 				return false
// 			}

// 			// Terminates the search if the player is found
// 			if player.Nickname == targetNickname {
// 				playerID = id
// 				playerDirection, err = cardinal.GetComponent[comp.Direction](world, id)
// 				if err != nil {
// 					return false
// 				}
// 				return false
// 			}

// 			// if player.Nickname == targetNickname {
// 			// 	playerID = id
// 			// 	playerPosition, err = cardinal.GetComponent[comp.Position](world, id)
// 			// 	if err != nil {
// 			// 		return false
// 			// 	}
// 			// 	return false
// 			// }

// 			// Continue searching if the player is not the target player
// 			return true
// 		})
// 	if searchErr != nil {
// 		return 0, nil, err
// 	}
// 	if err != nil {
// 		return 0, nil, err
// 	}
// 	if playerDirection == nil {
// 		return 0, nil, fmt.Errorf("player %q does not exist", targetNickname)
// 	}

// 	return playerID, playerDirection, /*playerPosition,*/ err
// }