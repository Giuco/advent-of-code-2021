package main

import "testing"

func TestPart1(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	fuel := calculateFuel(input)
	if fuel != 37 {
		t.Errorf("Got %v. Expected %v", fuel, 37)
	}
}

func TestPart2(t *testing.T) {
	input := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	fuel := calculateFuel2(input)
	if fuel != 168 {
		t.Errorf("Got %v. Expected %v", fuel, 168)
	}
}
