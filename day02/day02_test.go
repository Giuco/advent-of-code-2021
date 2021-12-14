package main

import "testing"

const test_input = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func TestParseInput(t *testing.T) {
	output := parseInput(test_input)
	if len(output) != 6 {
		t.Errorf("Expected Size: %v Actual Size %v", 6, len(output))
	}
}

func TestMoveShip(t *testing.T) {
	input := parseInput(test_input)
	output := moveShip(input, false)
	if output.result() != 150 {
		t.Errorf("Expected result: %v Actual Result %v", 150, output.result())
	}
}

func TestMoveShipAim(t *testing.T) {
	input := parseInput(test_input)
	output := moveShip(input, true)
	if output.result() != 900 {
		t.Errorf("Expected result: %v Actual Result %v", 900, output.result())
	}
}
