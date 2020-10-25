package main

import (
	"math/rand"
	"time"
)

func pickRandomPerson(listFilePath string) (string, error) {
	masterList, err := readPersonStore(listFilePath + "masterlist")
	if err != nil {
		return "`Error:` I couldn't find the master list", &CustomError{}
	}
	return masterList.randomFromList(), nil
}

func (p personStore) randomFromList() string {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator - is this ok if triggered every week same time?
	return p[rand.Intn(len(p))]
}

func getSongOfTheWeek() {}

//randomiser function takes in person list and returns selection as string
// func pickRandomName(p personStore) string {}
