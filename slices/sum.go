package slices

func SumArray(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}

func SumSlice(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}
