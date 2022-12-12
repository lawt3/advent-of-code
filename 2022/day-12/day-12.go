package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

func main() {
	grid := make(map[image.Point]rune)
	var start, end image.Point

	scanner := bufio.NewScanner(os.Stdin)
	var row int
	for scanner.Scan() {
		for col, r := range scanner.Text() {
			p := image.Point{col, row}
			if r == 'S' {
				start = p
				r = 'a'
			}
			if r == 'E' {
				end = p
				r = 'z'
			}
			grid[p] = r
		}
		row++
	}

	fmt.Println(part1(grid, start, end))
	fmt.Println(part2(grid, end))
}

func adjacent(a, b rune) bool {
	return b <= a+1
}

func part1(grid map[image.Point]rune, start, end image.Point) int {
	queue := []image.Point{start}
	dist := map[image.Point]int{start: 0}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, dir := range []image.Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
			next := curr.Add(dir)

			// Check in bounds
			if _, ok := grid[next]; !ok {
				continue
			}

			// Check seen
			if _, ok := dist[next]; ok {
				continue
			}

			// Check elevation
			if !adjacent(grid[curr], grid[next]) {
				continue
			}

			dist[next] = dist[curr] + 1
			queue = append(queue, next)
		}
	}

	return dist[end]
}

func part2(grid map[image.Point]rune, start image.Point) int {
	queue := []image.Point{start}
	dist := map[image.Point]int{start: 0}

	var curr image.Point

	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]

		if grid[curr] == 'a' {
			break
		}

		for _, dir := range []image.Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} {
			next := curr.Add(dir)

			// Check in bounds
			if _, ok := grid[next]; !ok {
				continue
			}

			// Check seen
			if _, ok := dist[next]; ok {
				continue
			}

			// Check elevation (backwards)
			if !adjacent(grid[next], grid[curr]) {
				continue
			}

			dist[next] = dist[curr] + 1
			queue = append(queue, next)
		}
	}

	return dist[curr]
}
