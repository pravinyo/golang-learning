package main

func Sum(numbers []int) int {

	var sum int

	for _, elem := range numbers {
		sum += elem
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numbers))
		}
	}

	return sums
}
