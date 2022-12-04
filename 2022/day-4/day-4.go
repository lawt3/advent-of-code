package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Range struct {
	start int
	end   int
}

func part1(first, second []Range) int {
	var count int
	for i := range first {
		a := first[i].start
		b := first[i].end

		c := second[i].start
		d := second[i].end

		if (a <= c && d <= b) || (c <= a && b <= d) {
			count++
		}
	}

	return count
}

func part2(first, second []Range) int {
	var overlap int
	for i := range first {
		a := first[i].start
		b := first[i].end

		c := second[i].start
		d := second[i].end

		if (a <= c && c <= b) || (c <= a && a <= d) {
			overlap++
		}
	}

	return overlap
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	first := make([]Range, 0)
	second := make([]Range, 0)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")

		var a, b int
		fmt.Sscanf(parts[0], "%d-%d", &a, &b)
		first = append(first, Range{a, b})

		var c, d int
		fmt.Sscanf(parts[1], "%d-%d", &c, &d)
		second = append(second, Range{c, d})
	}

	fmt.Println(part1(first, second))
	fmt.Println(part2(first, second))
}
