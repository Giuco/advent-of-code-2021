package main

import "testing"

var input = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}

func TestCountIncreases(t *testing.T) {
	output := countIncreases(input)
	if output != 7 {
		t.Errorf("Wrong output. Expected %v. Got %v", 7, output)
	}
}

func TestCountIncreasesWindow(t *testing.T) {
	output := countIncreasesWindow(input, 3)
	if output != 5 {
		t.Errorf("Wrong output. Expected %v. Got %v", 5, output)
	}
}
