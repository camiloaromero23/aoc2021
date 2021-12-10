package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/camiloaromero23/aoc2021/utils"
)

var s = slot{0, false}
var winners []int

type slot struct {
	value  int
	marked bool
}

type board *[5][5]slot

func initBoards(boardsInput []string) *[]board {
	var err error
	boards := []board{}
	for i, l := range boardsInput {
		l = strings.TrimSpace(l)
		l = utils.RemoveDuplicateWhitespace(l)
		row := strings.Split(l, " ")
		if i%6 == 0 {
			// Create new board
			b := &[5][5]slot{{s, s, s, s, s}, {s, s, s, s, s}, {s, s, s, s, s}, {s, s, s, s, s}, {s, s, s, s, s}}
			boards = append(boards, b)
		}
		boardIndex := i / 6
		for j, v := range row {
			if i%6 == 0 {
				continue
			}
			boards[boardIndex][(i-boardIndex-1)%5][j].value, err = strconv.Atoi(v)
			utils.CheckErr(err)
		}
	}
	return &boards

}

func checkDraw(draw int, boards *[]board) (int, bool) {
	for i, b := range *boards {
		if checkBoard(draw, &b) && !contains(winners, i) {
			winners = append(winners, i)
		}
	}

	if len(winners) == len(*boards) {
		return winners[len(winners)-1], true
	}
	return 0, false
}

func checkBoard(draw int, b *board) bool {
	for i, r := range *b {
		for j, s := range r {
			if s.value == draw {
				board := *b
				board[i][j].marked = true
				return checkWinnerBoard(b, i, j)
			}
		}
	}
	return false
}

func checkWinnerBoard(b *board, r, c int) bool {
	for i := range *b {
		board := *b
		if !board[r][i].marked {
			break
		}
		if i == len(board)-1 {
			return true
		}
	}
	for i := range *b {
		board := *b
		if !board[i][c].marked {
			break
		}
		if i == len(board)-1 {
			return true
		}
	}
	return false
}

func checkWinnerScore(boards []board, wbIndex, winnerNumber int) (score int) {
	wb := boards[wbIndex]
	for _, r := range wb {
		for _, s := range r {
			if !s.marked {
				score += s.value
			}
		}
	}
	score *= winnerNumber
	return
}

func main() {
	in, err := os.ReadFile("day4.input")
	utils.CheckErr(err)
	input := strings.TrimSpace(string(in))
	lines := strings.Split(input, "\n")
	stringDraws := strings.Split(lines[0], ",")
	var draws []int
	for _, draw := range stringDraws {
		intDraw, err := strconv.Atoi(draw)
		utils.CheckErr(err)
		draws = append(draws, intDraw)
	}
	boardsInput := lines[1:]

	boards := initBoards(boardsInput)

	var winnerIndex int
	var winner bool
	var winnerNumber int
	for _, draw := range draws {
		if winnerIndex, winner = checkDraw(draw, boards); winner {
			winnerNumber = draw
			break
		}
	}

	score := checkWinnerScore(*boards, winnerIndex, winnerNumber)

	fmt.Printf("Winner score: %d\n", score)

}

// Snippet taken from Mostafa - stack overflow
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
