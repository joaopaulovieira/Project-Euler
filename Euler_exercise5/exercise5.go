package main

func is_multiple(number, multiple int) bool {
	return number%multiple == 0 
}

func smallest_multiple(number int) int {
	small := 0
	test := 1
	aux := false

	for small != test {
		for j := 1; j <= number; j++ {
			if is_multiple(test, j) == false {
				aux = false
			}
		}
		if aux == true {
			small = test
		}
		if aux == false {
			test++
		}
		aux = true
	}
	return small
}