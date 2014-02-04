package main

func factor_prime(number int) int {
	i := 2
	factor := 0
	for ; number != 1; i++ {
		number, i = MMC(number, i)
		factor = i
	}
	return factor
}

func MMC(number, minimum int) (int, int) {
	if number%minimum == 0 {
		number = number / minimum
	}
	return number, minimum
}