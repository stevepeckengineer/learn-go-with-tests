// summing with for loops
package sum

func Sum(nums []int) int {
	var solution int

	for _, num := range nums {
		solution += num
	}

	return solution
}
