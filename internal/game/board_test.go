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
	if NewBoard().GetSize() != 10 {
		t.Errorf("grid should not be len %v\n", len(grid))
	}
	pos := board.Retrieve(4, 4)
	if pos.GetContents().GetType() != WATER {
		t.Errorf("%v Not %v entity\n", pos.GetContents().GetType(), WATER)
	}
	if pos.GetContents().GetStartPosition().GetRow() != 4 {
		t.Errorf("Entity position row %v is not position row %v\n", pos.GetContents().GetStartPosition().GetRow(), 4)
	}
	if pos.GetContents().GetStartPosition().GetColumn() != 4 {
		t.Errorf("Entity position col %v is not position col %v\n", pos.GetContents().GetStartPosition().GetColumn(), 4)
	}

	entity := ships[CARRIER]
	entity.direction = Down
	board.Place(pos.GetRow(), pos.GetColumn(), entity)
	pos = board.Retrieve(pos.GetRow(), pos.GetColumn())
	entity.setStartPosition(pos)
	if pos.GetContents() != ships[CARRIER] {
		t.Errorf("%v Not %v entity\n", pos.GetContents().GetType(), ships[CARRIER])
	}
	if ships[CARRIER].GetStartPosition() != pos {
		t.Errorf("Ship start position %v is not %v \n", ships[CARRIER].GetStartPosition(), pos)
	}
	// for i := range ships[CARRIER].GetIntegrity() {
	// 	pos = board.Retrieve(pos.GetRow()+i, pos.GetColumn())
	// 	if pos.GetContents() != ships[CARRIER] {
	// 		t.Errorf("%v Not %v entity at pos %v\n", pos.GetContents().GetType(), ships[CARRIER].GetType(), pos)
	// 	}
	// }
}
