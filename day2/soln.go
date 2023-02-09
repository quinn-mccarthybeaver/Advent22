package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Soln1(file *os.File) int {
	reader := bufio.NewScanner(file)
	var lefts []string
	var rights []string
	for reader.Scan() {
		in := strings.Fields(reader.Text())
		lefts = append(lefts, in[0])
		rights = append(rights, in[1])
	}
	var score int
	for idx, val := range lefts {
		victory := 3
		switch rights[idx] {
		case "X":
			score += 1
			if val == "B" {
				victory -= 3
			} else if val == "C" {
				victory += 3
			}
		case "Y":
			score += 2
			if val == "C" {
				victory -= 3
			} else if val == "A" {
				victory += 3
			}
		case "Z":
			score += 3
			if val == "A" {
				victory -= 3
			} else if val == "B" {
				victory += 3
			}
		}
		score += victory
	}
	fmt.Println(score)
	return score
}

func Soln2(file *os.File) int {
	reader := bufio.NewScanner(file)
	var lefts []string
	var rights []string
	for reader.Scan() {
		in := strings.Fields(reader.Text())
		lefts = append(lefts, in[0])
		rights = append(rights, in[1])
	}
	var score int
	for idx, val := range lefts {
		switch rights[idx] {
		case "X":
			score += lose_score(val)
		case "Y":
			score += 3 + tie_score(val)
		case "Z":
			score += 6 + win_score(val)
		}
	}
	return score
}

// A B C -> Rock Paper Scisors
// 1 2 3
// win: 6
// tie: 3
// los; 0
func lose_score(x string) int {
	switch x {
	case "A":
		return 3
	case "B":
		return 1
	case "C":
		return 2
	default:
		return 0
	}
}

func win_score(x string) int {
	switch x {
	case "A":
		return 2
	case "B":
		return 3
	case "C":
		return 1
	default:
		return 0
	}
}

func tie_score(x string) int {
	switch x {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	default:
		return 0
	}
}