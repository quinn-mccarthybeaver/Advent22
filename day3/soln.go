package day3

import (
	"bufio"
	"fmt"
	"strings"
)

func Soln1(scanner *bufio.Scanner) int {
	var satchels []string
	for scanner.Scan() {
		satchels = append(satchels, scanner.Text())
	}

	var sum int
	scores := map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26, 'A': 27, 'B': 28, 'C': 29, 'D': 30, 'E': 31, 'F': 32, 'G': 33, 'H': 34, 'I': 35, 'J': 36, 'K': 37, 'L': 38, 'M': 39, 'N': 40, 'O': 41, 'P': 42, 'Q': 43, 'R': 44, 'S': 45, 'T': 46, 'U': 47, 'V': 48, 'W': 49, 'X': 50, 'Y': 51, 'Z': 52}
	for _, val := range satchels {
		sum += scores[satchel_prio(val)]
	}

	return sum
}

func Soln2(scanner *bufio.Scanner) int {
	var sum int
	var place int
	satchels := make([]string, 3)
	scores := map[rune]int{'a': 1, 'b': 2, 'c': 3, 'd': 4, 'e': 5, 'f': 6, 'g': 7, 'h': 8, 'i': 9, 'j': 10, 'k': 11, 'l': 12, 'm': 13, 'n': 14, 'o': 15, 'p': 16, 'q': 17, 'r': 18, 's': 19, 't': 20, 'u': 21, 'v': 22, 'w': 23, 'x': 24, 'y': 25, 'z': 26, 'A': 27, 'B': 28, 'C': 29, 'D': 30, 'E': 31, 'F': 32, 'G': 33, 'H': 34, 'I': 35, 'J': 36, 'K': 37, 'L': 38, 'M': 39, 'N': 40, 'O': 41, 'P': 42, 'Q': 43, 'R': 44, 'S': 45, 'T': 46, 'U': 47, 'V': 48, 'W': 49, 'X': 50, 'Y': 51, 'Z': 52}

	for scanner.Scan() {
		satchels[place] = scanner.Text()
		place = (place + 1) % 3

		if place == 0 {
			sum += scores[group_prio(satchels)]
		}
	}

	return sum
}

func satchel_prio(x string) rune {
	left := x[0 : len(x)/2]
	right := x[len(x)/2:]

	for _, val := range left {
		if strings.ContainsRune(right, val) {
			return val
		}
	}
	fmt.Println("Error")
	return 0
}

func group_prio(x []string) rune {
	var acc []rune
	for _, val := range x[0] {
		if strings.ContainsRune(x[1], val) {
			acc = append(acc, val)
		}
	}

	for _, val := range acc {
		if strings.ContainsRune(x[2], val) {
			return val
		}
	}
	fmt.Println("Error")
	return 0
}
