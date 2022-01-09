package main

import (
	"testing"
)

var rawTestInput = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`
var testInputCoord = []coord{
	{6, 10},
	{0, 14},
	{9, 10},
	{0, 3},
	{10, 4},
	{4, 11},
	{6, 0},
	{6, 12},
	{4, 1},
	{0, 13},
	{10, 12},
	{3, 4},
	{3, 0},
	{8, 4},
	{1, 10},
	{2, 14},
	{8, 10},
	{9, 0},
}
var testInputFold = []fold{
	{"y", 7},
	{"x", 5},
}
var testPaper = paper{
	{false, false, false, true, false, false, true, false, false, true, false},
	{false, false, false, false, true, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false},
	{true, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, true, false, false, false, false, true, false, true},
	{false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false},
	{false, false, false, false, false, false, false, false, false, false, false},
	{false, true, false, false, false, false, true, false, true, true, false},
	{false, false, false, false, true, false, false, false, false, false, false},
	{false, false, false, false, false, false, true, false, false, false, true},
	{true, false, false, false, false, false, false, false, false, false, false},
	{true, false, true, false, false, false, false, false, false, false, false},
}

func TestParseInput(t *testing.T) {
	gotCoord, gotFold := parseInput(rawTestInput)
	if len(gotCoord) != len(testInputCoord) {
		t.Errorf("CoordLen - Got: %v Expected: %v", len(gotCoord), len(testInputCoord))
	}

	if len(gotFold) != len(testInputFold) {
		t.Errorf("FoldLen - Got %v Expected %v", len(gotFold), len(testInputFold))
	}

	for i := 0; i < len(gotCoord); i++ {
		if gotCoord[i] != testInputCoord[i] {
			t.Errorf("i=%v Got %v Expected %v", i, gotCoord[i], testInputCoord[i])
		}
	}

	for i := 0; i < len(gotFold); i++ {
		if gotFold[i] != testInputFold[i] {
			t.Errorf("i=%v Got %v Expected %v", i, gotFold[i], testInputFold[i])
		}
	}

}

func TestBuildPaper(t *testing.T) {
	gotPaper := buildPaper(testInputCoord)
	for i := 0; i < len(gotPaper); i++ {
		for j := 0; j < len(gotPaper[i]); j++ {
			if gotPaper[i][j] != testPaper[i][j] {
				t.Errorf("i=%v, j=%v. Got %v. Expected %v", i, j, gotPaper[i][j], testPaper[i][j])
			}
		}
	}
}

func TestFoldPaper(t *testing.T) {
	paper := buildPaper(testInputCoord)
	fold := foldPaper(paper, testInputFold[0])
	total := sumMatrix(fold)
	if total != 17 {
		t.Errorf("Got %v, want %v", total, 17)
	}

	fold = foldPaper(fold, testInputFold[1])
	total = sumMatrix(fold)
	if total != 16 {
		t.Errorf("Got %v, want %v", total, 16)
	}
}
