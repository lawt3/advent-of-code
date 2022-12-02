package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	shapeScore = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
		"A": 1,
		"B": 2,
		"C": 3,
	}

	outcomeScore = map[string]int{
		"W": 6,
		"D": 3,
		"L": 0,
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	index = map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"X": 0,
		"Y": 1,
		"Z": 2,
	}

	choices = [3]string{"A", "B", "C"}
)

type Round struct {
	l string
	r string
}

func play1(r Round) string {
	switch {
	case index[r.l] == index[r.r]:
		return "D"
	case (index[r.l]+1)%3 == index[r.r]:
		return "W"
	default:
		return "L"
	}
}

func part1(r Round) int {
	return shapeScore[r.r] + outcomeScore[play1(r)]
}

func play2(r Round) string {
	switch {
	case r.r == "Y":
		return r.l
	case r.r == "X":
		// In Go, % doesn't work for negative ints!
		// Need an extra layer of %
		return choices[((index[r.l]-1)%3+3)%3]
	default:
		return choices[(index[r.l]+1)%3]
	}
}

func part2(r Round) int {
	return shapeScore[play2(r)] + outcomeScore[r.r]
}

func main() {
	strategy := make([]Round, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		strategy = append(strategy, Round{string(text[0]), string(text[2])})
	}

	var sum1 int
	var sum2 int
	for _, s := range strategy {
		sum1 += part1(s)
		sum2 += part2(s)
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}
