package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type connection struct {
	start string
	end   string
}

type Cave struct {
	name  string
	paths []*Cave
}

func (c *Cave) isBig() bool {
	if strings.ToUpper(c.name) == c.name {
		return true
	}

	return false
}

func ParseInput(text string) []connection {
	var cs []string
	connections := []connection{}

	for _, c := range strings.Split(text, "\n") {
		cs = strings.Split(c, "-")
		connections = append(connections, connection{cs[0], cs[1]})
	}

	return connections
}

func ReadInput() []connection {
	bs, err := ioutil.ReadFile("../data/day12.txt")
	if err != nil {
		panic(err)
	}
	raw_input := string(bs)
	return ParseInput(raw_input)
}

func BuildMap(conns []connection) map[string]*Cave {
	var startCave, endCave *Cave
	var ok bool

	caveDict := map[string]*Cave{}
	for _, c := range conns {
		startCave, ok = caveDict[c.start]
		if !ok {
			startCave = &Cave{name: c.start}
			caveDict[c.start] = startCave
		}
		endCave, ok = caveDict[c.end]
		if !ok {
			endCave = &Cave{name: c.end}
			caveDict[c.end] = endCave
		}

		startCave.paths = append(startCave.paths, endCave)
		endCave.paths = append(endCave.paths, startCave)
		// fmt.Println(caveDict)
	}

	return caveDict
}

func copyVisitCount(m map[string]int) map[string]int {
	cp := make(map[string]int)
	for k, v := range m {
		cp[k] = v
	}
	return cp
}

func Explore1(start *Cave, count int, visitCount map[string]int, path []string) (int, map[string]int, []string) {
	var newVisitCount map[string]int
	visitCount[start.name]++
	path = append(path, start.name)
	// fmt.Println("-------")

	for _, next := range start.paths {
		newVisitCount = copyVisitCount(visitCount)
		// fmt.Println(path, count, visitCount)
		// fmt.Println("Current", start.name, "Next", next.name)
		if next.name == "end" {
			count++
		} else if next.isBig() {
			count, _, _ = Explore1(next, count, newVisitCount, path)
		} else if visitCount[next.name] == 0 && next.name != "start" {
			count, _, _ = Explore1(next, count, newVisitCount, path)
		}
	}

	return count, visitCount, path
}

func Explore2(start *Cave, count int, visitCount map[string]int, path []string) (int, map[string]int, []string) {
	var newVisitCount map[string]int
	var visitedSmallTwoTime bool
	visitCount[start.name]++
	for k, v := range visitCount {
		if !(strings.ToUpper(k) == k) && v >= 2 {
			visitedSmallTwoTime = true
		}
	}
	path = append(path, start.name)
	// fmt.Println("-------")

	for _, next := range start.paths {
		newVisitCount = copyVisitCount(visitCount)
		// fmt.Println(path, count, visitCount, visitedSmallTwoTime)
		// fmt.Println("Current", start.name, "Next", next.name)
		if next.name == "end" {
			count++
		} else if next.name == "start" {
			continue
		} else if next.isBig() {
			count, _, _ = Explore2(next, count, newVisitCount, path)
		} else if visitCount[next.name] == 0 {
			count, _, _ = Explore2(next, count, newVisitCount, path)
		} else if !visitedSmallTwoTime {
			count, _, _ = Explore2(next, count, newVisitCount, path)
		}
	}

	return count, visitCount, path
}

func CountPaths1(conns []connection) int {
	caveDict := BuildMap(conns)
	count, _, _ := Explore1(caveDict["start"], 0, map[string]int{}, []string{})
	return count
}

func CountPaths2(conns []connection) int {
	caveDict := BuildMap(conns)
	count, _, _ := Explore2(caveDict["start"], 0, map[string]int{}, []string{})
	return count
}

func main() {
	input := ReadInput()
	fmt.Println("Part 1 -", CountPaths1(input))
	fmt.Println("Part 2 -", CountPaths2(input))
}
