// summing with for loops
package sum

func Sum(nums [5]int) int {
	var solution int

	for i := range len(nums) {
		solution += nums[i]
	}

	return solution
}
