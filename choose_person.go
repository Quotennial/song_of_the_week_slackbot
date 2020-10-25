package main

import (
	"math/rand"
	"time"
)

func pickRandomPerson() string {
	masterList, err := readPersonStore("masterlist")
	if err != nil {
		return "I couldn't find the master list"
	}
	return masterList.randomFromList()
}

func (p personStore) randomFromList() string {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator - is this ok if triggered every week same time?
	return p[rand.Intn(len(p))]
}

func getSongOfTheWeek() {}

//randomiser function takes in person list and returns selection as string
// func pickRandomName(p personStore) string {}
