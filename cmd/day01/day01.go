package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputFile string = "./assets/input/day01.txt"
const sumTarget int = 2020

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	entries := make(map[int]int)
	readInputFileIntoMap(f, entries)
	f.Close()

	checkMapForTargetSum1(entries, sumTarget)
	checkMapForTargetSum2(entries, sumTarget)
}

func readInputFileIntoMap(f *os.File, entries map[int]int) {
	input := bufio.NewScanner(f)
	i := 0
	for input.Scan() {
		intValue, err := strconv.Atoi(input.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		entries[i] = intValue
		i++
	}
}

func checkMapForTargetSum1(entries map[int]int, target int) int {
	fmt.Printf("Finding 2 entries that sum to %v...\n", target)
	var answer int
	for _, x := range entries {
		for _, y := range entries {
			if x+y == target {
				answer = x * y
				fmt.Printf("%v + %v = %v ; %v * %v = %v\n", x, y, target, x, y, answer)
			}
		}
	}
	return answer
}

func checkMapForTargetSum2(entries map[int]int, target int) int {
	fmt.Printf("Finding 3 entries that sum to %v...\n", target)
	var answer int
	for _, x := range entries {
		for _, y := range entries {
			for _, z := range entries {
				if x+y+z == target {
					answer = x * y * z
					fmt.Printf("%v + %v + %v = %v ; %v * %v * %v = %v\n", x, y, z, target, x, y, z, answer)
				}
			}
		}
	}
	return answer
}
