package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/camiloaromero23/aoc2021/utils"
)

type fissure struct {
	from coord
	to   coord
}

type coord struct {
	x, y int
}

func main() {
	in, err := os.ReadFile("day5.input")
	utils.CheckErr(err)
	input := strings.TrimSpace(string(in))
	lines := strings.Split(input, "\n")
	fissures := []fissure{}
	for _, v := range lines {
		fissureStr := strings.Split(v, " -> ")
		from := strings.Split(fissureStr[0], ",")
		to := strings.Split(fissureStr[1], ",")
		x1, err := strconv.Atoi(from[0])
		utils.CheckErr(err)
		y1, err := strconv.Atoi(from[1])
		utils.CheckErr(err)
		x2, err := strconv.Atoi(to[0])
		utils.CheckErr(err)
		y2, err := strconv.Atoi(to[1])
		utils.CheckErr(err)
		f := fissure{
			from: coord{x1, y1},
			to:   coord{x2, y2},
		}
		fissures = append(fissures, f)
	}

	xMax, yMax := func() (int, int) {
		var x, y int
		for _, f := range fissures {
			if f.from.x > x {
				x = f.from.x
			}
			if f.from.y > y {
				y = f.from.y
			}
		}
		return x + 1, y + 1
	}()
	board := make([][]int, xMax)
	for i := range board {
		board[i] = make([]int, yMax)
	}
	for _, f := range fissures {
		addFissure(&board, f)
	}

	points := countPoints(&board)
	fmt.Println("Total points:", points)
}

func addFissure(board *[][]int, f fissure) {
	if (f.from.x == f.from.y) && (f.to.x == f.to.y) {
		if f.from.x > f.to.x {
			increaseCoordDiagonally(board, "DIAGONAL_EVEN", f)
		} else {
			increaseCoordDiagonally(board, "DIAGONAL_EVEN", f)
		}
		return
	}
	if is45Deg(f) {
		if f.from.x > f.to.x {
			increaseCoordDiagonally(board, "DIAGONAL_UNEVEN", f)
		} else {
			increaseCoordDiagonally(board, "DIAGONAL_UNEVEN", f)
		}
		return
	}
	if (f.from.x != f.to.x) && (f.from.y != f.to.y) {
		return
	}
	if f.from.x != f.to.x {
		if f.from.x > f.to.x {
			increaseCoord(board, "VERTICAL", f.to.x, f.from.x, f.from.y)
			return
		}
		increaseCoord(board, "VERTICAL", f.from.x, f.to.x, f.from.y)
		return
	}

	if f.from.y > f.to.y {
		increaseCoord(board, "HORIZONTAL", f.to.y, f.from.y, f.from.x)
		return
	}
	increaseCoord(board, "HORIZONTAL", f.from.y, f.to.y, f.from.x)
	return
}

func increaseCoord(board *[][]int, direction string, from, to, basePos int) {
	b := *board
	for i := from; i <= to; i++ {
		switch direction {
		case "HORIZONTAL":
			b[basePos][i]++
		case "VERTICAL":
			b[i][basePos]++
		}
	}
}

func is45Deg(f fissure) bool {
	y := math.Abs(float64(f.to.y - f.from.y))
	x := math.Abs(float64(f.to.x - f.from.x))
	return x == y
}

func increaseCoordDiagonally(board *[][]int, direction string, f fissure) {
	b := *board
	var x, y, grx, gry int
	if f.to.x > f.from.x {
		x = f.from.x
		y = f.from.y
		grx = f.to.x
		gry = f.to.y
	} else {
		x = f.to.x
		y = f.to.y
		grx = f.from.x
		gry = f.from.y
	}

	for x <= grx {
		switch direction {
		case "DIAGONAL_UNEVEN":
			b[x][y]++
			x++
			if gry > y {
				y++
			} else {
				y--
			}
		case "DIAGONAL_EVEN":
			b[x][x]++
			x++
		}
	}
}

func countPoints(board *[][]int) (score int) {
	for _, r := range *board {
		for _, v := range r {
			if v > 1 {
				score++
			}
		}
	}
	return
}
