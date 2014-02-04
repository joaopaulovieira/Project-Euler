package main

import(
	"fmt"
	"math"
	)

func sumPythagoreanTripletEquals1000(first_number, second_number, third_number int) bool {
	result := first_number + second_number + third_number
	if result == 1000 {
		return  true
	}
	return false
}

func isPythagoreanTriplet(first_number, second_number, third_number int) bool {
	if math.Pow(float64(first_number), 2) + math.Pow(float64(second_number), 2) == math.Pow(float64(third_number), 2) {
		return true
	}
	return false
}

func pythagoreanTriplet() int {
	result := 0
	test := false
	aux := false

	for a := 1; a <= 1000; a++ {
		for b := a + 1; b <= 1000; b++ {
			for c := b + 1; c <= 1000; c++ {
				test = sumPythagoreanTripletEquals1000(a, b, c)
				if test == true {
					aux = isPythagoreanTriplet(a, b, c)
					if aux == true {
						result = a*b*c
						fmt.Println(result)
					}
				}
			}
		}
	}
	return result
}