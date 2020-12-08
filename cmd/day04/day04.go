package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

const inputFile string = "./assets/input/day04.txt"

var passportFormat = []*regexp.Regexp{
	regexp.MustCompile(`(?:^|\s)(byr):(?:(19[2-9]\d|200[0-2])(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(iyr):(?:(201\d|2020)(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(eyr):(?:(202\d|2030)(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(hgt):(?:((?:1[5-8]\d|19[0-3])cm|(?:59|6\d|7[0-6])in)(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(hcl):(?:(#[\da-f]{6})(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(ecl):(?:(amb|blu|brn|gry|grn|hzl|oth)(?:\s|$))?`),
	regexp.MustCompile(`(?:^|\s)(pid):(?:(\d{9})(?:\s|$))?`),
}

func main() {
	input, _ := ioutil.ReadFile(inputFile)

	part1, part2 := 0, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n\n") {
		valid1, valid2 := 1, 1
		for _, regex := range passportFormat {
			if match := regex.FindStringSubmatch(s); len(match) == 0 {
				valid1, valid2 = 0, 0
			} else if match[2] == "" {
				valid2 = 0
			}
		}
		part1, part2 = part1+valid1, part2+valid2
	}

	fmt.Printf("Part 1 solution: %v\n", part1)
	fmt.Printf("Part 2 solution: %v\n", part2)
}
