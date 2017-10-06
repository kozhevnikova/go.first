package main

import mgo "gopkg.in/mgo.v2"

func insertInDB(s *mgo.Session, c *mgo.Collection, count int, percent int, seconds int) error {
	err := c.Insert(
		&Statistics{Count: count, Percent: percent, Time: seconds},
	)
	if err != nil {
		return err
	}
	return nil
}
