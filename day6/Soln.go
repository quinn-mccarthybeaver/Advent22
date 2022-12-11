package day6

import (
	"bufio"
)

func Soln1(scanner *bufio.Scanner) int {
	scanner.Scan()
	line := scanner.Text()
	var buffer []rune

	var count int
	for idx, val := range line {
		if idx < 14 {
			buffer = append(buffer, val)
			continue
		}

		if all_diff(buffer) {
			count = idx
			break
		} else {
			for idx := range buffer {
				if idx == len(buffer)-1 {
					break
				} else {
					buffer[idx] = buffer[idx+1]
				}
			}
			buffer[len(buffer)-1] = val
		}

	}
	return count
}

func all_diff(buf []rune) bool {
	for idx, val := range buf {
		for _, other := range buf[idx+1:] {
			if val == other {
				return false
			}
		}
	}

	return true
}