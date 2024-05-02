package array

func Sum(nums [5]int) int {
	var sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
