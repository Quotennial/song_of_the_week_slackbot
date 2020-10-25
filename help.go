package main

func helpString() string {
	helpMesg := ` Welcome to SongBot!
	- *newband*: takes in list of slack usernames to be used as master list for SOTW
	- *showband*: shows the master list
	- *showlineup*: shows who is next for song of the week (interim list)
	- *help*: prints help info
	
	Project found at - https://github.com/Quotennial/song_of_the_week_slackbot`

	return helpMesg
}
