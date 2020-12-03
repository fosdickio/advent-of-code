/*
Package utils/sum implements a simple library for checking whether data from a given input file sums to a given target.
*/
package utils

// TwoSum reads in data from an integer slice and checks if there are two values that sum to the given target.  If so,
// the multiple of those two numbers is returned.  If not, 0 is returned.
func TwoSum(data []int, target int) int {
	var seen = make(map[int]bool, len(data))

	for _, v := range data {
		if seen[v] {
			return v * (target - v)
		}
		seen[target-v] = true
	}

	return 0
}

// ThreeSum reads in data from an integer slice and checks if there are three values that sum to the given target.  If
// so, the multiple of those three numbers is returned.  If not, 0 is returned.
func ThreeSum(data []int, target int) int {
	seen := make(map[int]bool, len(data))

	for i, x := range data {
		for _, y := range data[i+1:] {
			z := target - x - y
			if seen[z] {
				return x * y * z
			}
		}
		seen[x] = true
	}

	return 0
}
