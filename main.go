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
			// ***************INTERACTION DECISION TREE**************
			// add arguments to these if statements, return string, then put this string into the reply function

			// takes in string of slack usernames and stores in file
			if strings.Contains(ev.Msg.Text, "newband") {
				reply := createNewPersonStore(ev.Msg.Text)
				replyBasic(ev, reply)
				fmt.Println("New Band List")
			}

			if strings.Contains(ev.Msg.Text, "showband") {
				masterList, err := readPersonStore("masterlist")
				if err != nil {
					replyBasic(ev, "I couldn't find the master list")
					continue
				}
				reply := "[" + masterList.toString() + "]"
				replyBasic(ev, reply)
			}

			if strings.Contains(ev.Msg.Text, "pickrandom") {
				pickedPerson := pickRandomPerson()
				replyBasic(ev, pickedPerson)
			}

			if strings.Contains(ev.Msg.Text, "showlist") {
				fmt.Println("show list")
				continue
			}

			if strings.Contains(ev.Msg.Text, "help") {
				reply := helpString()
				replyBasic(ev, reply)
			}

			replyToUser(ev) // need a channeling function to filter out what is being asked
		}
	}
}

func replyToUser(ev *slack.MessageEvent) { //change this to the channel for the actual app
	fmt.Printf("Channel: %v/n", ev.Msg.Channel)
	// candidatePeople := peopleList() // make this master list and can be the generator function
	sentence := "test"

	msg := slack.MsgOptionText(sentence, false)
	channelID, timestamp, err := slackClient.PostMessage(ev.Channel, msg)
	// need to accept/ reject (maybe needs to be sepreate func- chack the beer tutorial)
	fmt.Println(channelID, timestamp)

	if err != nil {
		log.Println("error sending to slack: " + err.Error())
	}
	return

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
