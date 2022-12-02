package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/samber/lo"
)

func main() {
	elves := make([]int, 0)

	var sum int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			elves = append(elves, sum)
			sum = 0
		}
		sum += n
	}

	fmt.Println(lo.Max(elves))

	sort.Ints(elves)
	fmt.Println(elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3])
}
