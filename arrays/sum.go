// summing with for loops
package sum

func Sum(nums []int) int {
	var solution int

	for _, num := range nums {
		solution += num
	}

	return solution
}

// takes multiple array slices, sums each one, returns each sum in a slice
func SumAll(slicesToSum ...[]int) (solution []int) {
	for _, nums := range slicesToSum {
		solution = append(solution, Sum(nums))
	}
	return
}
