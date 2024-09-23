package query

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal/search/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "starter-game/component"

	"pkg.world.dev/world-engine/cardinal"
)

type PlayerPositionRequest struct {
	Nickname string
}

type PlayerPositionResponse struct {
	Position int
}

func PlayerPosition(world cardinal.WorldContext, req *PlayerPositionRequest) (*PlayerPositionResponse, error) {
	var playerPosition *comp.Position
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Exact(filter.Component[comp.Player](), filter.Component[comp.Position](), filter.Component[comp.Health]())).
		Each(world, func(id types.EntityID) bool {
			var player *comp.Player
			// var playerPosition *comp.Position
			player, err = cardinal.GetComponent[comp.Player](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the player is found
			if player.Nickname == req.Nickname {
				playerPosition, err = cardinal.GetComponent[comp.Position](world, id)
				if err != nil {
					return false
				}
				return false
			}

			// Continue searching if the player is not the target player
			return true
		})
	if searchErr != nil {
		return nil, searchErr
	}
	if err != nil {
		return nil, err
	}

	if playerPosition == nil {
		return nil, fmt.Errorf("player %s does not exist", req.Nickname)
	}

	return &PlayerPositionResponse{Position: playerPosition.X}, nil
	// return &PlayerPositionResponse{Position: playerPosition.Y}, nil
}
