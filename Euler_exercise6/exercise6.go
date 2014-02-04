package main

import(
	"math"
)

func sum_squares(number float64) float64 {
	var sum float64 = 0
	var i float64 = 1
	for ; i <= number; i++ {
		sum += math.Pow(i, 2)
	}
	return sum
}

func square_sum(number float64) float64 {
	var sum float64 = 0
	var i float64 = 1
	for ; i <= number; i++ {
		sum += i
	}
	return math.Pow(sum, 2)
}

func diference_between_square_sum_and_sum_squares(number float64) float64 {
	return square_sum(number) - sum_squares(number)
}