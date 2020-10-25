package main

import "testing"

// TestAddNewBandCorrectly add new list of names and check they can be read, covering newband and showband and pickrandom
func TestAddNewBandCorrectly(t *testing.T) {
	// set up vars to make new list
	channelNameTest := "testingChannel"
	listFolderName := "nameListFolder"                     // store the files with names in one folder
	listFilePath := listFolderName + "/" + channelNameTest // use this file path

	// add incorrect/ empty band string - should get errors
	showBand(listFilePath)
	pickRandomPerson(listFilePath)

	// correctly add members
	msg := "@SongBot newband [@testuser1, @testuser2, @testuser3]"
	createNewPersonStore(msg, listFilePath)
	showBand(listFilePath)
	pickRandomPerson(listFilePath)

	// add new band correctly and incorrectly
}
