package day10

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Soln(file *os.File) (int, int) {
	scanner := bufio.NewScanner(file)

	var signals []int
	signals = append(signals, 1)
	for scanner.Scan() {
		cmd := strings.Fields(scanner.Text())

		switch cmd[0] {
		case "noop":
			signals = append(signals, signals[len(signals)-1])
		case "addx":
			signals = append(signals, signals[len(signals)-1])
			delta, _ := strconv.Atoi(cmd[1])
			signals = append(signals, signals[len(signals)-1]+delta)
		}
	}

	var res1 int
	for i := 20; i < len(signals); i += 40 {
		res1 += i * signals[i]
	}

	return res1, 0
}