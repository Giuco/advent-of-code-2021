package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func runDays(fishes []int, days int) int {
	newFishs := make([]int, len(fishes))
	copy(newFishs, fishes)
	for d := 0; d < days; d++ {
		for i, v := range newFishs {
			if v == 0 {
				newFishs[i] = 6
				newFishs = append(newFishs, 8)
			} else {
				newFishs[i] -= 1
			}
		}
	}
	return len(newFishs)
}

func runDaysOptimized(fishes []int, days int) int {
	lifeCount := []int{}
	for l := 0; l <= 8; l++ {
		lifeCount = append(lifeCount, 0)
	}

	for _, f := range fishes {
		lifeCount[f] += 1
	}

	newLifeCount := []int{}
	for d := 0; d < days; d++ {
		newLifeCount = []int{}
		for _, l := range lifeCount[1:] {
			newLifeCount = append(newLifeCount, l)
		}
		newLifeCount = append(newLifeCount, lifeCount[0])
		newLifeCount[6] += lifeCount[0]
		lifeCount = newLifeCount
	}

	count := 0
	for _, l := range newLifeCount {
		count += l
	}
	return count
}

func readInput() []int {
	bs, err := ioutil.ReadFile("../data/day06.txt")
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
	fishes := readInput()
	nFishes := runDays(fishes, 80)
	fmt.Println("Part 1 -", nFishes)
	nFishes = runDaysOptimized(fishes, 256)
	fmt.Println("Part 2 -", nFishes)
}
