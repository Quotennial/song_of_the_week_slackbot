package main

import (
	"math/rand"
	"time"
)

func pickPerson(listOfPeople []string) string {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator - is this ok if triggered every week same time?
	sentence := sentenceGenerator(listOfPeople[rand.Intn(len(listOfPeople))])
	return sentence
}

func sentenceGenerator(personID string) string {
	// add lots of different sentences and pick one at random - find a way to do this
	sentence := "<@" + personID + ">, you're up for song of the week!"
	// create lots of strings to trail the name and then randomly select from that list
	return sentence
}

func getSongOfTheWeek() {}

//randomiser function takes in person list and returns selection as string
// func pickRandomName(p personStore) string {}
