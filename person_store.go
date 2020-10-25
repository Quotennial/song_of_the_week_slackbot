package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

type personStore []string

// Functions turning person store into files and reading those files
func (p personStore) toString() string {
	//use a go std library: strings, input and then separator
	//first turn deck d into a string
	return strings.Join([]string(p), ",")

}
func (p personStore) saveToFile(filename string) error {
	//using the writefile from ioutil library, takes target name, data and permissions
	// also call our toString function to convert the deck to string
	return ioutil.WriteFile(filename, []byte(p.toString()), 066)
}

func readPersonStore(filename string) personStore {
	// byteslice and error obj is returned from the function
	bs, err := ioutil.ReadFile(filename)
	// if error is there
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// need to resplit the string by comma, use std package
	s := strings.Split(string(bs), ",")
	// convert back into deck type
	return personStore(s)
}

func createNewPersonStore() {}

func getSongOfTheWeek() {}

//randomiser function takes in person list and returns selection as string
// func pickRandomName(p personStore) string {}

func inputPeopleList(inputStringList string) []string {
	re := regexp.MustCompile(`(?s)\[(.*)\]`)           // create the regular expression
	m := re.FindAllStringSubmatch(inputStringList, -1) //
	masterPeople := []string{
		m[0][1],
	}
	fmt.Println(masterPeople)
	return masterPeople
}

func peopleList() []string {
	//maybe play around with array vs. slice? Master is array and slice is gone vs. ToGo
	masterPeople := []string{
		"U0123E4S1S7",
		"USER2",
		"USER3",
	}
	return masterPeople
}

// allow someone elese to input the array of people they want / change the people list
