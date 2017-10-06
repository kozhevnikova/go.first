package main

import "fmt"

func clearByEnter() {
	var enter string
	fmt.Scanln(&enter)
	print("\033[H\033[2J")
}

func clearByTime() {
	print("\033[H\033[2J")
}
