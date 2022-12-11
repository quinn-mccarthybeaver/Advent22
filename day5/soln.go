package day5

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Soln1(scanner *bufio.Scanner) int {
	crates := process_crates(scanner)
	process_moves(&crates, scanner)
	for _, val := range crates {
		fmt.Print(string(val[0]))
	}
	fmt.Println()
	return 0
}

func Soln2(scanner *bufio.Scanner) int {
	crates := process_crates(scanner)
	crates = process_moves1(crates, scanner)
	for _, val := range crates {
		fmt.Print(string(val[0]))
	}
	fmt.Println()
	return 0
}

func process_crates(scanner *bufio.Scanner) [][]rune {
	res := make([][]rune, 9)
	scanner.Scan()
	line := scanner.Text()

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		var count int
		for idx, val := range line {
			if (idx-1)%4 == 0 {
				if val != ' ' {
					res[count] = append(res[count], val)
				}
				count++
			}
		}

		line = scanner.Text()
	}
	return res
}

func process_moves(crates *[][]rune, scanner *bufio.Scanner) {
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		quantity, _ := strconv.Atoi(line[1])
		from, _ := strconv.Atoi(line[3])
		to, _ := strconv.Atoi(line[5])

		from--
		to--

		for idx, val := range (*crates)[from] {
			if idx < quantity {
				(*crates)[to] = append([]rune{val}, (*crates)[to]...)
			} else {
				break
			}
		}
		(*crates)[from] = (*crates)[from][quantity:]
	}
}

func process_moves1(crates [][]rune, scanner *bufio.Scanner) [][]rune {
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		quantity, _ := strconv.Atoi(line[1])
		from, _ := strconv.Atoi(line[3])
		to, _ := strconv.Atoi(line[5])

		from--
		to--

		moved := make([]rune, quantity)
		copy(moved, crates[from][0:quantity])
		crates[to] = append(moved, crates[to]...)
		crates[from] = crates[from][quantity:]
	}

	return crates
}
