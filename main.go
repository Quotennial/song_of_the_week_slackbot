package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

const (
	// action is used for slack attament action.
	actionSelect = "select"
	actionStart  = "start"
	actionCancel = "cancel"
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

	// Register handler to receive interactive message
	// responses from slack (kicked by user action)

	for msg := range rtm.IncomingEvents { //for all incoming messages
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
				reply := masterPickRandomPerson(listFilePath)
				replyBasic(ev, reply)
				continue
			}

			if strings.Contains(ev.Msg.Text, "help") {
				reply := helpString()
				replyBasic(ev, reply)
				continue
			}

			if strings.Contains(ev.Msg.Text, "sotw") {
				reply := songOfTheWeekSelection(listFilePath)
				replySOTW(ev, reply, slackClient)
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

func replySOTW(ev *slack.MessageEvent, replySOTW string, slackClient *slack.Client) {

	// value is passed to message handler when request is approved.
	attachment := slack.Attachment{
		Text:       "Would you like to share your song of the week? :musical_note: :control_knobs: :headphones:",
		CallbackID: fmt.Sprintf("ask_%s", ev.User),
		Color:      "#666666",
		Actions: []slack.AttachmentAction{
			slack.AttachmentAction{
				Name:  "action",
				Text:  "No thanks!",
				Type:  "button",
				Value: "no",
			},
			slack.AttachmentAction{
				Name:  "action",
				Text:  "Yes, please!",
				Type:  "button",
				Value: "yes",
			},
		},
	}
	fmt.Println(fmt.Sprintf("ask_%s", ev.User))

	// from tutoiral https://github.com/nlopes/slack/blob/master/examples/buttons/buttons.go
	message := slack.MsgOptionAttachments(attachment)
	sotwMessage := slack.MsgOptionText("", false)
	channelID, timestamp, err := slackClient.PostMessage(ev.Channel, sotwMessage, message)
	if err != nil {
		fmt.Printf("Could not send message: %v", err)
	}
	fmt.Printf("Message with buttons sucessfully sent to channel %s at %s", channelID, timestamp)
	http.HandleFunc("/actions", actionHandler)
	http.ListenAndServe(":3000", nil)

}

func actionHandler(w http.ResponseWriter, r *http.Request) {
	var payload slack.InteractionCallback
	err := json.Unmarshal([]byte(r.FormValue("payload")), &payload)
	if err != nil {
		fmt.Printf("Could not parse action response JSON: %v", err)
	}
	fmt.Printf("Message button pressed by user %s with value %s", payload.User.Name, payload.Value)
}
