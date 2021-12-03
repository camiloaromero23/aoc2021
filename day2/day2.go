package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/camiloaromero23/aoc2021/utils"
)

var horizontal int
var depth int
var aim int

var movements = map[string]func(int){
	"down": func(amount int) {
		aim += amount
	},
	"up": func(amount int) {
		aim -= amount
	},
	"forward": func(amount int) {
		horizontal += amount
		depth += aim * amount
	},
}

func main() {
	in, err := os.ReadFile("day2.input")
	utils.CheckErr(err)
	input := strings.TrimSpace(string(in))

	aux := strings.Split(input, "\n")
	for _, move := range aux {
		moveAux := strings.Split(move, " ")
		d := moveAux[0]
		u, err := strconv.Atoi(moveAux[1])
		utils.CheckErr(err)
		movements[d](u)
	}
	fmt.Printf("Horizantal position * depth = %d\n", horizontal*depth)

}
