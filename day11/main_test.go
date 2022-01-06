package main

import (
	"testing"
)

var rawInput = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

func TestParse(t *testing.T) {
	expectedOutput := [][]int{
		{5, 4, 8, 3, 1, 4, 3, 2, 2, 3},
		{2, 7, 4, 5, 8, 5, 4, 7, 1, 1},
		{5, 2, 6, 4, 5, 5, 6, 1, 7, 3},
		{6, 1, 4, 1, 3, 3, 6, 1, 4, 6},
		{6, 3, 5, 7, 3, 8, 5, 4, 7, 8},
		{4, 1, 6, 7, 5, 2, 4, 6, 4, 5},
		{2, 1, 7, 6, 8, 4, 1, 7, 2, 1},
		{6, 8, 8, 2, 8, 8, 1, 1, 3, 4},
		{4, 8, 4, 6, 8, 4, 8, 5, 5, 4},
		{5, 2, 8, 3, 7, 5, 1, 5, 2, 6},
	}
	output := parseInput(rawInput)

	if len(expectedOutput) != len(output) {
		t.Errorf("Expected n rows %v. Got %v", len(expectedOutput), len(output))
	}

	for i := 0; i < len(expectedOutput); i++ {
		if len(expectedOutput[i]) != len(output[i]) {
			t.Errorf("Expect n cols %v for row %v, got %v", len(expectedOutput[i]), i, len(output[i]))
		}

		for j := 0; j < len(expectedOutput[i]); j++ {
			if expectedOutput[i][j] != output[i][j] {
				t.Errorf("Cell (%v, %v) = %v, expected %v", i, j, output[i][j], expectedOutput[i][j])
			}
		}
	}
}

func TestRunStep(t *testing.T) {
	rawInput := `11111
19991
19191
19991
11111`
	octoMap := parseInput(rawInput)
	nFlashes := runStep(octoMap)
	if nFlashes != 9 {
		t.Errorf("Expected %v. Got %v", 9, nFlashes)
	}
	nFlashes = runStep(octoMap)
	if nFlashes != 0 {
		t.Errorf("Expected %v. Got %v", 0, nFlashes)
	}
}

func TestRunNSteps(t *testing.T) {
	octoMap := parseInput(rawInput)
	output := runNSteps(octoMap, 100)

	if output != 1656 {
		t.Errorf("Expect %v. Got %v", 1656, output)
	}
}
