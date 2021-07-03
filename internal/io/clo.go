package io

import (
	"battleship/internal/game"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type DisplayMode = string

const (
	Simple DisplayMode = "Simple"
	Pretty DisplayMode = "Pretty"
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
	displayMode DisplayMode
	info        *bytes.Buffer
	input       *bytes.Buffer
	page        *bytes.Buffer
	result      *bytes.Buffer
	err         *bytes.Buffer
	outBuffer   *os.File
}

func NewPrinter(displayMode DisplayMode, outBuffer *os.File) *Printer {
	var info bytes.Buffer
	var input bytes.Buffer
	var page bytes.Buffer
	var result bytes.Buffer
	var err bytes.Buffer
	return &Printer{displayMode, &info, &input, &page, &result, &err, outBuffer}
}

func (printer *Printer) PrintInfo() {
	displayMode := printer.displayMode
	writer := printer.outBuffer
	if printer.info.Len() > 0 {
		label := ""
		if displayMode == Pretty {
			label = "\nPhase: "
		}
		writer.WriteString(fmt.Sprintf("%v%v\n", label, printer.info.String()))
	}
}

func (printer *Printer) PrintPage() {
	displayMode := printer.displayMode
	writer := printer.outBuffer
	if displayMode == Pretty {
		writer.WriteString(printer.page.String())
		if printer.input.Len() > 0 {
			writer.WriteString(fmt.Sprintf("Input: %v\n", printer.input.String()))
		}
	}
	printer.page.Reset()
}

func (printer *Printer) PrintResult() {
	displayMode := printer.displayMode
	writer := printer.outBuffer
	if printer.result.Len() > 0 {
		label := ""
		if displayMode == Pretty {
			label = "Response: "
		}
		writer.WriteString(fmt.Sprintf("%v%v\n", label, printer.result.String()))

	}
	printer.result.Reset()
}

func (printer *Printer) PrintError() {
	displayMode := printer.displayMode
	writer := printer.outBuffer
	if displayMode == Pretty && printer.err.Len() > 0 {
		writer.WriteString(fmt.Sprintf("ERROR: %v\n", printer.err.String()))
	}
	printer.err.Reset()
}

func (printer *Printer) PrintInput() {
	displayMode := printer.displayMode
	writer := printer.outBuffer
	if displayMode == Pretty {
		writer.WriteString(fmt.Sprintln("Input Command:"))
	}
}

func (printer *Printer) PrintAll() {

	displayMode := printer.displayMode
	if displayMode == Pretty {
		printer.PrintInfo()
	}
	printer.PrintPage()
	printer.PrintResult()
	printer.PrintError()
	printer.input.Reset()
}

func (printer *Printer) ResetAnimation(secFreq time.Duration) {
	if printer.outBuffer == os.Stdout {
		time.Sleep(secFreq * time.Millisecond)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
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

func (printer *Printer) UpdateInfo(info string) {
	printer.info.Reset()
	printer.info.WriteString(info)
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
