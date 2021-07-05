package io

import (
	"os"
	"testing"
)

func TestCLI(t *testing.T) {

	buffer := os.Stdin
	cli := NewCommandInterpreter(buffer)
	if !cli.IsStdIn() {
		t.Errorf("cli should be stdin")
	}

	col, err := toGridColumnIndex(1)
	if err != nil {
		t.Errorf("%v", err)
	}
	if col != 0 {
		t.Errorf("Columns should be %v but was %v", 0, col)
	}
	row, err := toGridRowIndex("A")
	if err != nil {
		t.Errorf("%v", err)
	}
	if row != 0 {
		t.Errorf("Columns should be %v but was %v", 0, row)
	}

}
