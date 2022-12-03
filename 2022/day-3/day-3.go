package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/samber/lo"
)

// There's a smarter way to implement this with two iterators
func intersect(l, r string) (res []byte) {
	for _, x := range l {
		for _, y := range r {
			if x == y {
				res = append(res, byte(x))
			}
		}
	}
	return
}

func priority(b byte) int {
	if b >= 'a' {
		return int(b-'a') + 1
	}

	return int(b-'A') + 27
}

func part1(rucksacks []string) (sum int) {
	for _, r := range rucksacks {
		n := len(r) / 2
		first := r[:n]
		second := r[n:]

		types := intersect(first, second)
		types = lo.Uniq(types)

		priorities := lo.Map(types, func(b byte, _ int) int {
			return priority(b)
		})

		sum += lo.Reduce(priorities, func(agg int, x int, _ int) int {
			return agg + x
		}, 0)
	}

	return
}

func part2(rucksacks []string) (sum int) {
	for i := 0; i < len(rucksacks); i += 3 {
		one := rucksacks[i]
		two := rucksacks[i+1]
		three := rucksacks[i+2]

		// In Go, strings don't convert to []char so it's better to use []string
		// strings.Split takes string -> []string
		// strings.Join takes []string -> string
		types := intersect(string(intersect(one, two)), three)
		sum += priority(types[0])
	}

	return
}

func main() {
	rucksacks := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rucksacks = append(rucksacks, scanner.Text())
	}

	fmt.Println(part1(rucksacks))
	fmt.Println(part2(rucksacks))
}
