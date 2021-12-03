package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/camiloaromero23/aoc2021/utils"
)

func isLarger(a, b [3]int) bool {
	var accA, accB int
	for i := 0; i < len(a); i++ {
		accA += a[i]
		accB += b[i]
	}
	return accB > accA
}

func main() {
	in, err := os.ReadFile("day1.input")
	utils.CheckErr(err)
	input := strings.TrimSpace(string(in))

	aux := strings.Split(input, "\n")
	depths := make([]int, len(aux))

	for i, v := range aux {
		val, err := strconv.Atoi(v)
		utils.CheckErr(err)
		depths[i] = val
	}

	prev := [3]int{depths[0], depths[1], depths[2]}
	curr := [3]int{depths[1], depths[2], depths[3]}
	var count int

	for i, d := range depths {
		if i < 3 {
			continue
		}

		if isLarger(prev, curr) {
			count++
		}

		prev[i%3] = curr[(i-1)%3]
		curr[i%3] = d
	}

	if isLarger(prev, curr) {
		count++
	}

	fmt.Printf("Measurement increases = %d\n", count)
}
