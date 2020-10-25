package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// CustomError to be raised when regex doesn't work
type CustomError struct{}

func (m *CustomError) Error() string {
	return "Custom Error Occured"
}

type personStore []string

//Create and save new master list
func createNewPersonStore(message string) string {
	personStoreList, err := inputPeopleList(message) // first clean the message - see if in right format
	if err != nil {
		fmt.Println("Can't add new band")
		return "`Error:` can not add new band, please ensure format follows - `@SongBot newband [@user1, @user2]`"
	}
	personStoreList.saveToFile("masterList")
	fmt.Println("Successfully added new band")
	return "Successfully added new band"
}

func inputPeopleList(inputStringList string) (personStore, error) {
	re := regexp.MustCompile(`(?s)\[(.*)\]`)               // create the regular expression
	match := re.FindAllStringSubmatch(inputStringList, -1) //test the regular expression to see if throws errors
	if match == nil {                                      ///if there is no match (nil) then return error
		return nil, &CustomError{}
	}

	m := re.FindAllStringSubmatch(inputStringList, -1) //use the regular expression
	masterPeople := []string{
		m[0][1],
	}
	return masterPeople, nil // return the list of people to be stored
}

func (p personStore) toString() string {
	//use a go std library: strings, input and then separator
	//first turn deck d into a string
	return strings.Join([]string(p), ",")

}
func (p personStore) saveToFile(filename string) error {
	//using the writefile from ioutil library, takes target name, data and permissions
	// also call our toString function to convert the deck to string
	return ioutil.WriteFile(filename, []byte(p.toString()), 0666)
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

func peopleList() []string {
	//maybe play around with array vs. slice? Master is array and slice is gone vs. ToGo
	masterPeople := []string{
		"U0123E4S1S7",
		"USER2",
		"USER3",
	}
	return masterPeople
}
