package main

func is_multiple(number, multiple int) bool {
	return number % multiple == 0
}

func sum_numbers(number int) int {
	result := 0
	for i := 0; i < 1000; i++ {
		if is_multiple(i, 3) || is_multiple(i, 5) {
			result += i
		}
	}
	return result
}
