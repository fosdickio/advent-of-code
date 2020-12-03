/*
Package utils/password implements a simple library for checking password policies of a given input.

Input will look like the following:

	1-3 a: abcde
	1-3 b: cdefg
	2-9 c: ccccccccc

Each line gives the password policy and then the password. The password policy indicates the lowest and highest number
of times a given letter must appear for the password to be valid. For example, 1-3 a means that the password must
contain a at least 1 time and at most 3 times.
*/
package utils

import (
	"fmt"
	"strings"
)

// CountValidPasswords takes a byte slice of various password policies and passwords and returns the count of the total
// number of passwords that adhere to the password policy.
func CountValidPasswords(input []byte) (solutionPart1 int, solutionPart2 int) {
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
