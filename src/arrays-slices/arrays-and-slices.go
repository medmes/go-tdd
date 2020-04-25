// Arrays allow you to store multiple elements of the same type in a variable in a particular order.
package arrays_slices

func Sum(numbers []int) int {

	sum := 0
	for _, v := range numbers { // (k, v) -> key (array index), value (array value).
		sum += v
	}
	return sum
}

func SumAll(numbers ...[]int) (sums []int) {
	lengthOfNumbers := len(numbers)
	sums = make([]int, lengthOfNumbers)

	// iterating on list of numbers and use Sum function to calculate
	for k, v := range numbers {
		sums[k] += Sum(v)
	}
	return
}
