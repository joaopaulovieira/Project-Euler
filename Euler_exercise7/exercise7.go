package main

func is_prime(number int) bool {
	test := 0

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			test++
		}
	}
	if test > 2 {
		return false
	}
	return true
}

func find_prime_number(number int) int {
	result := 0
	position := 0

	for i := 2; position != number; i++ {
		if is_prime(i) == true {
			result = i
			position++
		}
	}
	return result
}
