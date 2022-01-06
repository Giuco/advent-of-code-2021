package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInput(text string) [][]int {
	var output [][]int
	var intLine []int

	for _, line := range strings.Split(text, "\n") {
		intLine = []int{}
		for _, c := range strings.Split(line, "") {
			i, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			intLine = append(intLine, i)
		}
		output = append(output, intLine)
	}
	return output
}

func readInput() [][]int {
	bs, err := ioutil.ReadFile("../data/day11.txt")
	if err != nil {
		panic(err)
	}
	raw_input := string(bs)
	return parseInput(raw_input)
}

func runFlash(octoMap [][]int, i, j int) {
	octoMap[i][j] = 0

	if i > 0 {
		if octoMap[i-1][j] > 0 {
			octoMap[i-1][j]++
		}
	}

	if j > 0 {
		if octoMap[i][j-1] > 0 {
			octoMap[i][j-1]++
		}
	}

	if i+1 < len(octoMap) {
		if octoMap[i+1][j] > 0 {
			octoMap[i+1][j]++
		}
	}

	if j+1 < len(octoMap[i]) {
		if octoMap[i][j+1] > 0 {
			octoMap[i][j+1]++
		}
	}

	if i > 0 && j > 0 {
		if octoMap[i-1][j-1] > 0 {
			octoMap[i-1][j-1]++
		}
	}

	if i+1 < len(octoMap) && j+1 < len(octoMap[i]) {
		if octoMap[i+1][j+1] > 0 {
			octoMap[i+1][j+1]++
		}
	}

	if i > 0 && j+1 < len(octoMap[i]) {
		if octoMap[i-1][j+1] > 0 {
			octoMap[i-1][j+1]++
		}
	}

	if i+1 < len(octoMap) && j > 0 {
		if octoMap[i+1][j-1] > 0 {
			octoMap[i+1][j-1]++
		}
	}
}

func runStep(octoMap [][]int) int {
	var noFlash bool
	var count int

	for i := 0; i < len(octoMap); i++ {
		for j := 0; j < len(octoMap[i]); j++ {
			octoMap[i][j]++
		}
	}

	for {
		noFlash = true
		for i, row := range octoMap {
			for j, cell := range row {
				if cell > 9 {
					runFlash(octoMap, i, j)
					noFlash = false
					count++
				}
			}
		}

		if noFlash {
			break
		}
	}

	return count
}

func runNSteps(octoMap [][]int, nSteps int) int {
	var count int
	for i := 0; i < nSteps; i++ {
		count += runStep(octoMap)
	}
	return count
}

func formatOctoMap(octoMap [][]int) string {
	var output string
	for _, line := range octoMap {
		for _, c := range line {
			output += strconv.Itoa(c)
		}
		output += "\n"
	}
	return output
}

func getSyncStep(octoMap [][]int) int {
	var n int
	var allEqual bool

	for k := 0; k < 1000; k++ {
		runStep(octoMap)
		
		allEqual = true
		n = octoMap[0][0]
		for i := 0; i < len(octoMap); i++ {
			for j := 0; j < len(octoMap[i]); j++ {
				if octoMap[i][j] != n {
					allEqual = false
					break
				}
			}
			if !allEqual {
				break
			}
		}

		if allEqual {
			return k + 1
		}
	}

	panic("Max iter")
}

func main() {
	octoMap := readInput()
	fmt.Println("Part 1 -", runNSteps(octoMap, 100))
	octoMap = readInput()
	fmt.Println("Part 2 -", getSyncStep(octoMap))
}
