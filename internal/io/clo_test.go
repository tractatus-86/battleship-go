package io

import (
	"battleship/internal/game"
	"fmt"
	"os"
	"testing"
)

func TestCLO(t *testing.T) {

	buffer := os.Stdout
	clo := NewPrinter(Simple, buffer)

	clo.UpdateError(fmt.Errorf("test"))
	clo.UpdateInfo("test")
	clo.UpdateGameGrid(game.NewBoard().GetGrid())
	clo.UpdateResult("test")
	clo.UpdateInput("test")

	fmt.Println(clo.info.ReadString('\n'))
	clo.PrintAll()
}
