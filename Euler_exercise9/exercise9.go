package main

// import(
// 	"math"
// )

// func tryPythagoreanTriplet(first_number, second_number float64) int {
// 	third_number := int(math.Sqrt(math.Pow(first_number, 2) + math.Pow(second_number, 2)))
// 	return third_number
// }

// func isPythagoreanTriplet(first_number, second_number, third_number int) bool {
// 	if first_number >= second_number || first_number >= third_number {
// 		return false
// 	}
// 	if second_number <= first_number || second_number >= third_number {
// 		return false
// 	}
// 	if third_number <= first_number || third_number <= second_number {
// 		return false
// 	}
// 	return true
// }

// func sumPythagoreanTripletEquals1000() int {
// 	result := 0
// 	i := 1
// 	j := 2

// 	for result != 1000 {
// 		third_number := tryPythagoreanTriplet(float64(i), float64(j))
// 		if isPythagoreanTriplet(i, j, third_number) == true {
// 			result = i + j + third_number
// 		}
// 		j++
// 		if j == 30 {
// 			j =1
// 		}
// 		i++
//  	}
//  	return result
// }