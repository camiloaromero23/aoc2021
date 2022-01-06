package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/camiloaromero23/aoc2021/utils"
)

var numbers = map[string]string{
	"cagedb":  "0",
	"ab":      "1",
	"gcdfa":   "2",
	"fbcad":   "3",
	"eafb":    "4",
	"cdfbe":   "5",
	"cdfgeb":  "6",
	"dab":     "7",
	"acedgfb": "8",
	"cefabd":  "9",
}

func main() {
	buf, err := os.ReadFile("day8.testinput")
	// buf, err := os.ReadFile("day8.input")
	utils.CheckErr(err)
	in := strings.TrimSpace(string(buf))
	lines := strings.Split(in, "\n")

	// var count int
	var sum int
	for k, v := range numbers {
		delete(numbers, k)
		key := utils.SortStringByCharacter(k)
		numbers[key] = v
	}
	fmt.Printf("%+v\n", numbers)
	// for i, v := range lines {
	for _, v := range lines {
		// if i > 1 {
		// 	break
		// }
		var num string
		l := strings.Split(v, " | ")
		input := l[0]
		output := l[1]
		in := strings.Split(input, " ")
		out := strings.Split(output, " ")
		in = filterByLengths(in, 2, 3, 4, 7)
		for i := range in {
			in[i] = utils.SortStringByCharacter(in[i])
			out[i] = utils.SortStringByCharacter(out[i])
		}
		// fmt.Printf("%v | %v\n", in, out)

		for _, ow := range out {
			// for _, iw := range in {
			// if iw == ow {
			// count++
			// fmt.Println("*******************************************************")
			// fmt.Println("ow", ow)
			// fmt.Println("nums[ow]", numbers[ow])
			n, ok := numbers[ow]
			if ok {
				num += n
			} else {
				if len(ow) == 2 {
					num += "1"
				}
				if len(ow) == 3 {
					num += "7"
				}
				if len(ow) == 4 {
					num += "4"
				}
				if len(ow) == 7 {
					num += "8"
				}
			}
			// }
			// }
		}
		fmt.Printf("%s = %s\n", out, num)
		s, err := strconv.Atoi(num)
		utils.CheckErr(err)
		sum += s
	}

	// fmt.Println("Current count:", count)
	fmt.Println("Sum", sum)
}

func filterByLengths(slice []string, lengths ...int) []string {
	filtered := []string{}
	for _, length := range lengths {
		for _, n := range slice {
			if len(n) == length {
				filtered = append(filtered, n)
			}
		}

	}
	return filtered
}
