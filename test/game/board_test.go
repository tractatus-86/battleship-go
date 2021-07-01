package game

import (
	"battleship/internal/game"
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := game.NewBoard()

	var grid game.Grid
	grid = board.GetGrid()
	if grid == nil {
		t.Errorf("grid %v should not be \n", grid)
	}
	if len(grid) != 10 {
		t.Errorf("grid should not be len %v\n", len(grid))
	}

}

func TestBoardPlace(t *testing.T) {

}

func TestBoardRetrieve(t *testing.T) {

}

func TestBoardGetGrid(t *testing.T) {

}

func TestBoardGetSize(t *testing.T) {

}

func TestPositionGetRow(t *testing.T) {

}

func TestPositionGetColumn(t *testing.T) {

}

func TestPositionGetContents(t *testing.T) {

}
