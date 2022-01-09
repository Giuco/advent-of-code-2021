package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type coord struct {
	y int
	x int
}

type fold struct {
	direction string
	value     int
}

type paper [][]bool

func (p *paper) maxX() int {
	return len(*p)
}

func (p *paper) maxY() int {
	return len((*p)[0])
}

func (p *paper) size() (int, int) {
	return p.maxX(), p.maxY()
}

func (p *paper) asString() string {
	var result, v string

	for i := 0; i < p.maxX(); i++ {
		for j := 0; j < p.maxY(); j++ {
			if (*p)[i][j] {
				v = "#"
			} else {
				v = "."
			}
			result += v
		}
		result += "\n"
	}
	return result
}

func parseInput(text string) ([]coord, []fold) {
	var splitTemp []string
	var stringCoord, stringFold string
	var x, y, value int
	var parsedCoords []coord
	var parsedFolds []fold
	var err error

	splitTemp = strings.Split(text, "\n\n")
	stringCoord = splitTemp[0]
	stringFold = splitTemp[1]

	// Parsing the coordinates (first part of input)
	for _, line := range strings.Split(stringCoord, "\n") {
		splitTemp = strings.Split(line, ",")
		x, err = strconv.Atoi(splitTemp[0])
		if err != nil {
			log.Fatal("Could not parse x coord", err)
		}

		y, err = strconv.Atoi(splitTemp[1])
		if err != nil {
			log.Fatal("Could not parse y coord", err)
		}
		parsedCoords = append(parsedCoords, coord{x, y})
	}

	// Parsing the foldds (second part of the input)
	for _, line := range strings.Split(stringFold, "\n") {
		splitTemp = strings.Split(line, " ")
		splitTemp = strings.Split(splitTemp[2], "=")
		value, err = strconv.Atoi(splitTemp[1])
		if err != nil {
			log.Fatal("Could not parse value coord of fold", err)
		}
		parsedFolds = append(parsedFolds, fold{splitTemp[0], value})
	}

	return parsedCoords, parsedFolds
}

func readInput() ([]coord, []fold) {
	bs, err := ioutil.ReadFile("../data/day13.txt")
	if err != nil {
		panic(err)
	}
	raw_input := string(bs)
	return parseInput(raw_input)
}

func createEmptyPaper(x, y int) paper {
	var pap paper
	var l []bool

	for i := 0; i < x; i++ {
		l = []bool{}
		for j := 0; j < y; j++ {
			l = append(l, false)
		}
		pap = append(pap, l)
	}
	return pap
}

func buildPaper(coordinates []coord) paper {
	var maxX, maxY int

	// Getting the max coordinates
	for _, c := range coordinates {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}

	// Building the empty paper
	pap := createEmptyPaper(maxX+1, maxY+1)

	// Populating the paper with the marks
	for _, c := range coordinates {
		pap[c.x][c.y] = true
	}

	return pap
}

func cutPaper(pap paper, fo fold) (paper, paper) {
	var half1, half2 paper

	if fo.direction == "y" {
		half1 = createEmptyPaper(fo.value, pap.maxY())
		half2 = createEmptyPaper(fo.value, pap.maxY())
		for i := 0; i < half1.maxX(); i++ {
			for j := 0; j < half1.maxY(); j++ {
				half1[i][j] = pap[i][j]
				half2[i][j] = pap[pap.maxX()-i-1][j]
			}
		}
	} else {
		half1 = createEmptyPaper(pap.maxX(), fo.value)
		half2 = createEmptyPaper(pap.maxX(), fo.value)
		for i := 0; i < half1.maxX(); i++ {
			for j := 0; j < half1.maxY(); j++ {
				half1[i][j] = pap[i][j]
				half2[i][j] = pap[i][pap.maxY()-j-1]
			}
		}
	}

	return half1, half2
}

func foldPaper(pap paper, fo fold) paper {
	half1, half2 := cutPaper(pap, fo)
	newPaper := createEmptyPaper(half1.maxX(), half1.maxY())

	for i := 0; i < half1.maxX(); i++ {
		for j := 0; j < half1.maxY(); j++ {
			newPaper[i][j] = half1[i][j] || half2[i][j]
		}
	}

	return newPaper
}

func sumMatrix(matrix paper) int {
	var total int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] {
				total++
			}
		}
	}
	return total
}

func part1() int {
	coords, folds := readInput()
	paper := buildPaper(coords)
	foldedPaper := foldPaper(paper, folds[0])
	return sumMatrix(foldedPaper)
}

func part2() string {
	coords, folds := readInput()
	paper := buildPaper(coords)
	foldedPaper := foldPaper(paper, folds[0])

	for _, f := range folds[1:] {
		foldedPaper = foldPaper(foldedPaper, f)
	}

	return foldedPaper.asString()
}

func main() {
	fmt.Println("Part 1 -", part1())
	fmt.Println("Part 2 -")
	fmt.Println(part2())
}
