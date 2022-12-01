package main

import (
	"bufio"
	"fmt"
	"os"
)

func input_reader() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		fmt.Println("input: day number to run")
	}

	return scanner
}