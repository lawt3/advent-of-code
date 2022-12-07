package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed sample
var sample string

//go:embed input
var input string

func main() {
	fmt.Println("Sample")
	solve(sample)

	fmt.Println("Input")
	solve(input)
}

func solve(s string) {
	lines := strings.Split(s, "\n")
}
