package main

func Sum(numbers []int) int {

	var sum int

	for _, elem := range numbers {
		sum += elem
	}

	return sum
}
