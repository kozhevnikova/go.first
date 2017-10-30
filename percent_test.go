package main

import (
	"testing"
)

func TestHundredPercentForNumbers(t *testing.T) {
	var input = [][]int{{2, 1}, {1, 1}, {-1, -1}}
	var random = [][]int{{0, 0}, {100, 100}}
	var count = 2
	var percentOfWrong = 100

	for _, randomVariable := range random {
		for _, input := range input {
			result := calculatePercentForNumbers(randomVariable, input, count)
			if result != percentOfWrong {
				t.Errorf("Percent of %d and %d must be equal %d, but result %d",
					randomVariable, input, percentOfWrong, result)
			}
		}
	}

}

func TestFiftyPercentForNumbers(t *testing.T) {
	var input = [][]int{{0, 1}, {1, 0}, {11, 1}}
	var random = [][]int{{1, 1}}
	var count = 2
	var percentOfWrong = 50

	for _, randomVariable := range random {
		for _, input := range input {
			result := calculatePercentForNumbers(randomVariable, input, count)
			if result != percentOfWrong {
				t.Errorf("Percent of %d and %d must be equal %d, but result %d",
					randomVariable, input, percentOfWrong, result)
			}
		}
	}

}
func TestOnePercentForNumbers(t *testing.T) {
	var input = [][]int{{22, 22}}
	var random = [][]int{{22, 22}}
	var count = 2
	var percentOfWrong = 0

	for _, randomVariable := range random {
		for _, input := range input {
			result := calculatePercentForNumbers(randomVariable, input, count)
			if result != percentOfWrong {
				t.Errorf("Percent of %d and %d must be equal %d, but result %d",
					randomVariable, input, percentOfWrong, result)
			}
		}
	}

}

func TestOneHundredPercentStrings(t *testing.T) {
	var input = [][]string{{"a", "a"}, {"0", " 1"}, {"-1", "1"}, {"qq", "qq"}}
	var random = [][]string{{"b", "b"}, {"q", "b"}}
	var count = 2
	var percentOfWrong = 100

	for _, randomVariable := range random {
		for _, inputVariable := range input {
			result := calculatePercentForStrings(randomVariable, inputVariable, count)
			if result != percentOfWrong {
				t.Errorf("Percent of %s and %s must be equal %d, but result %d",
					randomVariable, inputVariable, percentOfWrong, result)
			}
		}
	}

}

func TestFiftyPercentForStrings(t *testing.T) {
	var input = [][]string{{"b", "bb"}}
	var random = [][]string{{"b", "b"}}
	var count = 2
	var percentOfWrong = 50

	for _, randomVariable := range random {
		for _, inputVariable := range input {
			result := calculatePercentForStrings(randomVariable, inputVariable, count)
			if result != percentOfWrong {
				t.Errorf("Percent of %s and %s must be equal %d, but result %d",
					randomVariable, inputVariable, percentOfWrong, result)
			}
		}
	}

}

func TestOnePercentStrings(t *testing.T) {
	var input = [][]string{{"b", "b"}}
	var random = [][]string{{"b", "b"}}
	var count = 2
	var percentOfWrong = 0

	for _, randomVariable := range random {
		for _, inputVariable := range input {
			result := calculatePercentForStrings(randomVariable, inputVariable, count)
			if result != percentOfWrong {
				t.Errorf("Percent of %s and %s must be equal %d, but result %d",
					randomVariable, inputVariable, percentOfWrong, result)
			}
		}
	}

}
