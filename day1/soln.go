package day1

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type gnome struct {
	food   []int
	calsum int
}

func (self *gnome) add(food int) {
	self.calsum += food
	self.food = append(self.food, food)
}

func Soln(file *os.File) int {
	input := bufio.NewScanner(file)
	var gnomes []gnome
	new_gnome := &(gnome{[]int{}, 0})

	for input.Scan() {
		if input.Text() == "" {
			gnomes = append(gnomes, *new_gnome)
			new_gnome = &(gnome{[]int{}, 0})
		} else {
			new_food, _ := strconv.Atoi(input.Text())
			new_gnome.add(new_food)
		}
	}

	sort.Slice(gnomes, func(i, j int) bool { return gnomes[i].calsum > gnomes[j].calsum })
	return gnomes[0].calsum + gnomes[1].calsum + gnomes[2].calsum
}