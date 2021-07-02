package game

type Entity interface {
	GetType() EntityName
	GetIntegrity() []bool
	GetDirection() Direction
	GetStartPosition() *Position
	GetState() bool
	setDirection(direction Direction)
	setStartPosition(position *Position)
}

type EntityName string

const (
	WATER      EntityName = "WATER"
	CARRIER    EntityName = "CARRIER"
	BATTLESHIP EntityName = "BATTLESHIP"
	CRUISER    EntityName = "CRUISER"
	SUBMARINE  EntityName = "SUBMARINE"
	DESTROYER  EntityName = "DESTROYER"
)

type Water struct {
	name          EntityName
	integrity     []bool
	direction     Direction
	startPosition *Position
}

func NewWater() *Water {
	return &Water{WATER, make([]bool, 1), None, nil}
}

func (water *Water) GetType() EntityName {
	return water.name
}

func (water *Water) GetIntegrity() []bool {
	return water.integrity
}

func (water *Water) GetDirection() Direction {
	return water.direction
}

func (water *Water) GetStartPosition() *Position {
	return water.startPosition
}

func (water *Water) GetState() bool {
	return water.integrity[0]
}

func (water *Water) setDirection(direction Direction) {
	water.direction = None
}

func (water *Water) setStartPosition(position *Position) {
	water.startPosition = position
}

type Ship struct {
	class         EntityName
	integrity     []bool
	direction     Direction
	startPosition *Position
}

func (ship *Ship) GetType() EntityName {
	return ship.class
}

func (ship *Ship) GetIntegrity() []bool {
	return ship.integrity
}

func (ship *Ship) GetDirection() Direction {
	return ship.direction
}

func (ship *Ship) GetStartPosition() *Position {
	return ship.startPosition
}

func (ship *Ship) GetState() bool {
	sunk := true
	for _, b := range ship.integrity {
		sunk = sunk && b
	}
	return sunk
}

func (ship *Ship) setDirection(direction Direction) {
	ship.direction = direction
}

func (ship *Ship) setStartPosition(position *Position) {
	ship.startPosition = position
}

var ships = map[EntityName]*Ship{
	CARRIER:    {CARRIER, make([]bool, 5), None, nil},
	BATTLESHIP: {BATTLESHIP, make([]bool, 4), None, nil},
	CRUISER:    {CRUISER, make([]bool, 3), None, nil},
	SUBMARINE:  {SUBMARINE, make([]bool, 3), None, nil},
	DESTROYER:  {DESTROYER, make([]bool, 2), None, nil},
}

var ShipNames = map[string]EntityName{
	"CARRIER":    CARRIER,
	"BATTLESHIP": BATTLESHIP,
	"CRUISER":    CRUISER,
	"SUBMARINE":  SUBMARINE,
	"DESTROYER":  DESTROYER,
}
