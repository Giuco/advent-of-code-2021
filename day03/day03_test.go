package main

import "testing"

const test_input = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func TestToDecimal(t *testing.T) {
	var i binNum = "10110"
	d := i.toDecimal()

	if d != 22 {
		t.Errorf("%v != %v", d, 22)
	}
}

func TestGetGammaAndEpsilon(t *testing.T) {
	input := parseInput(test_input)
	gamma := input.getGamma()
	epsilon := input.getEpsilon()

	if gamma != 22 {
		t.Errorf("Expected gamma: %v Actual gamma %v", 22, gamma)
	}

	if epsilon != 9 {
		t.Errorf("Expected epsilon: %v Actual epsilon %v", 9, epsilon)
	}
}

func TestGetOxygenCO2(t *testing.T) {
	input := parseInput(test_input)
	oxygen := input.getOxygen()
	co2 := input.getCO2()

	if oxygen != 23 {
		t.Errorf("Expected oxygen: %v Actual oxygen %v", 22, oxygen)
	}

	if co2 != 10 {
		t.Errorf("Expected co2: %v Actual co2 %v", 9, co2)
	}
}
