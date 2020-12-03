package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const inputFile string = "./assets/input/day02.txt"

func main() {
	input, _ := ioutil.ReadFile(inputFile)
	part1, part2 := countValidPasswords(input)

	fmt.Printf("Part 1 solution: %v\n", part1)
	fmt.Printf("Part 2 solution: %v\n", part2)
}

func countValidPasswords(input []byte) (solutionPart1 int, solutionPart2 int) {
	var part1, part2 int = 0, 0

	for _, split := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var min, max int
		var char byte
		var pass string

		fmt.Sscanf(split, "%v-%v %c: %v", &min, &max, &char, &pass)
		count := strings.Count(pass, string(char))

		if count >= min && count <= max {
			part1++
		}
		if (pass[min-1] == char) != (pass[max-1] == char) {
			part2++
		}
	}

	return part1, part2
}
