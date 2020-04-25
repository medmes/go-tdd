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
	// we will worry less if we use append instead of a direct slice, because as we know slices have a capacity
	// and if we do mySlice[10] = 2 it gives a runtime error
	for _, numbers := range numbers {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
