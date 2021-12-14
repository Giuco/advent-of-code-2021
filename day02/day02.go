package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type command struct {
	direction string
	units     int
}

type position struct {
	depth      int
	horizontal int
	aim        int
}

func (p position) move(c command) position {
	if c.direction == "forward" {
		p.horizontal += c.units
	} else if c.direction == "up" {
		p.depth -= c.units
	} else if c.direction == "down" {
		p.depth += c.units
	} else {
		panic("Unexpected direction")
	}
	return p
}

func (p position) moveAim(c command) position {
	if c.direction == "forward" {
		p.depth += p.aim * c.units
		p.horizontal += c.units
	} else if c.direction == "up" {
		p.aim -= c.units
	} else if c.direction == "down" {
		p.aim += c.units
	} else {
		panic("Unexpected direction")
	}
	return p
}

func (p position) result() int {
	return p.depth * p.horizontal
}

func parseInput(text string) []command {
	lines := strings.Split(text, "\n")
	commands := []command{}
	for _, l := range lines {
		objs := strings.Split(l, " ")
		unit, err := strconv.Atoi(objs[1])
		if err != nil {
			panic(err)
		}
		c := command{direction: objs[0], units: unit}
		commands = append(commands, c)
	}
	return commands
}

func readInput() string {
	bs, err := ioutil.ReadFile("../data/day02.txt")
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(string(bs), "\r", "")
}

func moveShip(commands []command, aim bool) position {
	p := position{}
	for _, c := range commands {
		if aim {
			p = p.moveAim(c)
		} else {
			p = p.move(c)
		}
	}
	return p
}

func main() {
	raw_input := readInput()
	input := parseInput(raw_input)
	output := moveShip(input, false)
	fmt.Println("Part 1 - Output: ", output.result())
	output = moveShip(input, true)
	fmt.Println("Part 2 - Output: ", output.result())
}
