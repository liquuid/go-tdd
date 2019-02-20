package array_slices

func Sum(arr []int) int {
	sum := 0
	for _, item := range arr {
		sum += item
	}
	return sum
}

func SumAll(arr ...[]int) []int{
	var sums []int
	for _, num := range arr{
		sums = append(sums, Sum(num))
	}
	return sums
}

func SumAllTails(arr ...[]int) []int{
	var sums []int
	for _, nums := range arr{
		if len(nums) != 0 {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
