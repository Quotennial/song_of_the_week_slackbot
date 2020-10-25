Following [this video tutorial](https://www.youtube.com/watch?v=zkB_c3cgtd0&t=1186s)

Got stuck on bringing the bot online - so used this [gopher tutorial](https://blog.gopheracademy.com/advent-2017/go-slackbot/)

GO documentation https://godoc.org/github.com/nlopes/slack#AuthTestResponse

Also quite useful https://medium.com/mercari-engineering/writing-an-interactive-message-bot-for-slack-in-golang-6337d04f36b9

Github implementation https://github.com/shomali11/slacker/blob/master/response.go

To test the bot https://app.slack.com/client/T011WFT94GN/C0121S6UU9K 


## Set up

Put the bot online 

```go
go run main.go person_store.go choose_person.go 
```

To test the bot add it to your [workspace](https://app.slack.com/client/T011WFT94GN/C0121S6UU9K)

### Adding Arguments
`add_bandmembers` - this takes in a string (in slice format) and adds as the list of people to choose from. 

`todaydj` - picks random member from list-without-replacement
- `accept` - confrims the pick
- `reject` - replaces and starts again
- `cancel` - cancels operation

`djtogo` - prints list of people not yet picked
`random_all` - picks random list from all 

# Requirements
**Functional Requirements**
- FR1: periodically (weekly) post random name from list
- FR2: user is able to accept/ reject name from list and immedeatley regenerates
- FR3: list regenerates when finished 
- FR4: can be polled/ asked for name out of cycle
- FR5: extra functionality to pick random name from full list
- FR6: list available fucntions (info)
- FR6: command to save song of the week to DB (maybe easier to search slack?)
- FR7: reminders if haven't posted

**Non-Functional Requirements**

- Have different phrases for song of the week (pick at random from them)