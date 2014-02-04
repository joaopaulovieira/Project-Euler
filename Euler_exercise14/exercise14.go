package main

func makeCollatzSequence(number int) int {
	if number%2 == 0  {
		number = number/2
	} else {
		number = (3*number) + 1
	}
	return number
}

func largestCollatzSequence(limit int) int {
	count := 0
	collatz_result := 0
	aux := 0
	result := 0

	for i := 2; i < limit; i++ {
		collatz_result = i
		for collatz_result != 1 {
			collatz_result = makeCollatzSequence(collatz_result)
			count++
		}
		collatz_result = 0
		if count > aux {
			aux = count
			result = i
		}
		count = 0
	}
	return result
}