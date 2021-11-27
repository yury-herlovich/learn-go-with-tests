package main

func Sum(numbers []int) int {
	sum := 0

	for _, n := range numbers {
		sum += n
	}

	return sum
}

func SumAll(args ...[]int) []int {
	var sums []int

	for _, arg := range args {
		sums = append(sums, Sum(arg))
	}

	return sums
}

func SumAllTrails(args ...[]int) []int {
	var sums []int

	for _, nums := range args {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(nums[1:]))
		}
	}

	return sums
}
