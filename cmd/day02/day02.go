package main

import (
	"fmt"
	"io/ioutil"

	"github.com/fosdickio/advent-of-code/pkg/utils"
)

const inputFile string = "./assets/input/day02.txt"

func main() {
	input, _ := ioutil.ReadFile(inputFile)
	part1, part2 := utils.CountValidPasswords(input)

	fmt.Printf("Part 1 solution: %v\n", part1)
	fmt.Printf("Part 2 solution: %v\n", part2)
}
