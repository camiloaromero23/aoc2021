package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/camiloaromero23/aoc2021/utils"
)

func main() {
	// buf, err := os.ReadFile("day7.testinput")
	buf, err := os.ReadFile("day7.input")
	utils.CheckErr(err)
	input := strings.TrimSpace(string(buf))
	line := strings.Split(input, ",")
	crabs := []int{}

	for _, v := range line {
		c, err := strconv.Atoi(v)
		utils.CheckErr(err)
		crabs = append(crabs, c)
	}
	min, max := func() (min int, max int) {
		for _, c := range crabs {
			if c < min {
				min = c
			}
			if c > max {
				max = c
			}
		}
		return
	}()

	fuel := func() (f int) {
		f = math.MaxInt
		distances := []int{}
		for i := min; i < max; i++ {
			var distance int
			for _, c := range crabs {
				d := int(math.Abs((float64(c - i))))
				distance += d * (d + 1) / 2
			}
			distances = append(distances, distance)
		}
		for _, d := range distances {
			f = int(math.Min(float64(f), float64(d)))
		}
		return
	}()
	fmt.Println("min fuel:", fuel)
}
