package main

import (
	"strings"
	"testing"
)

const rawInput = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func TestFindFirstWrongClosing(t *testing.T) {
	input := []string{
		"{([(<{}[<>[]}>{[]{[(<()>",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
	}
	expectedOutput := []string{"}", ")", "]", ")", ">"}

	for i := 0; i < len(input); i++ {
		e := expectedOutput[i]
		o := findFirstWrongClosing(input[i])

		if e != o {
			t.Errorf("Expected e: %v Actual e %v", e, o)
		}
	}
}

func TestPart1(t *testing.T) {
	input := parseInput(rawInput)
	output := runPart1(input)
	expected := 26397

	if output != expected {
		t.Errorf("Expect %v. Got %v", expected, output)
	}
}

func TestAutoComplete(t *testing.T) {
	input := []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"(((({<>}<{<{<>}{[]{[]{}",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}
	expectedOutput := []string{
		"}}]])})]",
		")}>]})",
		"}}>}>))))",
		"]]}}]}]}>",
		"])}>",
	}

	for i := 0; i < len(input); i++ {
		realOutput := autoComplete(input[i])
		joinedRealOutput := strings.Join(realOutput, "")
		if joinedRealOutput != expectedOutput[i] {
			t.Errorf("Expected: %v. Actual %v", expectedOutput[i], joinedRealOutput)
		}
	}
}

func TestRunPart2(t *testing.T) {
	input := parseInput(rawInput)
	output := runPart2(input)
	expected := 288957

	if output != expected {
		t.Errorf("Expect %v. Got %v", expected, output)
	}

}
