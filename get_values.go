package main

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

func getValuesForNumbers(seconds int, count int, randomNumbersArray []int,
	sessionForNumbers *mgo.Session, collectionForNumbers *mgo.Collection) error {

	fmt.Print("Enter numbers -->")
	inputNumbersArray, err := scanNumbers(count)
	if err != nil {
		return fmt.Errorf("input values must be numbers, %s", err)
	}
	percentOfWrongNumbers := calculatePercentForNumbers(
		randomNumbersArray, inputNumbersArray, count,
	)
	err = insertInDB(
		sessionForNumbers, collectionForNumbers, count,
		percentOfWrongNumbers, seconds,
	)
	if err != nil {
		return fmt.Errorf("can't insert data in database %s", err)
	}
	fmt.Println("percent of wrong answers", percentOfWrongNumbers, "%")
	return nil
}

func getValuesForStrings(
	seconds int, count int, randomStringArray []string,
	sessionForStrings *mgo.Session, collectionForStrings *mgo.Collection) error {

	fmt.Print("Enter string -->")
	inputStringArray, err := scanStrings(count)
	if err != nil {
		return fmt.Errorf("input values must be letters,%s", err)
	}
	percentOfWrongString := calculatePercentForStrings(
		randomStringArray, inputStringArray, count,
	)
	err = insertInDB(
		sessionForStrings, collectionForStrings, count,
		percentOfWrongString, seconds,
	)
	if err != nil {
		return fmt.Errorf("can't insert data in collection, %s", err)
	}
	fmt.Println("Percent of wrong answers", percentOfWrongString, "%")

	return nil
}
