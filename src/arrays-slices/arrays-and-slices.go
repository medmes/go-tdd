// Arrays allow you to store multiple elements of the same type in a variable in a particular order.
package arrays_slices

func Sum(numbers [5]int) int {

	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}
