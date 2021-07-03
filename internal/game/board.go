package game

type Position struct {
	row    int
	col    int
	entity *Entity
}

func newPosition(row, col int, entity *Entity) *Position {
	return &Position{row, col, entity}
}

func (position *Position) GetRow() int {
	return position.row
}

func (position *Position) GetColumn() int {
	return position.col
}

func (position *Position) GetContents() Entity {
	return *position.entity
}

type Direction = string

type Grid = [][]*Position

type Board struct {
	grid Grid
}

func (board *Board) GetGrid() Grid {
	return board.grid
}

func (board *Board) GetSize() int {
	return len(board.grid)
}

const (
	Down  Direction = "DOWN"
	Right Direction = "RIGHT"
	None  Direction = "NONE"
)

func NewBoard() *Board {
	size := 10
	var grid = make([][]*Position, size)
	board := &Board{grid}
	for j := range grid {
		row := make([]*Position, size)
		grid[j] = row
		for i := range row {
			board.Place(j, i, NewWater())
			board.Retrieve(j, i).GetContents().setStartPosition(board.grid[j][i])
		}
	}
	return board
}

func (board *Board) Place(row, column int, entity Entity) {
	for i := range entity.GetIntegrity() {
		switch entity.GetDirection() {
		case Down:
			board.grid[row+i][column] = newPosition(row+i, column, &entity)
		default:
			board.grid[row][column+i] = newPosition(row, column+i, &entity)
		}
	}
}

func (board *Board) Retrieve(row, column int) *Position {
	return board.grid[row][column]
}
