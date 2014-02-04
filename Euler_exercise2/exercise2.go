package main

func below_four_million(first int) bool {
	if first > 4000000 {
		return true
	}
	return false
}

func return_even_number(number int) bool {
	if number%2 == 0 {
		return true
	}
	return false
}

func fibonacci_sequence() func() int {
	first := 1
	next := 1
	return func() int {
		if below_four_million(first) == true {
			return 0
		}
		first, next = next, next+first
		if return_even_number(first) == false {
			return 0
		}
		return first
	}
}

func sum_sequence(range_sequence int) int {
	result := 0
	fibonacci := fibonacci_sequence()

	for i := 1; i <= range_sequence; i++ {
		result += fibonacci()
	}
	return result
}