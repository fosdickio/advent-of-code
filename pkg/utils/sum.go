package utils

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
