package main

import (
	"testing"
)

var rawInputs = []string{
	`start-A
start-b
A-c
A-b
b-d
A-end
b-end`,
	`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`,
	`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`,
}
var expectedPathCount1 = []int{10, 19, 226}
var expectedPathCount2 = []int{36, 103, 3509}
var expectedCaveMap = map[string][]string{
	"start": {"A", "b"},
	"A":     {"start", "c", "end", "b"},
	"c":     {"A"},
	"b":     {"start", "A", "d", "end"},
	"d":     {"b"},
	"end":   {"A", "b"},
}

func TestParseInput(t *testing.T) {
	got := ParseInput(rawInputs[0])
	expected := []connection{
		{"start", "A"},
		{"start", "b"},
		{"A", "c"},
		{"A", "b"},
		{"b", "d"},
		{"A", "end"},
		{"b", "end"},
	}

	if len(got) != len(expected) {
		t.Errorf("Got %v, expected %v", len(got), len(expected))
	}

	for i := 0; i < len(got); i++ {
		if got[i] != expected[i] {
			t.Errorf("Got %v, expected %v. On index=%v", got[i], expected[i], i)
		}
	}
}

func TestBuildMap(t *testing.T) {
	var cave *Cave
	var ok, foundAll, foundSpecificCave bool

	input := ParseInput(rawInputs[0])
	got := BuildMap(input)
	// fmt.Println(got)

	for name, paths := range expectedCaveMap {
		// fmt.Println("Cave:", name)
		cave, ok = got[name]
		if !ok {
			t.Errorf("Could not find cave: %v", name)
		}

		if len(paths) != len(cave.paths) {
			t.Errorf("Expected %v size, Got %v size. Cave: %v", len(paths), len(cave.paths), name)
			continue
		}

		foundAll = true
		for _, expectedCaveName := range paths {
			// fmt.Println("Finding path to:", expectedCaveName)
			foundSpecificCave = false
			for _, foundCave := range cave.paths {
				// fmt.Println("Found path to", foundCave.name)
				if foundCave.name == expectedCaveName {
					foundSpecificCave = true
					break
				}
			}
			if !foundSpecificCave {
				foundAll = false
				break
			}
		}

		if !foundAll {
			t.Errorf("Could not find all paths for cave %v", name)
		}
	}
}

func TestCountPaths1(t *testing.T) {
	for i := 0; i < len(rawInputs); i++ {
		input := ParseInput(rawInputs[i])
		got := CountPaths1(input)
		expected := expectedPathCount1[i]

		if got != expected {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	}
}

func TestCountPaths2(t *testing.T) {
	for i := 0; i < len(rawInputs); i++ {
		input := ParseInput(rawInputs[i])
		got := CountPaths2(input)
		expected := expectedPathCount2[i]

		if got != expected {
			t.Errorf("Got %v, expected %v", got, expected)
		}
	}
}
