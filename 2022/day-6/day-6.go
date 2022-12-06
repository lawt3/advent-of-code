package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/samber/lo"
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

func solve(input string) {
	fmt.Println(part1(input) + 1)
	fmt.Println(part2(input) + 1)
}

type uniq struct {
	arr []string
	n   int
}

func NewUniq(xs []string) *uniq {
	n := len(lo.Uniq(xs))
	return &uniq{xs, n}
}

func (u *uniq) update(s string) {
	u.arr = append(u.arr[1:], s)
	u.n = len(lo.Uniq(u.arr))
}

// Should return when four chars are uniq
func part1(s string) int {
	xs := strings.Split(s, "")

	u := NewUniq(xs[:4])
	for i := 4; i < len(xs); i++ {
		u.update(xs[i])
		if u.n == 4 {
			return i
		}
	}

	return 0
}

func part2(s string) int {
	xs := strings.Split(s, "")

	u := NewUniq(xs[:14])
	for i := 14; i < len(xs); i++ {
		u.update(xs[i])
		if u.n == 14 {
			return i
		}
	}

	return 0
}
