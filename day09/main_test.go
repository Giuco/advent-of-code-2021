package main

import "testing"

const raw_input = `2199943210
3987894921
9856789892
8767896789
9899965678`

func TestParseInput(t *testing.T) {
	parsed := parseInput(raw_input)
	expected := smokeMap{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
	}

	if len(parsed) != len(expected) {
		t.Errorf("Wrong number of rows")
	}

	for i := 0; i < len(parsed); i++ {
		if len(parsed[i]) != len(expected[i]) {
			t.Errorf("Wrong number of columns for row %v", i)
		}

		for j := 0; j < len(parsed[i]); j++ {
			if parsed[i][j] != expected[i][j] {
				t.Errorf("Wrong number on cell (%v,%v). Expected %v, Got %v", i, j, expected[i][j], parsed[i][j])
			}
		}
	}

}

func TestSumLowPoints(t *testing.T) {
	parsedInput := parseInput(raw_input)
	total := sumLowPoints(parsedInput)
	if total != 15 {
		t.Errorf("Expected %v. Got %v", 15, total)
	}
}

func TestPart2(t *testing.T) {
	parsedInput := parseInput(raw_input)
	result := getPart2(parsedInput)
	if result != 1134 {
		t.Errorf("Expected %v. Got %v", 1134, result)
	}
}
