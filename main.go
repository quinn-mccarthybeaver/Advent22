package main

import (
	"22advent/day1"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	day1.Soln1(scanner)
}