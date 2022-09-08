package main

func Sum(numbers [5]int) int {

	var sum int

	for _, elem := range numbers {
		sum += elem
	}

	return sum
}
