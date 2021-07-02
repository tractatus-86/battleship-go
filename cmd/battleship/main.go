package main

import (
	"battleship/internal/battleship"
	"battleship/internal/game"
	"battleship/internal/io"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	inBuffer, outBuffer, display := parseArgs()
	gameState := game.NewGame()
	interpreter := io.NewCommandInterpreter(inBuffer)
	printer := io.NewPrinter(display, outBuffer)
	defer inBuffer.Close()
	defer outBuffer.Close()
	battleship.Start(interpreter, gameState, printer)
}

func parseArgs() (*os.File, *os.File, io.DisplayMode) {
	inputFile := flag.String("input_file", "", "input file of commands")
	outputFile := flag.String("output_file", "", "output file of responses")
	displayMode := flag.Bool("pretty", false, "parse printer mode")
	flag.Parse()

	inBuffer := os.Stdin
	var err error
	if strings.TrimSpace(*inputFile) != "" {
		inBuffer, err = os.OpenFile(*inputFile, os.O_RDONLY, 0644)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}

	outBuffer := os.Stdout
	if strings.TrimSpace(*outputFile) != "" {
		outBuffer, err = os.OpenFile(*outputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}

	display := io.Simple
	if *displayMode {
		display = io.Pretty
	}
	return inBuffer, outBuffer, display
}
