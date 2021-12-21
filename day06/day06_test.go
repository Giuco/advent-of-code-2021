package main

import "testing"

func TestPart1(t *testing.T) {
	la := []int{3, 4, 3, 1, 2}
	count := runDays(la, 18)
	if count != 26 {
		t.Errorf("Got %v. Expected %v", count, 26)
	}

	count = runDays(la, 80)
	if count != 5934 {
		t.Errorf("Got %v. Expected %v", count, 5934)
	}
}

func TestPart2(t *testing.T) {
	la := []int{3, 4, 3, 1, 2}
	count := runDaysOptimized(la, 256)
	if count != 26984457539 {
		t.Errorf("Got %v. Expected %v", count, 26984457539)
	}
}
