package main

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func sumAll(numberSlices ...[]int) (sums []int) {
	sums = make([]int, len(numberSlices))

	for i, numbers := range numberSlices {
		sums[i] = sum(numbers)
	}

	return
}
