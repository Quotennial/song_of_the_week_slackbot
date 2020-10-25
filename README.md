# SongBot - Song of the Week Selecter Tool

A tool to help pick a member of your team at random! It was originally intended to select a team memeber to be in charge of posting the Friday song of the week, but can be implemented for any weekly selection tool. Essentially picks a name from a list without replacement. There is also a command that picks name at random from whole origninal team (with teplacement).

## Set up
Put the bot online 

```go
go run main.go person_store.go choose_person.go help.go
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


# Building Info
Following [this video tutorial](https://www.youtube.com/watch?v=zkB_c3cgtd0&t=1186s)

Got stuck on bringing the bot online - so used this [gopher tutorial](https://blog.gopheracademy.com/advent-2017/go-slackbot/)

GO documentation https://godoc.org/github.com/nlopes/slack#AuthTestResponse

Also quite useful https://medium.com/mercari-engineering/writing-an-interactive-message-bot-for-slack-in-golang-6337d04f36b9

Github implementation https://github.com/shomali11/slacker/blob/master/response.go

To test the bot https://app.slack.com/client/T011WFT94GN/C0121S6UU9K 


## Requirements and Roadmap
**Functional Requirements**
- FR1: can add list of users to be selected from (rejects incorrect format) :heavy_check_mark:
- FR2: can edit the list of users to be selected from :heavy_check_mark:
- FR3: has *help* info sheet with list available commands (info) :heavy_check_mark:
- FR4: can select name from list at random (with replacement) :heavy_check_mark:
- FR5: can select "song of the week" person from list without replacement 
- FR6: user is able to accept/ reject name 
- FR6: command to save song of the week to DB (maybe easier to search slack?)
- FR7: have multiple channels/users per slack account (possible with naming_lists by channel?)


**Non-Functional Requirements**
- Image for SongBot
- Have different phrases for song of the week (pick from them at random)
- Testing :heavy_check_mark: (scafolding and basic test built)