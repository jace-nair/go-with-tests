package slices

var MyArray1 [5]int = [5]int{0, 1, 2, 3, 4}
var MyArray2 [5]int = [5]int{6, 7, 8, 9, 10}

var MySlice1 []int = []int{0, 1, 2, 3, 4}
var MySlice2 []int = []int{6, 7, 8, 9, 10}

// Adds up elements of an Array
func SumArray(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}

// Adds up elements of a Slice
func SumSlice(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

// Sums all elements of the slices passed in and returns it as a new slice
func SumAll(numbersToSum ...[]int) []int {

	var newSlice []int

	for _, numbers := range numbersToSum {
		newSlice = append(newSlice, SumSlice(numbers))
	}
	return newSlice
}

// Sums tail elements of the slices passed in and returns it as a new Slice
func SumAllTails(slices ...[]int) []int {

	var newSlice []int

	for _, v := range slices {
		if len(v) == 0 {
			newSlice = append(newSlice, 0)
		} else {
			tail := v[1:] // choose values from index 1 to the end omitting the value for the first index 0
			newSlice = append(newSlice, SumSlice(tail))
		}
	}
	return newSlice
}
