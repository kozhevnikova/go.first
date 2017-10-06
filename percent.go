package main

func calculatePercentForStrings(random []string, input []string, count int) int {
	var wrong int
	var percent int
	for i, _ := range random {
		if input[i] != random[i] {
			wrong = wrong + 1
		}
	}
	percent = (wrong * 100) / count
	return percent
}

func calculatePercentForNumbers(random []int, input []int, count int) int {
	var wrong int
	var percent int
	for i, _ := range random {
		if input[i] != random[i] {
			wrong = wrong + 1
		}
	}
	percent = (wrong * 100) / count
	return percent
}
