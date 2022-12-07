package main

import (
	"22advent/day1"
	"22advent/day2"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if len(os.Args) != 2 {
		fmt.Println("give a single number as an arg")
		return
	}

	var day int
	var err error
	day, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
	}

	switch day {
	case 1:
		day1.Soln1(scanner)
	case 2:
		day2.Soln2(scanner)
	default:
		fmt.Println("not implemented yet")
	}
}