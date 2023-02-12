package day9

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	row int
	col int
}

func Soln(file *os.File) (int, int) {
	scanner := bufio.NewScanner(file)

	trail1 := make(map[coord]bool)
	trail2 := make(map[coord]bool)
	trail1[coord{0, 0}] = true
	trail2[coord{0, 0}] = true

	rope := make([]coord, 10)

	for scanner.Scan() {
		action := strings.Fields(scanner.Text())

		count, _ := strconv.Atoi(action[1])

		for i := 0; i < count; i++ {
			switch action[0] {
			case "U":
				rope[0].row += 1
			case "D":
				rope[0].row -= 1
			case "L":
				rope[0].col -= 1
			case "R":
				rope[0].col += 1
			}

			for i := 0; i < len(rope)-1; i++ {
				updateTail(&rope[i], &rope[i+1])
			}
			trail1[rope[1]] = true
			trail2[rope[9]] = true
		}
	}

	res1 := len(trail1)
	res2 := len(trail2)
	return res1, res2
}

func updateTail(head *coord, tail *coord) {
	var rowDelta int
	var colDelta int
	if head.row > tail.row+1 || (diff(head.col, tail.col) > 1 && head.row > tail.row) {
		rowDelta++
	} else if head.row < tail.row-1 || (diff(head.col, tail.col) > 1 && head.row < tail.row) {
		rowDelta--
	}

	if head.col > tail.col+1 || (diff(head.row, tail.row) > 1 && head.col > tail.col) {
		colDelta++
	} else if head.col < tail.col-1 || (diff(head.row, tail.row) > 1 && head.col < tail.col) {
		colDelta--
	}
	tail.row += rowDelta
	tail.col += colDelta
}

func diff(x int, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}