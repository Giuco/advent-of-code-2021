package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type entry struct {
	first  []string
	second []string
}

type stone map[string]string

func parseInput(text string) []entry {
	lines := strings.Split(text, "\n")
	var firstRaw string
	var first []string
	var secondRaw string
	var second []string
	var separated []string
	var toReturn []entry

	for _, line := range lines {
		separated = strings.Split(line, " | ")
		firstRaw = separated[0]
		secondRaw = separated[1]

		first = strings.Split(firstRaw, " ")
		second = strings.Split(secondRaw, " ")

		for i := 0; i < len(first); i++ {
			first[i] = strings.TrimSpace(first[i])
		}

		for i := 0; i < len(second); i++ {
			second[i] = strings.TrimSpace(second[i])
		}

		toReturn = append(toReturn, entry{first: first, second: second})

	}
	return toReturn
}

func readInput() []entry {
	bs, err := ioutil.ReadFile("../data/day08.txt")
	if err != nil {
		panic(err)
	}
	raw_input := string(bs)
	return parseInput(raw_input)
}

func count1478(entries []entry) int {
	count := 0
	for _, e := range entries {
		for _, s := range e.second {
			if len(s) == 2 || len(s) == 4 || len(s) == 3 || len(s) == 7 {
				count++
			}
		}
	}
	return count
}

func translateEntry(first []string) stone {
	countsChrs := map[rune]int{}
	rosettaStone := stone{}

	for _, seq := range first {
		for _, digit := range seq {
			countsChrs[digit]++
		}
	}

	for _, seq := range first {
		if len(seq) == 2 {
			rosettaStone[sortSeqtring(seq)] = "1"
		} else if len(seq) == 3 {
			rosettaStone[sortSeqtring(seq)] = "7"
		} else if len(seq) == 4 {
			rosettaStone[sortSeqtring(seq)] = "4"
		} else if len(seq) == 7 {
			rosettaStone[sortSeqtring(seq)] = "8"
		} else if len(seq) == 5 {
			for _, chr := range seq {
				if countsChrs[chr] == 4 {
					rosettaStone[sortSeqtring(seq)] = "2"
				} else if countsChrs[chr] == 6 {
					rosettaStone[sortSeqtring(seq)] = "5"
				}
			}

			if _, twoOrFive := rosettaStone[sortSeqtring(seq)]; !twoOrFive {
				rosettaStone[sortSeqtring(seq)] = "3"
			}
		} else if len(seq) == 6 {
			countSegmentsWithKeyInstances := make(map[int]int)
			for _, segment := range seq {
				countSegmentsWithKeyInstances[countsChrs[segment]]++
			}
			if countSegmentsWithKeyInstances[4] == 0 {
				rosettaStone[sortSeqtring(seq)] = "9"
			} else if countSegmentsWithKeyInstances[7] == 1 {
				rosettaStone[sortSeqtring(seq)] = "0"
			} else {
				rosettaStone[sortSeqtring(seq)] = "6"
			}
		}

	}

	return rosettaStone
}

func decodeEntry(entries []string, stone stone) int {
	decodedEntry := ""
	for _, entry := range entries {
		decodedEntry += stone[sortSeqtring(entry)]
	}
	output, err := strconv.Atoi(decodedEntry)
	if err != nil {
		panic(err)
	}
	return output
}

func decodeAll(entries []entry) int {
	sumDecoded := 0
	for _, entry := range entries {
		stone := translateEntry(entry.first)
		decoded := decodeEntry(entry.second, stone)
		sumDecoded += decoded
	}
	return sumDecoded
}

func sortSeqtring(s string) string {
	sorted := ""
	for _, l := range []string{"a", "b", "c", "d", "e", "f", "g"} {
		if strings.Contains(s, l) {
			sorted += l
		}
	}
	return sorted
}

func main() {
	input := readInput()
	fmt.Println("Part 1: ", count1478(input))
	fmt.Println("Part 2", decodeAll(input))
}
