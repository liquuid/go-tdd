package array_slices

func Sum(arr []int) int {
	sum := 0
	for _, item := range arr {
		sum += item
	}
	return sum
}
