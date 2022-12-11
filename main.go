package main

import (
	"22advent/day1"
	"22advent/day2"
	"22advent/day3"
	"22advent/day4"
	"22advent/day5"
	"22advent/day6"
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
		fmt.Println(day1.Soln(scanner))
	case 2:
		fmt.Println(day2.Soln2(scanner))
	case 3:
		fmt.Println(day3.Soln2(scanner))
	case 4:
		fmt.Println(day4.Soln2(scanner))
	case 5:
		fmt.Println(day5.Soln2(scanner))
	case 6:
		fmt.Println(day6.Soln1(scanner))
	default:
		fmt.Println("not implemented yet")
	}
}