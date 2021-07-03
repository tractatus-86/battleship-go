package game

import (
	"testing"
)

func TestRules(t *testing.T) {
	gameState := NewGame()
	if gameState.GetGamePhase() != Setup {
		t.Errorf("New Board should be %v phase", Setup)
	}
	if gameState.GetBoard() == nil {
		t.Errorf("Game board shoudl be init")
	}
	err := outOfBoundsCheck(9, 9, Down, 2, gameState.GetBoard().GetGrid())
	if err == nil {
		t.Errorf("outOfBounds should have erred")
	}
	err = outOfBoundsCheck(9, 9, Right, 2, gameState.GetBoard().GetGrid())
	if err == nil {
		t.Errorf("outOfBounds should have erred")
	}
	_, err = gameState.PlaceShip(BATTLESHIP, Down, 0, 0)
	err = shipCollisionDetection(0, 0, Down, ships[BATTLESHIP], *gameState.board)
	if err == nil {
		t.Errorf("shipCollisionDetection should have erred not %v", err)
	}
	gameState.PlaceShip(BATTLESHIP, Down, 0, 1)
	if !shipAlreadyPlaced(*ships[BATTLESHIP]) {
		t.Errorf("This ship was already placed")
	}
	gameState.PlaceShip(EntityName(CARRIER), Down, 0, 1)
	if allShipsPlaced() {
		t.Errorf("All ships haven't been placed")
	}
	gameState.PlaceShip(EntityName(CRUISER), Down, 0, 2)
	if allShipsPlaced() {
		t.Errorf("All ships haven't been placed")
	}
	gameState.PlaceShip(EntityName(DESTROYER), Down, 0, 3)
	if allShipsPlaced() {
		t.Errorf("All ships haven't been placed")
	}
	gameState.PlaceShip(EntityName(SUBMARINE), Down, 0, 4)
	if !allShipsPlaced() {
		t.Errorf("All ships are been placed")
	}
	if gameState.GetGamePhase() != Battle {
		t.Errorf("After placements Board should be %v phase", Battle)
	}

	_, err = gameState.Fire(10, 10)
	if err == nil {
		t.Errorf("fire should have been out of bounds")
	}
	for i := 9; i >= 0; i-- {

		for j := 9; j >= 0; j-- {
			if allShipsSunk() {
				t.Errorf("ships not yet sunk")
			}
			gameState.Fire(i, j)

		}
	}
	if !allShipsSunk() {
		t.Errorf("all ships should be sunk")
	}
	_, err = gameState.Fire(8, 8)
	if err == nil {
		t.Errorf("already fired here")
	}
	if gameState.GetGamePhase() != GameOver {
		t.Errorf("After placements Board should be %v phase", GameOver)
	}

}

func TestFire(t *testing.T) {

}

func TestAllShipsSunk(t *testing.T) {

}
