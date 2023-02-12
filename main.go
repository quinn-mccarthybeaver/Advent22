package main

import (
	"22advent/day1"
	"22advent/day2"
	"22advent/day3"
	"22advent/day4"
	"22advent/day5"
	"22advent/day6"
	"22advent/day7"
	"22advent/day8"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("give a single number as an arg")
		return
	}

	var day int
	var err error
	day, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	path := fmt.Sprintf("day%d/input", day)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("unable to open file with error:", err)
		os.Exit(1)
	}

	var ans1 int
	var ans2 int

	switch day {
	case 1:
		fmt.Println(day1.Soln(file))
	case 2:
		fmt.Println(day2.Soln2(file))
	case 3:
		fmt.Println(day3.Soln2(file))
	case 4:
		fmt.Println(day4.Soln2(file))
	case 5:
		fmt.Println(day5.Soln2(file))
	case 6:
		fmt.Println(day6.Soln1(file))
	case 7:
		ans1, ans2 = day7.Soln(file)
	case 8:
		ans1, ans2 = day8.Soln(file)
	default:
		fmt.Println("not implemented yet")
	}
	fmt.Println(ans1)
	fmt.Println(ans2)
}