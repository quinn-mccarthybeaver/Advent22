package day4

import (
	"bufio"
	// "fmt"
	"strconv"
	"strings"
)

func Soln1(scanner *bufio.Scanner) int {
	var sum int
	for scanner.Scan() {
		f := strings.Split(scanner.Text(), ",")
		left := strings.Split(f[0], "-")
		right := strings.Split(f[1], "-")

		if contains(left, right) {
			sum++
		}
	}
	return sum
}

func Soln2(scanner *bufio.Scanner) int {
	var sum int
	for scanner.Scan() {
		f := strings.Split(scanner.Text(), ",")
		left := strings.Split(f[0], "-")
		right := strings.Split(f[1], "-")

		if contains1(left, right) {
			sum++
		}
	}
	return sum
}

func contains(left []string, right []string) bool {
	l0, _ := strconv.Atoi(left[0])
	l1, _ := strconv.Atoi(left[1])
	r0, _ := strconv.Atoi(right[0])
	r1, _ := strconv.Atoi(right[1])

	// fmt.Println(l0, l1, r0, r1)

	return (l0 <= r0 && l1 >= r1) || (r0 <= l0 && r1 >= l1)
}

func contains1(left []string, right []string) bool {
	l0, _ := strconv.Atoi(left[0])
	l1, _ := strconv.Atoi(left[1])
	r0, _ := strconv.Atoi(right[0])
	r1, _ := strconv.Atoi(right[1])

	// fmt.Println(l0, l1, r0, r1)

	return (r0 <= l0 && r1 >= l0) || (r0 <= l1 && r1 >= l1) || (l0 <= r0 && l1 >= r0) || (l0 <= r1 && l1 >= r1)
}