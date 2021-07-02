package io

import (
	"battleship/internal/game"
	"bytes"
	"fmt"
)

const (
	Miss      Mark = "o"
	Hit       Mark = "x"
	Water     Mark = "~"
	Something Mark = "#"
	Roof      Mark = "_"
	Wall      Mark = "|"
	Delimiter Mark = " "
)

type Mark = string

func toColumn(i int) int {
	return i + 1
}

func toRow(i int) string {
	return string(rune(i + 65))
}

type Printer struct {
	input  *bytes.Buffer
	page   *bytes.Buffer
	result *bytes.Buffer
	err    *bytes.Buffer
}

func NewPrinter() *Printer {
	var input bytes.Buffer
	var page bytes.Buffer
	var result bytes.Buffer
	var err bytes.Buffer
	return &Printer{&input, &page, &result, &err}
}

func (printer *Printer) PrintDisplay() {
	fmt.Println(printer.page.String())
	if printer.input.Len() > 0 {
		fmt.Printf("Input: %v\n", printer.input.String())
	}
	if printer.result.Len() > 0 {
		fmt.Printf("Response: %v\n", printer.result.String())
	}
	if printer.err.Len() > 0 {
		fmt.Printf("ERROR: %v\n", printer.err.String())
	}
	fmt.Println("Input Command:")
	printer.page.Reset()
	printer.result.Reset()
	printer.err.Reset()
	printer.input.Reset()
}

func writeColumnLabels(page *bytes.Buffer, gridLen int) {
	page.WriteString(Delimiter)
	page.WriteString(Delimiter)
	page.WriteString(Delimiter)
	page.WriteString(Delimiter)
	for i := range make([]int, gridLen) {
		page.WriteString(fmt.Sprint(toColumn(i)))
		page.WriteString(Delimiter)
		page.WriteString(Delimiter)
		page.WriteString(Delimiter)
	}
	page.WriteString("\n")
}

func writeRowLabel(i int, page *bytes.Buffer) {
	page.WriteString(toRow(i))
	page.WriteString(Delimiter)
}

func (printer *Printer) UpdateInput(input string) {
	printer.input.Reset()
	printer.input.WriteString(input)
}

func (printer *Printer) UpdateResult(result string) {
	printer.result.Reset()
	printer.result.WriteString(result)
}

func (printer *Printer) UpdateError(err error) {
	printer.err.Reset()
	if err != nil {
		printer.err.WriteString(fmt.Sprintf("%v", err))
	}
}

func writeCellLeftBoarder(page *bytes.Buffer) {
	page.WriteString(Wall)
	page.WriteString(Delimiter)
}

func writeCellRightBoarder(page *bytes.Buffer) {
	page.WriteString(Delimiter)
}

func writeRowTerminator(page *bytes.Buffer) {
	page.WriteString("\n")
}

func writeCell(pos *game.Position, page *bytes.Buffer) {
	contents := pos.GetContents()
	if contents.GetType() == game.WATER {
		if !contents.GetState() {
			page.WriteString(Water)
		} else {
			page.WriteString(Miss)
		}
	} else {
		pRow := pos.GetRow()
		sRow := contents.GetStartPosition().GetRow()
		pCol := pos.GetColumn()
		sCol := contents.GetStartPosition().GetColumn()
		if contents.GetDirection() == game.Down && contents.GetIntegrity()[pRow-sRow] {
			page.WriteString(Hit)
		} else if contents.GetDirection() == game.Right && contents.GetIntegrity()[pCol-sCol] {
			page.WriteString(Hit)
		} else {
			page.WriteString(Something)
		}

	}
}

func (printer *Printer) UpdateGameGrid(grid game.Grid) {
	page := printer.page
	page.Reset()
	writeColumnLabels(page, len(grid))
	for i, j := range grid {
		writeRowLabel(i, page)
		for _, pos := range j {
			writeCellLeftBoarder(page)
			writeCell(pos, page)
			writeCellRightBoarder(page)
		}
		writeRowTerminator(page)
	}

}
