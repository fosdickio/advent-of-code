package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fosdickio/advent-of-code/pkg/utils"
)

const inputFile string = "./assets/input/day01.txt"
const sumTarget int = 2020

func main() {
	input, err := utils.ReadIntInputFileIntoSlice(inputFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}

	fmt.Printf("Part 1 solution: %v\n", utils.TwoSum(input, sumTarget))
	fmt.Printf("Part 2 solution: %v\n", utils.ThreeSum(input, sumTarget))
}
