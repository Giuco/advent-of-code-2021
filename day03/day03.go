package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type binNum string
type bitList []binNum

func (b binNum) toDecimal() int64 {
	i, err := strconv.ParseInt(string(b), 2, 64)
	if err != nil {
		panic(err)
	} else {
		return i
	}
}

func (bl bitList) getNumSize() int {
	return len(bl[0])
}

func (bl bitList) getMostCommon(pos int) string {
	totalSum := 0
	size := len(bl)

	for _, i := range bl {
		ii, err := strconv.Atoi(string(i[pos]))
		if err != nil {
			panic(err)
		}
		totalSum += ii
	}

	var half int
	if size%2 == 0 {
		half = size / 2
	} else {
		half = size/2 + 1
	}

	if totalSum >= half {
		return "1"
	} else {
		return "0"
	}
}

func (bl bitList) getGamma() int64 {
	var gamma binNum

	for i := 0; i < bl.getNumSize(); i++ {
		most_common := bl.getMostCommon(i)
		if most_common == "1" {
			gamma += "1"
		} else {
			gamma += "0"
		}
	}
	return gamma.toDecimal()
}

func (bl bitList) getEpsilon() int64 {
	var epsilon binNum

	for i := 0; i < bl.getNumSize(); i++ {
		mostCommon := bl.getMostCommon(i)
		if mostCommon == "0" {
			epsilon += "1"
		} else {
			epsilon += "0"
		}
	}
	return epsilon.toDecimal()
}

func (bl bitList) getOxygen() int64 {
	binO2 := binNum("")
	filteredBL := bl
	for len(filteredBL) > 1 {
		for pos := 0; pos < bl.getNumSize(); pos++ {
			mostCommon := filteredBL.getMostCommon(pos)
			
		}
	}
	return binO2.toDecimal()
}

func (bl bitList) getCO2() int64 {

}

func parseInput(text string) bitList {
	bl := bitList{}
	for _, i := range strings.Split(text, "\n") {
		bl = append(bl, binNum(i))
	}
	return bl
}

func readInput() string {
	bs, err := ioutil.ReadFile("../data/day03.txt")
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(string(bs), "\r", "")
}

func main() {
	raw_input := readInput()
	input := parseInput(raw_input)
	gamma := input.getGamma()
	epsilon := input.getEpsilon()
	fmt.Printf("Part 1 - gamma=%v epsilon%v energy=%v", gamma, epsilon, gamma*epsilon)
}
