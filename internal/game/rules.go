package game

import (
	"battleship/internal/utils"
	"fmt"
)

var GamePhase = Setup

type Phase = string

type Effect = string

const (
	Setup    Phase  = "Setup"
	Battle   Phase  = "Battle"
	GameOver Phase  = "Game Over"
	Hit      Effect = "Hit"
	Miss     Effect = "Miss"
	Sunk     Effect = "You Sunk My"
	Placed   Effect = "Placed"
)

type GameState struct {
	gamePhase Phase
	board     *Board
}

func (gameState *GameState) GetBoard() *Board {
	return gameState.board
}

func (gameState *GameState) GetGamePhase() Phase {
	return gameState.gamePhase
}

func (gameState *GameState) setGamePhase(phase Phase) {
	gameState.gamePhase = phase
}

func NewGame() *GameState {
	board := NewBoard()
	gamePhase := Setup
	return &GameState{gamePhase, board}
}

func (gameState *GameState) Fire(pRow1, pCol1 int) (string, error) {
	board := gameState.board
	err := outOfBoundsCheck(pRow1, pCol1, None, 1, board.GetGrid())
	if err != nil {
		return "", err
	}
	entity := board.Retrieve(pRow1, pCol1).GetContents()
	sRow1 := entity.GetStartPosition().GetRow()
	sCol1 := entity.GetStartPosition().GetColumn()
	result := Miss
	switch entity.GetDirection() {
	case Down:
		if entity.GetIntegrity()[pRow1-sRow1] {
			return "", fmt.Errorf("row: %v column: %v already marked", pRow1, pCol1)
		}
		result = Hit
		entity.GetIntegrity()[pRow1-sRow1] = true
	case Right:
		if entity.GetIntegrity()[pCol1-sCol1] {
			return "", fmt.Errorf("row: %v column: %v already marked", pRow1, pCol1)
		}
		result = Hit
		entity.GetIntegrity()[pCol1-sCol1] = true
	default:
		if entity.GetIntegrity()[0] {
			return "", fmt.Errorf("row: %v column: %v already marked", pRow1, pCol1)
		}
		entity.GetIntegrity()[0] = true
		return result, nil
	}
	if entity.GetState() {
		result = fmt.Sprintf("%v %v!", Sunk, utils.CamelCase(string(entity.GetType())))
	}
	if allShipsSunk() {
		gameState.gamePhase = GameOver
	}
	return result, nil
}

func allShipsSunk() bool {
	all_ships_sunk := true
	for _, v := range ships {
		all_ships_sunk = all_ships_sunk && v.GetState()
	}
	return all_ships_sunk
}

func shipAlreadyPlaced(ship Ship) bool {
	return ship.direction != None
}

func allShipsPlaced() bool {
	all_ships_placed := true
	for _, v := range ships {
		all_ships_placed = all_ships_placed && v.direction != None
	}
	return all_ships_placed
}

func outOfBoundsCheck(row, col int, direction Direction, length int, grid Grid) error {
	var start int
	switch direction {
	case Right:
		start = col
	default:
		start = row
	}
	collisionCheck := start + length
	if collisionCheck > len(grid) {
		return fmt.Errorf("ship out of grid bounds")
	}
	return nil
}

func shipCollisionDetection(row, col int, direction Direction, ship *Ship, board Board) error {
	for i := range ship.integrity {
		var iRow, iCol = row, col
		switch direction {
		case Right:
			iCol = iCol + i
		case Down:
			iRow = iRow + i
		default:
			return fmt.Errorf("bad direction %v", direction)
		}
		shipD := board.Retrieve(iRow, iCol).GetContents().GetType()
		if shipD != WATER {
			return fmt.Errorf("%v collides with ship %v", utils.CamelCase(string(ship.GetType())), utils.CamelCase(string(shipD)))
		}
	}
	return nil
}

func (gameState *GameState) Exit() {
	gameState.gamePhase = GameOver
}

func (gameState *GameState) PlaceShip(ship_name EntityName, direction Direction, row, col int) (string, error) {
	ship := ships[ship_name]
	if shipAlreadyPlaced(*ship) {
		return "", fmt.Errorf("ship %v already placed", utils.CamelCase(string(ship_name)))
	}
	board := gameState.board
	err := outOfBoundsCheck(row, col, direction, len(ship.integrity), board.grid)
	if err != nil {
		return "", err
	}
	err = shipCollisionDetection(row, col, direction, ship, *board)
	if err != nil {
		return "", err
	}
	ship.direction = direction
	board.Place(row, col, ship)
	ship.startPosition = board.Retrieve(row, col)
	if allShipsPlaced() {
		gameState.setGamePhase(Battle)
	}
	return fmt.Sprintf("%v %v", Placed, utils.CamelCase(string(ship_name))), nil
}
