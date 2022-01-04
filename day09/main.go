package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

type smokeMap [][]int
type coord struct {
	i int
	j int
}

func parseInput(text string) smokeMap {
	sm := smokeMap{}
	for _, l := range strings.Split(text, "\n") {
		cells := strings.Split(l, "")
		intCells := []int{}
		for _, cell := range cells {
			i, err := strconv.Atoi(cell)
			if err != nil {
				panic(err)
			}
			intCells = append(intCells, i)
		}
		sm = append(sm, intCells)
	}
	return sm
}

func readInput() smokeMap {
	bs, err := ioutil.ReadFile("../data/day09.txt")
	if err != nil {
		panic(err)
	}
	raw_input := string(bs)
	return parseInput(raw_input)
}

func findLowPoints(sm smokeMap) []coord {
	var lowPoints []coord
	var point int
	var isSmallest bool

	for i := 0; i < len(sm); i++ {
		for j := 0; j < len(sm[i]); j++ {
			isSmallest = true
			point = sm[i][j]
			if i > 0 && sm[i-1][j] <= point {
				isSmallest = false
			} else if i+1 < len(sm) && sm[i+1][j] <= point {
				isSmallest = false
			} else if j > 0 && sm[i][j-1] <= point {
				isSmallest = false
			} else if j+1 < len(sm[i]) && sm[i][j+1] <= point {
				isSmallest = false
			}

			if isSmallest == true {
				lowPoints = append(lowPoints, coord{i, j})
			}
		}
	}
	return lowPoints

}

func sumLowPoints(sm smokeMap) int {
	lowPoints := findLowPoints(sm)
	total := 0

	for _, point := range lowPoints {
		total += sm[point.i][point.j] + 1
	}
	return total

}

func findBasinSizes(sm smokeMap) []int {
	var sizes []int
	var smCopy smokeMap
	lowPoints := findLowPoints(sm)

	for _, lp := range lowPoints {
		smCopy = append(smokeMap{}, sm...)
		sizes = append(sizes, getBasinSize(smCopy, lp))
	}

	return sizes
}

func getBasinSize(sm smokeMap, low coord) int {
	size := 0
	_, size = explore(sm, low, size)
	return size
}

func explore(sm smokeMap, p coord, size int) (smokeMap, int) {
	if sm[p.i][p.j] != 9 {
		size++
		sm[p.i][p.j] = 9
		if p.i > 0 {
			sm, size = explore(sm, coord{p.i - 1, p.j}, size)
		}

		if p.i+1 < len(sm) {
			sm, size = explore(sm, coord{p.i + 1, p.j}, size)
		}

		if p.j > 0 {
			sm, size = explore(sm, coord{p.i, p.j - 1}, size)
		}

		if p.j+1 < len(sm[0]) {
			sm, size = explore(sm, coord{p.i, p.j + 1}, size)
		}
	}
	return sm, size
}

func getPart2(sm smokeMap) int {
	basinSizes := findBasinSizes(sm)
	sort.Ints(basinSizes)
	return basinSizes[len(basinSizes)-1] * basinSizes[len(basinSizes)-2] * basinSizes[len(basinSizes)-3]
}

func main() {
	input := readInput()
	fmt.Println("Part 1 -", sumLowPoints(input))
	fmt.Println("Part 2 -", getPart2(input))
}
