package main

import (
	"fmt"
	"os"
	"strconv"

	docopt "github.com/docopt/docopt-go"
)

func main() {
	config, err := parseConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't find config", err)
	}

	usage := `Memory

Usage:
  memory  num  --count <count> [--time <time>] 
  memory  str  --count <count> [--time <time>] 
  memory -h | --help
  memory --version

Options:
  -h --help     	Show this screen.
  --version     	Show version.
  --count <count>	Count of numbers or letters. 
  --time <time>		Time for memorization. 
`

	arguments, err := docopt.Parse(usage, nil, true, "memory", false)
	if err != nil {
		fmt.Println(err)
	}
	var num = arguments["num"].(bool)
	if num {
		var c = arguments["--count"].(string)
		count, _ := strconv.Atoi(c)
		if t, ok := arguments["--time"].(string); ok {
			seconds, err := strconv.Atoi(t)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			err = handleNumbers(count, seconds, config)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			os.Exit(1)
		}
		seconds := DEFAULT_TIME
		err := handleNumbers(count, seconds, config)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	}
	var str = arguments["str"].(bool)
	if str {
		var c = arguments["--count"].(string)
		count, _ := strconv.Atoi(c)
		if t, ok := arguments["--time"].(string); ok {
			seconds, err := strconv.Atoi(t)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			err = handleStrings(count, seconds, config)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
		}
		seconds := DEFAULT_TIME
		err := handleStrings(count, seconds, config)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}

	}
}
