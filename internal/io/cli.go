package io

import (
	"battleship/internal/game"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CommandName = string

type CommandParams interface {
	NumParams() int
}

type FireParams struct {
	Row int
	Col int
}

func (fireParams *FireParams) NumParams() int {
	return 2
}

type ExitParams struct {
}

func (exitParams *ExitParams) NumParams() int {
	return 0
}

type PlaceShipParams struct {
	ShipName  game.EntityName
	Direction game.Direction
	Row       int
	Col       int
}

func (placeShipParams *PlaceShipParams) NumParams() int {
	return 4
}

const (
	Fire      CommandName = "FIRE"
	PlaceShip CommandName = "PLACE_SHIP"
	Exit      CommandName = "EXIT"
)

type CommandInterpreter struct {
	scanner *bufio.Scanner
}

func NewCommandInterpreter() *CommandInterpreter {
	return &CommandInterpreter{bufio.NewScanner(os.Stdin)}
}

func (commandInterpreter *CommandInterpreter) Next() bool {
	return commandInterpreter.scanner.Scan()
}

func (commandInterpreter *CommandInterpreter) Content() string {
	return commandInterpreter.scanner.Text()
}

func ParseBattleInput(input string) (CommandName, CommandParams, error) {
	commandInstruction := strings.Split(strings.ToUpper(input), " ")
	switch len(commandInstruction) {
	case 2:
		return validateCommandFire(commandInstruction)
	case 1:
		return validateCommandExit(commandInstruction)
	default:
		return validateCommandUnknown(commandInstruction)
	}

}

func ParseSetupInput(input string) (CommandName, CommandParams, error) {
	commandInstruction := strings.Split(strings.ToUpper(input), " ")
	switch len(commandInstruction) {
	case 4:
		return validateCommandPlaceShip(commandInstruction)
	case 1:
		return validateCommandExit(commandInstruction)
	default:
		return validateCommandUnknown(commandInstruction)
	}
}
func validateCommandPlaceShip(commandInstruction []string) (CommandName, CommandParams, error) {
	if commandInstruction[0] != PlaceShip {
		return "", nil, fmt.Errorf("unknown commmand %v", commandInstruction[0])
	}

	ok := true
	var direction game.Direction
	shipName, ok := game.ShipNames[commandInstruction[1]]
	if !ok {
		keys := make([]string, 0, len(game.ShipNames))
		for key := range game.ShipNames {
			keys = append(keys, key)
		}
		return "", nil, fmt.Errorf("no such ship %v choose %v", commandInstruction[1], keys)
	}
	switch c := commandInstruction[2]; {
	case c == game.Right:
		direction = game.Right
	case c == game.Down:
		direction = game.Down
	default:
		return "", nil, fmt.Errorf("no such directions %v choose: %v ", c, []string{game.Down, game.Right})
	}

	row, err := toGridRowIndex(string(commandInstruction[3][0]))
	if err != nil {
		return "", nil, err
	}
	i, err := strconv.Atoi(string(commandInstruction[3][1:]))
	if err != nil {
		return "", nil, err
	}
	column, err := toGridColumnIndex(i)
	if err != nil {
		return "", nil, err
	}

	return PlaceShip, &PlaceShipParams{shipName, direction, row, column}, nil
}

func toGridRowIndex(c string) (int, error) {
	size := len(c)
	if size < 1 {
		return -1, fmt.Errorf("no row submitted")
	}
	row := int(rune(c[0]))
	if row < 65 || row > 74 {
		return -1, fmt.Errorf("row %v is out of bounds", c)
	}
	return row - 65, nil
}

func toGridColumnIndex(i int) (int, error) {
	if i > 0 && i <= 10 {
		return i - 1, nil
	} else {
		return -1, fmt.Errorf("%v is not a valid column", i)
	}
}
func validateCommandFire(commandInstruction []string) (CommandName, CommandParams, error) {
	if commandInstruction[0] != Fire {
		return "", nil, fmt.Errorf("unknown commmand %v", commandInstruction[0])
	}
	row, err := toGridRowIndex(string(commandInstruction[1][0]))
	if err != nil {
		return "", nil, err
	}
	i, err := strconv.Atoi(string(commandInstruction[1][1:]))
	if err != nil {
		return "", nil, err
	}
	column, err := toGridColumnIndex(i)
	if err != nil {
		return "", nil, err
	}
	return Fire, &FireParams{row, column}, nil
}
func validateCommandExit(commandInstruction []string) (CommandName, CommandParams, error) {
	if commandInstruction[0] != Exit {
		return "", nil, fmt.Errorf("unknown commmand %v", commandInstruction[0])

	}
	return Exit, &ExitParams{}, nil
}
func validateCommandUnknown(commandInstruction []string) (CommandName, CommandParams, error) {
	return "", nil, fmt.Errorf("unknown commmand %v", commandInstruction[0])
}
