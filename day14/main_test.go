package main

import (
	"testing"
)

var rawInput = `NNCB
CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

func TestPart1(t *testing.T) {
	initial, instructions := parseInput(rawInput)
	expectedOutput := 1588
	actualOutput := partN(initial, instructions, 10)

	if actualOutput != expectedOutput {
		t.Errorf("Wrong! Expected %v, got %v", expectedOutput, actualOutput)
	}
}

func TestPart2(t *testing.T) {
	initial, instructions := parseInput(rawInput)
	expectedOutput := 2188189693529
	actualOutput := partN(initial, instructions, 40)

	if actualOutput != expectedOutput {
		t.Errorf("Wrong! Expected %v, got %v", expectedOutput, actualOutput)
	}
}
