package day7

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type dir struct {
	files   map[string]int
	subdirs map[string]*dir
	parent  *dir
	size    int
}

func Soln1(file *os.File) (int, int) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()

	// this makes a dir holding all zero valued fields
	cur_dir := &dir{
		files:   make(map[string]int),
		subdirs: make(map[string]*dir),
	}

	toplevel := cur_dir

	// This should probably be extracted
	for scanner.Scan() {
		tokens := strings.Fields(scanner.Text())

		is_cmd, _ := regexp.Match("^\\$$", []byte(tokens[0]))

		if is_cmd {

			// change directory
			if tokens[1] == "cd" {
				if tokens[2] == ".." {
					cur_dir = cur_dir.parent
				} else {
					new_dir := &dir{
						files:   make(map[string]int),
						subdirs: make(map[string]*dir),
						parent:  cur_dir,
						size:    0,
					}

					cur_dir.subdirs[tokens[2]] = new_dir
					cur_dir = new_dir
				}
			}

			// TODO: think about if I actually care about ls
			// probably only used in part 2
		}

		is_file, _ := regexp.Match("^\\d+$", []byte(tokens[0]))

		// TODO: possible bug if ls is called on a directory more than once
		if is_file {
			fsize, _ := strconv.Atoi(tokens[0])
			cur_dir.size += fsize
		}
	}

	res := toplevel.getAnswer()
	res2 := toplevel.getAnswer2()

	// 70000000
	// 30000000

	return res, res2
}

func (self *dir) countDirSize() int {
	res := self.size

	for _, val := range self.subdirs {
		res += val.countDirSize()
	}

	return res
}

func (self *dir) getAnswer() int {
	var acc int
	curDirSize := self.countDirSize()

	if curDirSize <= 100000 {
		acc += curDirSize
	}

	for _, val := range self.subdirs {
		acc += val.getAnswer()
	}

	return acc
}

func (self *dir) getAnswer2() int {
	candidates := make(map[string]int)

	candidates["/"] = self.countDirSize()
	diff := candidates["/"] - 40_000_000

	for key, val := range self.subdirs {
		score := val.helper(&candidates, diff)

		if score >= diff {
			candidates[key] = score
		}
	}

	score := candidates["/"]
	for _, val := range candidates {
		if val < score {
			score = val
		}
	}

	return score
}

func (self *dir) helper(candidates *map[string]int, diff int) int {
	res := self.countDirSize()

	for key, val := range self.subdirs {
		var score int
		score = val.helper(candidates, diff)

		if score >= diff {
			(*candidates)[key] = score
		}
	}
	return res
}