package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func countIncreases(depths []int) int {
	var count int
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			count++
		}
	}
	return count
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func countIncreasesWindow(depths []int, window int) int {
	var count int
	for i := window; i < len(depths); i++ {
		if depths[i] > depths[i-window] {
			count++
		}
	}
	return count
}

func readInput() []int {
	bs, err := ioutil.ReadFile("../data/day01.txt")
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(0)
	}
	t1 := strings.Split(strings.ReplaceAll(string(bs), "\r\n", "\n"), "\n")
	t2 := []int{}
	for _, n := range t1 {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, i)
	}
	return t2
}

func main() {
	input := readInput()
	output := countIncreases(input)
	fmt.Println("Part 1 - The number of increases was:", output)
	output = countIncreasesWindow(input, 3)
	fmt.Println("Part 2 - The number of increases was:", output)
}
