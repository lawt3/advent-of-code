package main

import (
	_ "embed"
	"fmt"
	"image"
	"strconv"
	"strings"
)

var dirs = []image.Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

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
	grid := make(map[image.Point]int)

	lines := strings.Split(s, "\n")
	for row, line := range lines {
		for col, s := range strings.Split(line, "") {
			height, _ := strconv.Atoi(s)
			grid[image.Point{row, col}] = height
		}
	}

	fmt.Println(part1(grid))
	fmt.Println(part2(grid))
}

func part1(grid map[image.Point]int) int {
	var acc int
	for point := range grid {
		if visible(point, grid) {
			acc++
		}
	}

	return acc
}

func part2(grid map[image.Point]int) int {
	var acc int
	for point := range grid {
		s := score(point, grid)
		if s > acc {
			acc = s
		}
	}

	return acc
}

func visible(point image.Point, grid map[image.Point]int) bool {
	for _, dir := range dirs {
		for np := point.Add(dir); ; np = np.Add(dir) {
			height, ok := grid[np]
			if !ok {
				return true
			}
			if height >= grid[point] {
				break
			}
		}
	}

	return false
}

func score(point image.Point, grid map[image.Point]int) int {
	acc := 1
	for _, dir := range dirs {
		var dist int
		for np := point.Add(dir); ; np = np.Add(dir) {
			height, ok := grid[np]
			if !ok {
				break
			}
			dist++
			if height >= grid[point] {
				break
			}
		}
		acc *= dist
	}

	return acc
}
