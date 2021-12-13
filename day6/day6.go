package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/camiloaromero23/aoc2021/utils"
)

func main() {
	// buf, err := os.ReadFile("day6.testinput")
	buf, err := os.ReadFile("day6.input")
	utils.CheckErr(err)
	input := strings.TrimSpace(string(buf))
	line := strings.Split(input, ",")
	// daysToWatch := 18
	// daysToWatch := 80
	daysToWatch := 256

	fishes := map[int]int{}

	for _, lf := range line {

		days, err := strconv.Atoi(lf)
		utils.CheckErr(err)
		fishes[days] += 1
	}

	for i := 0; i < daysToWatch; i++ {
		n := fishes[0]
		fishes[0] = 0
		for fd := 1; fd < 9; fd++ {
			fishes[fd-1] += fishes[fd]
			fishes[fd] = 0
		}
		fishes[6] += n
		fishes[8] = n
	}

	fishCount := func() (fc int) {
		for _, f := range fishes {
			fc += f
		}
		return
	}()

	fmt.Printf("fish amount = %d\n", fishCount)
}
