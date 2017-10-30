package main

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

func handleNumbers(count int, seconds int, config Config) error {
	//create session. start
	sessionForNumbers, err := mgo.Dial(config.Server.Address)
	if err != nil {
		return fmt.Errorf(
			"unable to establish connection with database %s, %s",
			config.Server.Address, err,
		)
	}
	collectionForNumbers := sessionForNumbers.DB(
		config.Database["numbers"].DBName).C(config.Database["numbers"].Collection)
	//end

	if count <= 0 {
		return errors.New(
			"specified --count flag must be greater than zero",
		)
	}

	if seconds < 0 {
		return errors.New("time flag must be positive number")
	}

	if seconds == 0 {
		randomNumbersArray := generateRandomNumbers(count)
		fmt.Print("Time is unlimited. Press Enter when you will be ready -->")
		fmt.Println(randomNumbersArray)
		clearByEnter()
		err := getValuesForNumbers(
			seconds, count, randomNumbersArray, sessionForNumbers,
			collectionForNumbers,
		)
		if err != nil {
			return err
		}
	}

	if seconds > 0 {
		randomNumbersArray := generateRandomNumbers(count)
		fmt.Printf("You have %d seconds to remember", seconds)
		fmt.Print("-->", randomNumbersArray)
		time.Sleep(time.Duration(seconds) * time.Second)
		clear()
		err := getValuesForNumbers(
			seconds, count, randomNumbersArray, sessionForNumbers,
			collectionForNumbers,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func handleStrings(count int, seconds int, config Config) error {
	//create session. start
	sessionForStrings, err := mgo.Dial(config.Server.Address)
	if err != nil {
		return fmt.Errorf("unable to establish connection with database %s, %s",
			config.Server.Address, err,
		)
	}
	collectionForStrings := sessionForStrings.DB(
		config.Database["strings"].DBName).C(config.Database["strings"].Collection)
	//end

	if count <= 0 {
		return errors.New(
			"specified --count  flag must be grater than zero",
		)
	}
	if seconds < 0 {
		return errors.New("time can't be negative or equal zero")
	}
	if seconds == 0 {
		randomStringsArray := generateRandomStrings(count)
		fmt.Print("Time is unlimited.Press Enter when you will be ready -->")
		fmt.Println(randomStringsArray)
		clearByEnter()
		err := getValuesForStrings(
			seconds, count, randomStringsArray, sessionForStrings,
			collectionForStrings,
		)
		if err != nil {
			return err
		}
	}
	if seconds > 0 {
		randomStringsArray := generateRandomStrings(count)
		fmt.Printf("You have %d seconds to remember", seconds)
		fmt.Print("-->", []string(randomStringsArray))
		time.Sleep(time.Duration(seconds) * time.Second)
		clear()
		err := getValuesForStrings(
			seconds, count, randomStringsArray, sessionForStrings,
			collectionForStrings,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
