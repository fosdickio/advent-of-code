package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/fosdickio/advent-of-code/pkg/utils"
)

const inputFile string = "./assets/input/day03.txt"

func main() {
	grid, err := utils.ReadInputFileIntoGrid(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	slopes := map[image.Point]int{
		{1, 1}: 0,
		{3, 1}: 0,
		{5, 1}: 0,
		{7, 1}: 0,
		{1, 2}: 0}
	multipliedSlopes := 1
	for s := range slopes {
		for x, y := 0, 0; y < len(grid); x, y = (x+s.X)%len(grid[0]), y+s.Y {
			if grid[y][x] == '#' {
				slopes[s]++
			}
		}
		multipliedSlopes *= slopes[s]
	}

	fmt.Printf("Part 1 solution: %v\n", slopes[image.Point{3, 1}])
	fmt.Printf("Part 2 solution: %v\n", multipliedSlopes)
}
