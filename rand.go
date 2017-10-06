package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateRandomNumbers(count int) []int {
	max := 100
	min := 0
	i := 0

	random := make([]int, count)
	for i = 0; i <= count-1; i++ {
		random[i] = rand.Intn(max) + min
	}
	return random
}

func generateRandomStrings(count int) []string {
	random := make([]string, count)
	for i := range random {
		random[i] = letters[rand.Intn(len(letters))]
	}
	return random
}
