package game

import (
	"testing"
)

func TestBoard(t *testing.T) {
	board := NewBoard()

	grid := board.GetGrid()
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
