package array

func Sum(nums []int) int {
	var sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(nums ...[]int) []int {
	sums := make([]int, len(nums))
	for i, num := range nums {
		sums[i] = Sum(num)
	}
	return sums
}
