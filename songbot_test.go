package main

import (
	"os"
	"testing"
)

// TestAddNewBandCorrectly add new list of names and check they can be read, covering newband and showband and pickrandom
func TestAddNewBandCorrectly(t *testing.T) {
	// set up vars to make new list
	channelNameTest := "testingChannel"
	listFolderName := "nameListFolder"                     // store the files with names in one folder
	listFilePath := listFolderName + "/" + channelNameTest // use this file path

	msg := "@SongBot newband [@testuser1, @testuser2, @testuser3]"
	if createNewPersonStore(msg, listFilePath) != "Successfully added new band" {
		t.Errorf("Band couldn't be added")
	}
	if showBand(listFilePath) != "The master list: [@testuser1, @testuser2, @testuser3]" {
		t.Errorf("Band couldn't be read")
	}
	os.Remove(listFilePath + "masterList") //clean up and remove testing file
}

func TestAddNewBandWrong(t *testing.T) {
	// set up vars to make new list
	channelNameTest := "testingChannelIncorrect"
	listFolderName := "nameListFolder"                     // store the files with names in one folder
	listFilePath := listFolderName + "/" + channelNameTest // use this file path

	msg := "@SongBot newband @testuser1, @testuser2, @testuser3"
	expectedReply := "`Error:` can not add new band, please ensure format follows - `@SongBot newband [@user1, @user2]`"

	if createNewPersonStore(msg, listFilePath) != expectedReply {
		t.Errorf("Error Parsing Band: createNewPersonStore ")
	}
}

func TestQueryNoBand(t *testing.T) {
	// set up vars to make new list
	channelNameTest := "testingChannelNoBand"
	listFolderName := "nameListFolder"                     // store the files with names in one folder
	listFilePath := listFolderName + "/" + channelNameTest // use this file path

	// empty band string - should get errors
	if showBand(listFilePath) != "`Error:` couldn't find the Master List, please add band members" {
		t.Errorf("Incorrect Error Handling for no file: showband")
	}
	if masterPickRandomPerson(listFilePath) != "`Error:` couldn't find the Master List, please add band members" {
		t.Errorf("Incorrect Error Handling for no file: masterPickRandomPerson")
	}
}
