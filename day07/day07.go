package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func calcMedian(numbers []int) int {
	sort.Ints(numbers)

	mNumber := len(numbers) / 2

	if len(numbers)%2 != 0 {
		return numbers[mNumber]
	}

	return (numbers[mNumber-1] + numbers[mNumber]) / 2
}

func calculateFuel(positions []int) int {
	median := calcMedian(positions)
	fuel := 0
	for _, pos := range positions {
		fuel += int(math.Abs(float64(pos - median)))
	}
	return fuel
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func calculateFuel2(positions []int) int {
	minPos, maxPos := MinMax(positions)
	minFuel := math.MaxInt64
	var fuel int

	for ref := minPos; ref <= maxPos; ref++ {
		fuel = 0
		for _, pos := range positions {
			for i := 1; i <= int(math.Abs(float64(pos-ref))); i++ {
				fuel += int(i)
			}
		}

		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return int(minFuel)
}

func readInput() []int {
	bs, err := ioutil.ReadFile("../data/day07.txt")
	if err != nil {
		panic(err)
	}
	raw_input := strings.Split(string(bs), ",")
	fishes := []int{}
	for _, f := range raw_input {
		i, err := strconv.Atoi(f)
		if err != nil {
			panic(err)
		}
		fishes = append(fishes, i)
	}
	return fishes
}

func main() {
	positions := readInput()
	fuel := calculateFuel(positions)
	fmt.Println("Part 1 -", fuel)
	fuel = calculateFuel2(positions)
	fmt.Println("Part 2 -", fuel)
}
