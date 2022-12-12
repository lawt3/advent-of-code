package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var buffer []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buffer = append(buffer, scanner.Text())
	}

	solve(buffer)
}

type CPU struct {
	cycle, X, strength int
	screen             string
}

func newCPU() *CPU {
	return &CPU{X: 1}
}

func (c *CPU) inc() {
	// Part 2
	if col := c.cycle % 40; col >= c.X-1 && col <= c.X+1 {
		c.screen += "#"
	} else {
		c.screen += "."
	}

	c.cycle += 1

	// Part 1
	if c.cycle%40 == 20 {
		c.strength += c.cycle * c.X
	}

}

func (c *CPU) add(X int) {
	c.X += X
}

func solve(buffer []string) {
	// Initial state
	cpu := newCPU()

	for _, line := range buffer {
		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case "noop":
			cpu.inc()
		case "addx":
			cpu.inc()
			cpu.inc()
			v, _ := strconv.Atoi(tokens[1])
			cpu.add(v)
		default:
			panic("unknown instruction")
		}
	}

	fmt.Println("strength", cpu.strength)
	for i := 1; i <= 240; i++ {
		fmt.Print(string(cpu.screen[i-1]))
		if i%40 == 0 {
			fmt.Println()
		}
	}
}
