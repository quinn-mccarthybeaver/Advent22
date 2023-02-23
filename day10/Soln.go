package day10

import (
	"bufio"
	"fmt"
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

		signals = append(signals, signals[len(signals)-1])
		if cmd[0] == "addx" {
			delta, _ := strconv.Atoi(cmd[1])
			signals = append(signals, signals[len(signals)-1]+delta)
		}
	}

	var res1 int
	for i := 20; i < len(signals); i += 40 {
		res1 += i * signals[i]
	}

	drawScreen(signals)

	return res1, 0
}

func drawScreen(signals []int) {
	line := ""
	for i := 0; i < 240; i++ {
		if i%40 == 0 {
			fmt.Println(line)
			line = ""
		}

		if signals[i]-(i%40) >= -1 && signals[i]-(i%40) <= 1 {
			line = line + "#"
		} else {
			line = line + "."
		}
	}

	fmt.Println(line)
}