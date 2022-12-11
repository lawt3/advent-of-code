package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

var dirs = map[string]image.Point{
	"R": {1, 0},
	"U": {0, 1},
	"L": {-1, 0},
	"D": {0, -1},
}

type move struct {
	dir   image.Point
	steps int
}

func main() {
	var moves []move

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var dir rune
		var steps int
		fmt.Sscanf(scanner.Text(), "%c %d", &dir, &steps)

		moves = append(moves, move{dirs[string(dir)], steps})
	}

	solve(moves)
}

func grid(rows, cols int, points map[image.Point]struct{}) {
	for row := rows; row >= 0; row-- {
		for col := 0; col < cols; col++ {
			if _, ok := points[image.Point{col, row}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func solve(moves []move) {
	part1(moves)
	part2(moves)
}

func part1(moves []move) {
	points := map[image.Point]struct{}{{0, 0}: {}}

	var head, tail, prev image.Point
	for _, m := range moves {
		for i := 0; i < m.steps; i++ {
			prev = head
			head = head.Add(m.dir)
			if !adjacent(head, tail) {
				tail = prev
				points[tail] = struct{}{}
			}
		}
	}

	fmt.Println(len(points))
	grid(6, 6, points)
}

func part2(moves []move) {
	points := map[image.Point]struct{}{{0, 0}: {}}
	rope := make([]image.Point, 10)

	for _, m := range moves {
		for i := 0; i < m.steps; i++ {
			rope[0] = rope[0].Add(m.dir)

			for j := 1; j < len(rope); j++ {
				if d := rope[j-1].Sub(rope[j]); abs(d.X) > 1 || abs(d.Y) > 1 {
					rope[j] = rope[j].Add(image.Point{sgn(d.X), sgn(d.Y)})
				}
			}

			points[rope[len(rope)-1]] = struct{}{}
		}
	}

	fmt.Println(len(points))
	grid(6, 6, points)
}

func adjacent(a, b image.Point) bool {
	delta := a.Sub(b)
	return abs(delta.X) <= 1 && abs(delta.Y) <= 1
}

func abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}

func sgn(x int) int {
	switch {
	case x < 0:
		return -1
	case x == 0:
		return 0
	default:
		return 1
	}
}
