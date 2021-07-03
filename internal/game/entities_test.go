package game

import (
	"testing"
)

func TestEntities(t *testing.T) {

	water := NewWater()

	if water.GetDirection() != None {
		t.Errorf("water should not have direction %v", water.direction)
	}
	if water.GetStartPosition() != nil {
		t.Errorf("water shoudl not have start position %v", water.startPosition)
	}
	if water.GetType() != WATER {
		t.Errorf("water should be named %v not %v", water.GetType(), WATER)
	}
	if len(water.GetIntegrity()) != 1 {
		t.Errorf("water should not have integristy %v", len(water.GetIntegrity()))
	}
	if water.GetState() {
		t.Errorf("water should not start with state %v", water.GetState())
	}
	var ent Entity = water
	water.setStartPosition(newPosition(3, 6, &ent))
	if water.GetStartPosition().GetRow() != 3 {
		t.Errorf("water  rowstart position %v not %v", water.startPosition, 3)
	}
	if water.GetStartPosition().GetColumn() != 6 {
		t.Errorf("water col start position %v not %v", water.startPosition, 6)
	}
	if water.GetStartPosition().GetContents() != ent {
		t.Errorf("Start position content is %v not %v", water.GetStartPosition().GetContents(), ent)
	}
	water.GetIntegrity()[0] = true
	if !water.GetState() {
		t.Errorf("water should not have state %v", water.GetState())
	}
	for shipName := range ShipNames {
		ship := ships[EntityName(shipName)]
		if ship.GetDirection() != None {
			t.Errorf("%v should not have direction %v", shipName, ship.direction)
		}
		if ship.GetStartPosition() != nil {
			t.Errorf("%v shoudl not have start position %v", shipName, ship.startPosition)
		}
		if ship.GetType() != EntityName(shipName) {
			t.Errorf("%v should be named %v not %v", shipName, ship.GetType(), shipName)
		}
		if ship.GetState() {
			t.Errorf("%v should not start with state %v", shipName, ship.GetState())
		}
		var ent Entity = ship
		ship.setStartPosition(newPosition(3, 6, &ent))
		if ship.GetStartPosition().GetRow() != 3 {
			t.Errorf("ship %v rowstart position %v not %v", shipName, ship.startPosition, 3)
		}
		if ship.GetStartPosition().GetColumn() != 6 {
			t.Errorf("ship %v col start position %v not %v", shipName, ship.startPosition, 6)
		}
		if ship.GetStartPosition().GetContents() != ent {
			t.Errorf("Start position content is %v not %v", ship.GetStartPosition().GetContents(), ent)
		}
	}

}
