package main

import "fmt"

func clearByEnter() {
	var enter string
	fmt.Scanln(&enter)
	print("\033[H\033[2J")
}

func clear() {
	print("\033[H\033[2J")
}
