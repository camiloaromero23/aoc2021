package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/camiloaromero23/aoc2021/utils"
)

func mostCommon(a, b int) (mostCommon int, leastCommon int) {
	if a > b {
		return 0, 1
	}
	return 1, 0
}

func countBits(arr *[]string, bitIndex int) (zeros, ones int) {
	bitsList := *arr
	for j := range bitsList {
		b := bitsList[j][bitIndex]
		if string(b) == "0" {
			zeros++
		} else {
			ones++
		}
	}
	return
}

func getRating(arr *[]string, bitIndex int, bitVal int) (rating []string) {
	bitList := *arr
	rating = []string{}
	for j, v := range bitList {
		b := bitList[j][bitIndex]
		if string(b) == fmt.Sprint(bitVal) {
			rating = append(rating, v)
		}
	}
	return
}

func updateRating(rating *[]string, ratingType string, index int) {
	if len(*rating) == 1 {
		return
	}

	zeros, ones := countBits(rating, index)
	most, least := mostCommon(zeros, ones)
	if ratingType == "oxygen" {
		*rating = getRating(rating, index, most)
	} else {
		*rating = getRating(rating, index, least)
	}
}

func main() {
	in, err := os.ReadFile("day3.input")
	utils.CheckErr(err)
	input := strings.TrimSpace(string(in))
	aux := strings.Split(input, "\n")

	var oxygen, co2 []string

	zeros, ones := countBits(&aux, 0)
	most, least := mostCommon(zeros, ones)
	oxygen = getRating(&aux, 0, most)
	co2 = getRating(&aux, 0, least)

	bitSize := len(aux[0])
	for i := 1; i < bitSize; i++ {
		updateRating(&oxygen, "oxygen", i)
		updateRating(&co2, "co2", i)
	}

	oxygenRating, err := strconv.ParseInt(oxygen[0], 2, 64)
	utils.CheckErr(err)
	co2Rating, err := strconv.ParseInt(co2[0], 2, 64)
	utils.CheckErr(err)

	fmt.Printf("Power consumption = %d", oxygenRating*co2Rating)
}
