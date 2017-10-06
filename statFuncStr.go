package main

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

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
