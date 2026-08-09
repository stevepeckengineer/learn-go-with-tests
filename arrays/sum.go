// summing with for loops
package sum

func Sum(nums [5]int) int {
	var solution int

	for _, num := range nums {
		solution += num
	}

	return solution
}
