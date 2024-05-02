package array

func Sum(nums []int) int {
	var sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(nums ...[]int) []int {
	//sums := make([]int, len(nums))
	var sums []int
	for _, num := range nums {
		sums = append(sums, Sum(num))
	}
	return sums
}
func SumAllTails(nums ...[]int) []int {
	var sums []int
	for _, num := range nums {
		if len(num) == 0 {
			sums = append(sums, 0)
		} else {
			tail := num[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
