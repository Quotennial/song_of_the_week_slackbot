package main

import "fmt"

func songOfTheWeekSelection(listFilePath string) (reply string) {
	restartList := listNeedsReset(listFilePath)

	if restartList == true {
		fmt.Println("need to reset list")
		err := resetSOTWList(listFilePath) // attempt to make the master list SOTW list
		if err != nil {
			return "`Error:` couldn't find the Master List, please add band members" // error if there is no master list
		}
	}

	sotwList, _ := readPersonStore(listFilePath + "interim")

	sotwSelection := sotwList.randomFromList()
	return "SOTW selection is" + sotwSelection

}

func resetSOTWList(listFilePath string) error {
	// todo - what if there is also no master list??? - NEED TO CATCH THAT
	sotwList, err := readPersonStore(listFilePath + "master")
	if err != nil {
		return &CustomError{}
	}
	sotwList.saveToFile(listFilePath + "interim")
	return nil
}

func listNeedsReset(listFilePath string) bool {
	sotwList, err := readPersonStore(listFilePath + "interim")
	if err != nil {
		return true
	}
	if len(sotwList) == 0 { // check that list is empty
		return true
	}
	return false

}
