package main

import (
	_ "embed"
	"fmt"
	"strings"
)

const (
	readStacks = iota
	readMoves
)

type move struct {
	src int
	dst int
	n   int
}

//go:embed sample
var sample string

//go:embed input
var input string

func pop(s string) (string, string) {
	return s[:1], s[1:]
}

func chop(s string, n int) (string, string) {
	return s[:n], s[n:]
}

func main() {
	stacks := make([]string, 3)
	var moves []move

	state := readStacks

	lines := strings.Split(sample, "\n")
	for _, line := range lines {
		fmt.Println("current line", line)
		switch state {
		case readStacks:
			if line == "" {
				fmt.Println("changing state to read moves")
				state = readMoves
				break
			}

			tokens := strings.Split(line, "")
			for i, t := range tokens {
				switch t {
				case " ", "[", "]":
					continue
				default:
					index := (i - 1) / 4
					stacks[index] += t
				}
			}
		case readMoves:
			var i, j, n int
			fmt.Sscanf(line, "move %d from %d to %d", &n, &i, &j)
			moves = append(moves, move{i - 1, j - 1, n})
		}
	}

	for _, m := range moves {
		for i := 0; i < m.n; i++ {
			c, rest := pop(stacks[m.src])
			stacks[m.dst] = c + stacks[m.dst]
			stacks[m.src] = rest
		}

		// xs, rest := chop(stacks[m.src], m.n)
		// stacks[m.dst] = xs + stacks[m.dst]
		// stacks[m.src] = rest
	}

	for _, s := range stacks {
		fmt.Print(strings.Split(s, "")[0])
	}
	fmt.Println()
}
