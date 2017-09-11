package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/urfave/cli"
	"gopkg.in/mgo.v2"
)

var letters = []string{"a", "b", "c", "d", "e", "f", "g",
	"h", "i", "j", "k", "l", "m", "n",
	"o", "p", "r", "s", "t", "u", "v", "w", "x",
	"y", "z"}

//const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Statistics struct {
	Name    string
	Count   int
	Percent int
	Time    int
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randInt(min int, max int, n int) []int {
	random := make([]int, n)
	var i int
	for i = 0; i <= n-1; i++ {
		random[i] = rand.Intn(max) + min
	}
	return random
}

func validateString(input []string) error {
	newInput := strings.Join(input, "")
	for range newInput {
		matched, err := regexp.MatchString("[a-z]", newInput)
		if matched == false {
			return errors.New("Required letters")
		}
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func getSession() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}

func getRandomNumbers(numCount int, seconds int) []int {
	random := randInt(0, 100, numCount)
	fmt.Printf("You have %d seconds to remember", seconds)
	fmt.Print("-->", random)
	time.Sleep(time.Duration(seconds) * time.Second)
	clear()
	return random
}

func getRandomString(strCount int, seconds int) []string {
	random := make([]string, strCount)
	fmt.Printf("You have %d seconds to remember", seconds)
	for i := range random {
		random[i] = letters[rand.Intn(len(letters))]
	}
	fmt.Print("-->", []string(random))
	time.Sleep(time.Duration(seconds) * time.Second)
	clear()
	return random
}

func scanNumbers(numCount int) ([]int, error) {
	input := make([]int, numCount)
	fmt.Print("Enter numbers -->")
	for i := range input {
		_, err := fmt.Scanln(&input[i])
		if err != nil {
			return input[:i], err
		}
	}
	return input, nil
}

func scanString(strCount int) ([]string, error) {
	input := make([]string, strCount)
	fmt.Print("Enter string -->")
	for i := range input {
		_, err := fmt.Scanln(&input[i])
		if err != nil {
			return input[:i], err
		}
	}
	return input, nil
}

func checkString(random []string, input []string, strCount int, seconds int) (int, error) {
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
	err = collection.Insert(
		&Statistics{Name: "string", Count: strCount, Percent: percent, Time: seconds},
	)
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
	err = collection.Insert(
		&Statistics{Name: "int", Count: numCount, Percent: percent, Time: seconds},
	)
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
						"Specified -n flag must be greater than zero",
					)
				}
				seconds := c.Int("t")
				if seconds <= 0 {
					return errors.New("Time can't be negative or equal zero")
				}
				randomArray := getRandomNumbers(numCount, seconds)
				numbers, err := scanNumbers(numCount)
				if err != nil {
					return errors.New("Required numbers")
				}
				percentNum, err := checkNumbers(
					randomArray, numbers, numCount, seconds,
				)
				if err != nil {
					return errors.New(
						"Percent can't be calculated",
					)
				}
				fmt.Println("Percent of wrong answers", percentNum, "%")
				return nil
			},
		},
		{
			Name:      "string",
			ShortName: "str",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "n,strCount",
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
					return errors.New(
						"Specified -n or -strCoutn flag must be grater than zero",
					)
				}
				seconds := c.Int("t")
				if seconds <= 0 {
					return errors.New("Time can't be negative or equal zero")
				}
				randomStringArray := getRandomString(strCount, seconds)
				stringByte, err := scanString(strCount)
				err = validateString(stringByte)
				percentNumString, err := checkString(
					randomStringArray, stringByte, strCount, seconds,
				)
				if err != nil {
					return errors.New(
						"Can't calculate percent",
					)
				}
				fmt.Println("Percent of wrong answers", percentNumString, "%")
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
