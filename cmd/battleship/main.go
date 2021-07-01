package main

import (
	"battleship/internal/game"
	"battleship/internal/io"
	"fmt"
)

func main() {

	gameState := game.NewGame()
	interpreter := io.NewCommandInterpreter()
	printer := io.NewPrinter()
	setup(interpreter, gameState, printer)
	battle(interpreter, gameState, printer)
	gameover(interpreter, gameState, printer)

}
func gameover(interpreter *io.CommandInterpreter, gameState *game.GameState, printer *io.Printer) {
	fmt.Println(gameState.GetGamePhase())
}
func setup(interpreter *io.CommandInterpreter, gameState *game.GameState, printer *io.Printer) {
	printer.UpdateGameGrid(gameState.GetBoard().GetGrid())
	printer.PrintDisplay()
	for gameState.GetGamePhase() == game.Setup && interpreter.Next() {

		input := interpreter.Content()
		printer.UpdateInput(input)
		_, commandParams, err := io.ParseSetupInput(input)

		if err != nil {
			printer.UpdateError(err)
		}
		switch v := commandParams.(type) {
		case *io.PlaceShipParams:
			effect, err := gameState.PlaceShip(v.ShipName, v.Direction, v.Row, v.Col)
			if err != nil {
				printer.UpdateError(err)

			} else {
				printer.UpdateResult(effect)
			}

		case *io.ExitParams:
			gameState.Exit()
		}
		printer.UpdateGameGrid(gameState.GetBoard().GetGrid())
		printer.PrintDisplay()
	}

}
func battle(interpreter *io.CommandInterpreter, gameState *game.GameState, printer *io.Printer) {
	printer.UpdateGameGrid(gameState.GetBoard().GetGrid())
	printer.PrintDisplay()
	for gameState.GetGamePhase() == game.Battle && interpreter.Next() {

		input := interpreter.Content()
		printer.UpdateInput(input)
		_, CommandParams, err := io.ParseBattleInput(input)
		if err != nil {
			printer.UpdateError(err)

		}

		switch v := CommandParams.(type) {
		case *io.FireParams:
			effect, err := gameState.Fire(v.Row, v.Col)
			if err != nil {
				printer.UpdateError(err)

			} else {
				printer.UpdateResult(effect)
			}
		case *io.ExitParams:
			gameState.Exit()
		}
		printer.UpdateGameGrid(gameState.GetBoard().GetGrid())
		printer.PrintDisplay()

	}

}
