package main

func palindromic_number(number int) bool {
	aux, invertido, digito := 0, 0, 0
	aux = number
	for aux != 0 {
		digito = aux % 10
		invertido = (invertido * 10) + digito
		aux = aux / 10
	}
	if number == invertido {
		return true
	}
	return false
}

func largest_palindromic_number() int {
	test := 0
	palindromic := 0
	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			if palindromic_number(i*j) == true {
				test = i*j
			}
			if test != 0 && palindromic < test {
			palindromic = test
			}
		}
	}
	return palindromic
}