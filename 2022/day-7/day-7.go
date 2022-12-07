package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
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

type stream struct {
	lines []string
	pos   int
}

func (s *stream) NextLine() []string {
	defer func() { s.pos++ }()
	return strings.Fields(s.lines[s.pos])
}

func (s *stream) Done() bool {
	return s.pos >= len(s.lines)
}

func (s *stream) PeekCmd() bool {
	return s.Done() || strings.ContainsRune(s.lines[s.pos], '$')
}

type node struct {
	parent   *node
	children map[string]*node
	size     int
}

func NewNode(parent *node) *node {
	return &node{
		parent:   parent,
		children: make(map[string]*node),
	}
}

func propagateSize(n *node) int {
	if n.children == nil {
		return n.size
	}

	for _, c := range n.children {
		n.size += propagateSize(c)
	}
	return n.size
}

func traverse(n *node, fn func(*node)) {
	if n.children == nil {
		fn(n)
		return
	}

	for _, c := range n.children {
		fn(c)
		traverse(c, fn)
	}
}

func solve(s string) {
	// Parse input and construct tree
	input := &stream{lines: strings.Split(s, "\n")}
	tree := NewNode(nil)
	curr := tree

	for !input.Done() {
		cmd := input.NextLine()
		switch cmd[1] {
		case "cd":
			switch cmd[2] {
			case "/":
				continue
			case "..":
				curr = curr.parent
			default:
				curr = curr.children[cmd[2]]
			}
		case "ls":
			for !input.PeekCmd() {
				tokens := input.NextLine()
				if tokens[0] == "dir" {
					curr.children[tokens[1]] = NewNode(curr)
				} else {
					size, _ := strconv.Atoi(tokens[0]) // can't fail
					curr.size += size
				}
			}
		}
	}

	fmt.Println("total", propagateSize(tree))

	fmt.Println(part1(tree))
	fmt.Println(part2(tree))
}

func part1(tree *node) int {
	// Filter <= 100,000 and sum those directories
	var ans int
	traverse(tree, func(n *node) {
		if n.size <= 100_000 {
			ans += n.size
		}
	})

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func part2(tree *node) int {
	// Find the smallest directory that gives enough room
	need := 30_000_000 - (70_000_000 - tree.size)

	ans := math.MaxInt
	traverse(tree, func(n *node) {
		if n.size >= need {
			ans = min(ans, n.size)
		}
	})

	return ans
}
