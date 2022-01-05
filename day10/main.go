package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func parseInput(text string) []string {
	return strings.Split(text, "\n")
}

func readInput() []string {
	bs, err := ioutil.ReadFile("../data/day10.txt")
	if err != nil {
		panic(err)
	}
	raw_input := string(bs)
	return parseInput(raw_input)
}

func isOpening(c string) bool {
	return c == "(" || c == "[" || c == "{" || c == "<"
}

func getInvert(c string) string {
	mapper := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",

		")": "(",
		"]": "[",
		"}": "{",
		">": "<",
	}
	return mapper[c]
}

func findFirstWrongClosing(text string) string {
	opens := []string{}

	for _, r := range strings.Split(text, "") {
		if isOpening(r) {
			opens = append(opens, r)
		} else if getInvert(r) != opens[len(opens)-1] {
			return r
		} else {
			opens = opens[:len(opens)-1]
		}
	}
	return ""
}

func getClosingPoints(c string) int {
	points := map[string]int{
		"":  0,
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	if isOpening(c) {
		return points[getInvert(c)]
	} else {
		return points[c]
	}
}

func runPart1(inputs []string) int {
	total := 0

	for _, x := range inputs {
		total += getClosingPoints(findFirstWrongClosing(x))
	}

	return total
}

func autoComplete(text string) []string {
	var opens, closes []string

	for _, r := range strings.Split(text, "") {
		if isOpening(r) {
			opens = append(opens, r)
		} else {
			opens = opens[:len(opens)-1]
		}
	}

	for i := len(opens) - 1; i >= 0; i-- {
		closes = append(closes, getInvert(opens[i]))
	}

	return closes
}

// func runPart2(inputs []string) int {
// 	var total int
// 	var autoCompleted []string

// 	for _, x := range inputs {
// 		if findFirstWrongClosing(x) != "" {
// 			autoCompleted = autoComplete(x)
// 			total *= 5
// 			total += getAutoCompletePoints(autoCompleted)
// 		}
// 	}
// 	return total
// }

// func getAutoCompletePoints(c string) int {
// 	points := map[string]int{
// 		")": 1,
// 		"]": 2,
// 		"}": 3,
// 		">": 4,
// 	}
// 	return points[c]
// }

func main() {
	input := readInput()
	fmt.Println("Part 1 -", runPart1(input))
}
