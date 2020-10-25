package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

var (
	slackClient *slack.Client //initialise the slack event
)

func main() {
	slackAccessToken := goDotEnvVariable("SLACK_ACCESS_TOKEN") // get the slack access token in .env
	slackClient = slack.New(slackAccessToken)
	rtm := slackClient.NewRTM() //create the realtime messaging objext
	go rtm.ManageConnection()   // set it up in a go routine

	for msg := range rtm.IncomingEvents { //for all incoming messages
		// fmt.Println(msg)
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if len(ev.User) == 0 {
				continue
			}

			// check if we have a DM, or standard channel post
			direct := strings.HasPrefix(ev.Msg.Channel, "D")
			fmt.Printf("Message Channel: %v/n", ev.Msg.Channel)
			authTest, _ := slackClient.AuthTest() //test the bot to get bot id
			// if it is not a direct message and not a mention - ignore
			if !direct && !strings.Contains(ev.Msg.Text, "@"+authTest.UserID) {
				fmt.Println("Message not for us")
				continue // does this break out of it? I think so
			}

			fmt.Println("Message for us")
			fmt.Println(ev.Msg.Text)
			fmt.Println(ev.Channel)
			// ***************INTERACTION DECISION TREE**************

			channelName := ev.Channel                          //save channel name to identify whic list to use
			listFolderName := "nameListFolder"                 // store the files with names in one folder
			listFilePath := listFolderName + "/" + channelName // use this file path

			// takes in string of slack usernames and stores in file
			if strings.Contains(ev.Msg.Text, "newband") {
				reply := createNewPersonStore(ev.Msg.Text, listFilePath)
				replyBasic(ev, reply)
				continue
			}

			// reads and shows the master list
			if strings.Contains(ev.Msg.Text, "showband") {
				reply := showBand(listFilePath)
				replyBasic(ev, reply)
				continue
			}

			if strings.Contains(ev.Msg.Text, "pickrandom") {
				pickedPerson, err := pickRandomPerson(listFilePath)
				if err != nil {
					replyBasic(ev, pickedPerson) // picked person carries the error message
					continue
				}
				reply := "Randomly selected person is: " + pickedPerson
				replyBasic(ev, reply)
			}

			if strings.Contains(ev.Msg.Text, "sotw") {
				continue
			}

			if strings.Contains(ev.Msg.Text, "help") {
				reply := helpString()
				replyBasic(ev, reply)
				continue
			}

		}
	}
}

func replyBasic(ev *slack.MessageEvent, replyString string) { //change this to the channel for the actual app
	msg := slack.MsgOptionText(replyString, false)
	channelID, timestamp, err := slackClient.PostMessage(ev.Channel, msg)
	// need to accept/ reject (maybe needs to be sepreate func- chack the beer tutorial)
	fmt.Println(channelID, timestamp)

	if err != nil {
		log.Println("error sending to slack: " + err.Error())
	}
	return

}
