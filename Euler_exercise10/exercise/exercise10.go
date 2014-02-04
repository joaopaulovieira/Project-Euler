package main

import(
	"fmt"
)

func is_prime(number float64) bool {
	test := 0
	var i float64 = 1

	for ; i <= number; i++ {
		if int(number)%int(i) == 0 {
			test++
		}
	}
	if test > 2 {
		return false
	}
	return true
}

func sumPrimeNumbersBelowsParamater(number float64) float64 {
	var result float64 = 0
	var i float64 = 2
	for ; i <= number; i++ {
		if is_prime(i) == true {
			result += i
		}
	}
	return result
}

func main() {
	fmt.Println(sumPrimeNumbersBelowsParamater(2000000))	
}