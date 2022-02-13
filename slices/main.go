package main

func sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func sumAll(numberSlices ...[]int) (sums []int) {
	for _, numbers := range numberSlices {
		sums = append(sums, sum(numbers))
	}

	return
}
