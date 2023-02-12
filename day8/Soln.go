package day8

import (
	"bufio"
	"os"
	"strconv"
)

func Soln(file *os.File) (int, int) {
	scanner := bufio.NewScanner(file)

	var grid [][]int
	for scanner.Scan() {
		var newline []int
		for _, val := range scanner.Text() {
			digit, _ := strconv.Atoi(string(val))
			newline = append(newline, digit)
		}
		grid = append(grid, newline)
	}

	res1 := visibleCoords(&grid)
	res2 := bestTree(&grid)
	return res1, res2
}

func visibleCoords(trees *[][]int) int {
	result := make(map[struct {
		row int
		col int
	}]string)
	// up
	for col := 0; col < len(*trees); col++ {
		largest := -1
		for row := len(*trees) - 1; row >= 0; row-- {
			if (*trees)[row][col] > largest {
				largest = (*trees)[row][col]
				result[struct {
					row int
					col int
				}{row: row, col: col}] = "bottom"
			}
		}
	}

	// down
	for col := 0; col < len(*trees); col++ {
		largest := -1
		for row := 0; row < len(*trees); row++ {
			if (*trees)[row][col] > largest {
				largest = (*trees)[row][col]
				result[struct {
					row int
					col int
				}{row: row, col: col}] = "top"
			}
		}
	}
	// left
	for row := 0; row < len(*trees); row++ {
		largest := -1
		for col := len(*trees) - 1; col >= 0; col-- {
			if (*trees)[row][col] > largest {
				largest = (*trees)[row][col]
				result[struct {
					row int
					col int
				}{row: row, col: col}] = "right"
			}
		}
	}
	// right
	for row := 0; row < len(*trees); row++ {
		largest := -1
		for col := 0; col < len(*trees); col++ {
			if (*trees)[row][col] > largest {
				largest = (*trees)[row][col]
				result[struct {
					row int
					col int
				}{row: row, col: col}] = "left"
			}
		}
	}

	return len(result)
}

func scoreTree(row int, col int, trees *[][]int) int {
	thisTree := (*trees)[row][col]
	res := 1
	count := 0
	// up
	for i := row - 1; i >= 0; i-- {
		if (*trees)[i][col] < thisTree {
			count++
		} else {
			count++
			break
		}
	}
	res *= count
	count = 0

	// down
	for i := row + 1; i < len(*trees); i++ {
		if (*trees)[i][col] < thisTree {
			count++
		} else {
			count++
			break
		}
	}
	res *= count
	count = 0

	// left
	for i := col - 1; i >= 0; i-- {
		if (*trees)[row][i] < thisTree {
			count++
		} else {
			count++
			break
		}
	}
	res *= count
	count = 0

	// right
	for i := col + 1; i < len(*trees); i++ {
		if (*trees)[row][i] < thisTree {
			count++
		} else {
			count++
			break
		}
	}
	res *= count

	return res
}

func bestTree(trees *[][]int) int {
	largestScore := 0
	for row := 0; row < len(*trees); row++ {
		for col := 0; col < len((*trees)[0]); col++ {
			newScore := scoreTree(row, col, trees)
			if newScore > largestScore {
				largestScore = newScore
			}
		}
	}

	return largestScore
}