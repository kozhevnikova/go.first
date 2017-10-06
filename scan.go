package main

import "fmt"

func scanNumbers(count int) ([]int, error) {
	input := make([]int, count)
	for i := range input {
		_, err := fmt.Scanln(&input[i])
		if err != nil {
			return input[:i], err
		}
	}
	return input, nil
}

func scanStrings(count int) ([]string, error) {
	input := make([]string, count)
	for i := range input {
		_, err := fmt.Scanln(&input[i])
		if err != nil {
			return input[:i], err
		}
	}
	return input, nil
}
