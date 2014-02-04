package main

import(
	"fmt"
)

func factorPrime(number int) int {
	factor := 0
	check := false
	for i := 1 ; i <= number; i++ {
		check = MMC(number, i)
		if check == true {
			factor++
		}
		check = false
	}
	return factor
}

func MMC(number, minimum int) bool {
	if number%minimum == 0 {
		return true
	}
	return false
}

func makeTriangularNumberReturnDivisors(limit int) int {
	result := 0
	aux := false
	factor_prime := 0

	for i := 1; aux != true; i++ {
		result = result + i
		factor_prime = factorPrime(result)
		if factor_prime >= limit {
			aux = true
		}
	}
	fmt.Println(result)
	return result
}