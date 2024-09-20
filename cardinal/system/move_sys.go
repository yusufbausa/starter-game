package system

// import (
// 	// "bufio"
// 	"fmt"
// 	// "os"
// 	"strings"

// 	"pkg.world.dev/world-engine/cardinal"
// 	"pkg.world.dev/world-engine/cardinal/search/filter"
// 	"pkg.world.dev/world-engine/cardinal/types"

// 	comp "starter-game/component"
// )

// type Character struct {
// 	x, y int
// }

// // RegenSystem replenishes the player's HP at every tick.
// // This provides an example of a system that doesn't rely on a transaction to update a component.
// func MoveSystem(world cardinal.WorldContext) error {
// 	return cardinal.NewSearch().Entity(
// 		filter.Exact(filter.Component[comp.Move](), filter.Component[comp.Move]())).
// 		Each(world, func(id types.EntityID) bool {
// 			move, err := cardinal.GetComponent[comp.Move](world, id)
// 			if err != nil {
// 				return true
// 			}
// 			// move.HP+1
// 			if err := cardinal.SetComponent[comp.Move](world, id, move); err != nil {
// 				return true
// 			}
// 			return true
// 		})
// }

// func (c *Character) Move(input string) {
// 	switch strings.ToUpper(input) {
// 	case "W":
// 		c.y++
// 	case "A":
// 		c.x--
// 	case "S":
// 		c.y--
// 	case "D":
// 		c.x++
// 	default:
// 		fmt.Println("Invalid input. Use W, A, A, or D.")
// 	}
// }

// func MoveDirection() {
// 	char := Character{x: 0, y: 0}
// 	// reader := bufio.NewReader(os.Stdin)

// 	fmt.Println("Use WASD keys to move the character. Type 'exit' to quit.")
// 	for {
// 		fmt.Printf("Current position: (%d, %d)\n", char.x, char.y)
// 		fmt.Print("Move: ")
// 		// input, _ := reader.ReaderString('\n')
// 		// input = strings.TrimSpace(input)

// 		// if input == "exit" {
// 		// 	break
// 		// }
// 		// char.Move(input)

// 	}

// }
