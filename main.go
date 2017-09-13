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
			return err
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

func getRandomNumbers(count int, seconds int) []int {
	random := randInt(0, 100, count)
	fmt.Printf("You have %d seconds to remember", seconds)
	fmt.Print("-->", random)
	time.Sleep(time.Duration(seconds) * time.Second)
	clear()
	return random
}

func getRandomString(count int, seconds int) []string {
	random := make([]string, count)
	fmt.Printf("You have %d seconds to remember", seconds)
	for i := range random {
		random[i] = letters[rand.Intn(len(letters))]
	}
	fmt.Print("-->", []string(random))
	time.Sleep(time.Duration(seconds) * time.Second)
	clear()
	return random
}

func scanNumbers(count int) ([]int, error) {
	input := make([]int, count)
	fmt.Print("Enter numbers -->")
	for i := range input {
		_, err := fmt.Scanln(&input[i])
		if err != nil {
			return input[:i], err
		}
	}
	return input, nil
}

func scanString(count int) ([]string, error) {
	input := make([]string, count)
	fmt.Print("Enter string -->")
	for i := range input {
		_, err := fmt.Scanln(&input[i])
		if err != nil {
			return input[:i], err
		}
	}
	return input, nil
}

func getStringPercent(random []string, input []string, count int, seconds int) int {
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

func getNumPercent(random []int, input []int, count int, seconds int) int {
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

func insertInDB(count int, percent int, seconds int, ShortName string) error {
	newSession, err := getSession()
	if err != nil {
		return err
	}
	collection := newSession.DB("go").C("data")
	err = collection.Insert(
		&Statistics{Name: ShortName, Count: count, Percent: percent, Time: seconds},
	)
	if err != nil {
		return err
	}
	return nil
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
					Name:  "n, count",
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
				count := c.Int("n")
				if count <= 0 {
					return errors.New(
						"Specified -n flag must be greater than zero",
					)
				}
				seconds := c.Int("t")
				if seconds <= 0 {
					return errors.New("Time can't be negative or equal zero")
				}
				randomArray := getRandomNumbers(count, seconds)
				numbers, err := scanNumbers(count)
				if err != nil {
					return errors.New("Required numbers")
				}
				percentNum := getNumPercent(
					randomArray, numbers, count, seconds,
				)
				err = insertInDB(count, percentNum, seconds, "num")
				if err != nil {
					return fmt.Errorf("Can't insert data: %s", err)
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
					Name:  "n,count",
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
				count := c.Int("n")
				if count <= 0 {
					return errors.New(
						"Specified -n or -strCoutn flag must be grater than zero",
					)
				}
				seconds := c.Int("t")
				if seconds <= 0 {
					return errors.New("Time can't be negative or equal zero")
				}
				randomStringArray := getRandomString(count, seconds)
				stringByte, err := scanString(count)
				if err != nil {
					return fmt.Errorf("Can't scan string: %s", err)
				}
				err = validateString(stringByte)
				if err != nil {
					return fmt.Errorf("%s", err)
				}
				percentString := getStringPercent(
					randomStringArray, stringByte, count, seconds,
				)
				err = insertInDB(count, percentString, seconds, "str")
				if err != nil {
					return fmt.Errorf("Can't insert data: %s", err)
				}
				fmt.Println("Percent of wrong answers", percentString, "%")
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
