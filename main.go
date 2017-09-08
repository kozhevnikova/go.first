package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/urfave/cli"
	"gopkg.in/mgo.v2"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Statistics struct {
	Name    string
	Count   int
	Percent int
	Time    int
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func getSession() (*mgo.Session, error) {
	err := errors.New("Can not connect to DB")
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Print(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}

func getRandomNumbers(numCount int, seconds int) []int {
	random := make([]int, numCount)
	for i, value := range rand.Perm(numCount) {
		random[i] = value
	}
	fmt.Printf("You have %d seconds to remember", seconds)
	fmt.Print(random)
	time.Sleep(time.Duration(seconds) * time.Second)
	clear()
	return random
}

func getRandomString(strCount int, seconds int) string {
	random := make([]byte, strCount)
	fmt.Printf("You have %d seconds to remember", seconds)
	for i := range random {
		random[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	fmt.Print(string(random))
	time.Sleep(time.Duration(seconds) * time.Second)
	clear()
	return string(random)
}

func scanNumbers(numCount int) ([]int, error) {
	input := make([]int, numCount)
	fmt.Print("Enter numbers: ")
	fmt.Print("->")
	for i := range input {
		_, err := fmt.Scan(&input[i])
		if err != nil {
			return input[:i], errors.New("Cannot scan numbers")
		}
	}
	return input, nil
}

func scanString() (string, error) {
	err := errors.New("Cannot read string")
	fmt.Print("Enter string: ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("->")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print(err)
	}
	return text, nil
}

func checkString(random string, input string, strCount int, seconds int) (int, error) {
	newSession, err := getSession()
	collection := newSession.DB("test").C("str")
	var wrong int
	var percent int
	for i, _ := range random {
		if input[i] != random[i] {
			wrong = wrong + 1
		}
	}
	percent = (wrong * 100) / strCount
	err = collection.Insert(&Statistics{Name: "string", Count: strCount, Percent: percent, Time: seconds})
	if err != nil {
		panic(err)
	}
	return percent, nil
}

func checkNumbers(random []int, input []int, numCount int, seconds int) (int, error) {
	newSession, err := getSession()
	collection := newSession.DB("test").C("int")
	var wrong int
	var percent int
	for i, _ := range random {
		if input[i] != random[i] {
			wrong = wrong + 1
		}
	}
	percent = (wrong * 100) / numCount
	err = collection.Insert(&Statistics{Name: "int", Count: numCount, Percent: percent, Time: seconds})
	if err != nil {
		panic(err)
	}
	return percent, nil
}

func main() {
	app := cli.NewApp()
	app.Usage = "Try your memory. Choose numbers or string."
	app.Commands = []cli.Command{
		{
			Name:      "numbers",
			ShortName: "num",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "n, numCount",
					Usage: "Count of numbers",
					Value: 10,
				},
				cli.IntFlag{
					Name:  "t",
					Usage: "Set seconds in seconds for memorization",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				numCount := c.Int("n")
				if numCount <= 0 {
					return errors.New(
						"specified -n flag must be greater than zero",
					)
				}
				seconds := c.Int("t")
				if seconds <= 0 {
					return errors.New("Time cannot be negative or equal zero")
				}
				randomArray := getRandomNumbers(numCount, seconds)
				numbers, err := scanNumbers(numCount)
				if err != nil {
					return errors.New("Cannot scan numbers")
				}
				percentNum, err := checkNumbers(randomArray, numbers, numCount, seconds)
				if err != nil {
					return errors.New("Cannot calculate percent")
				}
				fmt.Println("Percent of wrong answers", percentNum)
				return nil
			},
		},
		{
			Name:      "string",
			ShortName: "str",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "n",
					Usage: "Count of letters",
					Value: 10,
				},
				cli.IntFlag{
					Name:  "t",
					Usage: "Set time in seconds for memorization",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				strCount := c.Int("n")
				if strCount <= 0 {
					return errors.New("strCount must be grater than zero")
				}
				seconds := c.Int("t")
				if seconds <= 0 {
					return errors.New("Time cannot be negative")
				}
				randomStringArray := getRandomString(strCount, seconds)
				stringByte, err := scanString()
				if err != nil {
					errors.New("Cannot scan string")
				}
				percentNumString, err := checkString(randomStringArray, stringByte, strCount, seconds)
				if err != nil {
					return errors.New("Cannot calculate percent")
				}
				fmt.Println("Percent of wrong answers", percentNumString)
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
