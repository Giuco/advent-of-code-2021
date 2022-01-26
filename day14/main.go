package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type instruction struct {
	pair   string
	insert string
}

func parseInput(text string) (string, []instruction) {
	var instructions []instruction
	var split []string

	lines := strings.Split(text, "\n")

	for _, line := range lines[1:] {
		split = strings.Split(line, " -> ")
		instructions = append(instructions, instruction{split[0], split[1]})
	}

	return lines[0], instructions
}


func readInput() (string, []instruction) {
	bs, err := ioutil.ReadFile("../data/day14.txt")
	if err != nil {
		panic(err)
	}
	raw_input := string(bs)
	return parseInput(raw_input)
}

func runNSteps(seq string, insts []instruction, nSteps int) map[rune]int {
	pairCount := map[string]int{}
	runeCount := map[rune]int{}

	for i := 0; i < len(seq)-1; i++ {
		pairCount[seq[i:i+2]]++
		runeCount[rune(seq[i])]++
	}

	runeCount[rune(seq[len(seq)-1])]++

	for i := 0; i < nSteps; i++ {
		fmt.Println("Step:", i)
		newPairCount := make(map[string]int)
		for pair, count := range pairCount {
			for _, instruction := range insts {
				if pair == instruction.pair {
					newPairCount[pair[0:1]+instruction.insert] += count
					newPairCount[instruction.insert+pair[1:2]] += count
					runeCount[rune(instruction.insert[0])] += count
				}
			}
		}
		pairCount = newPairCount
	}

	return runeCount
}


func partN(seq string, insts []instruction, nSteps int) int {
	count := runNSteps(seq, insts, nSteps)

	min := count[rune(seq[0])]
	max := count[rune(seq[0])]

	for _, v := range count {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return max - min
}

func main() {
	initial, instructions := readInput()
	fmt.Println("Part 1", partN(initial, instructions, 10))
	fmt.Println("Part 2", partN(initial, instructions, 40))
}
