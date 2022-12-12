package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items                   []int
	op                      func(int) int
	test                    func(int) bool
	trueMonkey, falseMonkey int
}

func main() {
	var monkeys []*monkey

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Monkey") {
			monkeys = append(monkeys, parseMonkey(scanner))
		}
	}

	// part1(monkeys)
	part2(monkeys)
}

func parseMonkey(scanner *bufio.Scanner) *monkey {
	var m monkey

	scanner.Scan()
	startingItems := strings.Split(scanner.Text(), ": ")[1]
	tokens := strings.Split(startingItems, ", ")
	for _, item := range tokens {
		item, _ := strconv.Atoi(item)
		m.items = append(m.items, item)
	}

	scanner.Scan()
	var operator, operand string
	fmt.Sscanf(scanner.Text(), "  Operation: new = old %s %s", &operator, &operand)
	m.op = func(item int) int {
		var v int
		if operand == "old" {
			v = item
		} else {
			v, _ = strconv.Atoi(operand)
		}

		switch operator {
		case "+":
			return item + v
		case "*":
			return item * v
		default:
			panic("unknown operator")
		}
	}

	scanner.Scan()
	var v int
	fmt.Sscanf(scanner.Text(), "  Test: divisible by %d", &v)
	m.test = func(v int) func(int) bool {
		return func(item int) bool {
			return item%v == 0
		}
	}(v)

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "    If true: throw to monkey %d", &v)
	m.trueMonkey = v

	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "    If false: throw to monkey %d", &v)
	m.falseMonkey = v

	return &m
}

func part1(monkeys []*monkey) {
	inspect := make([]int, len(monkeys))

	for round := 0; round < 20; round++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				// Inspect item
				inspect[i]++
				item = m.op(item)

				// Relief
				item /= 3

				// Throw item
				if m.test(item) {
					monkeys[m.trueMonkey].items = append(monkeys[m.trueMonkey].items, item)
				} else {
					monkeys[m.falseMonkey].items = append(monkeys[m.falseMonkey].items, item)
				}
			}
			m.items = nil
		}
	}

	sort.Ints(inspect)
	fmt.Println(inspect)
	fmt.Println(inspect[len(inspect)-1] * inspect[len(inspect)-2])
}

func part2(monkeys []*monkey) {
	inspect := make([]int, len(monkeys))

	for round := 0; round < 10000; round++ {
		for i, m := range monkeys {
			for _, item := range m.items {
				// Inspect item
				inspect[i]++
				item = m.op(item)

				// Relief
				// Use LCM since all the divisor tests are prime
				item %= 9699690

				// Throw item
				if m.test(item) {
					monkeys[m.trueMonkey].items = append(monkeys[m.trueMonkey].items, item)
				} else {
					monkeys[m.falseMonkey].items = append(monkeys[m.falseMonkey].items, item)
				}
			}
			m.items = nil
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspect)))
	fmt.Println(inspect)
	fmt.Println(inspect[0] * inspect[1])
}
