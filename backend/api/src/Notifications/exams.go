// Package notifications provides notifications
package notifications

import (
	"errors"
	"fmt"
	"log"
	"math"
	"time"

	"github.com/abodsakah/BTH-appen/backend/api/src/DB"
	expo "github.com/noahhakansson/exponent-server-sdk-golang/sdk"
	"gorm.io/gorm"
)

// TODO: Go through each exam in exams slice and
// create a slice of expo.PushMessage structs
// each being a message for one exam that should be sent to all registered users of that exam.

var (
	// StopRunning variable
	// Allows stopping the expo notification server gracefully by setting it to true.
	StopRunning = false
	// max times to retry sending messages before giving up.
	maxRetries uint = 5
)

// StartExamServer function
//
// Starts notification server.
// Loops and sends message for due exams once a day.
// since GetExamsDueSoon only gets exams due in ONE and FIVE days,
// no duplicate notifications should be sent
// as they will not be due in ONE or FIVE days after one more day has passed
func StartExamServer(gormObj *gorm.DB) error {
	// setup GORM database object
	gormDB := gormObj

	// test testMessages
	var testMessages []expo.PushMessage
	msg := &expo.PushMessage{
		To:       []expo.ExponentPushToken{"ExponentPushToken[xxxxxxxxxxxxxxxxxxxxxx]", "ExponentPushToken[ePYzdGJlxthQk6_M-HiOzJ]"},
		Title:    "Notification Title",
		Body:     "This is a test notification",
		Data:     map[string]string{"withSome": "data"},
		Sound:    "default",
		Priority: expo.DefaultPriority,
	}
	msg1 := &expo.PushMessage{
		To:       []expo.ExponentPushToken{"ExponentPushToken[ePYzdGJlxthQk6_M-HiOzJ]"},
		Title:    "Notification Title",
		Body:     "This is a test notification",
		Data:     map[string]string{"withSome": "data"},
		Sound:    "default",
		Priority: expo.DefaultPriority,
	}

	testMessages = append(testMessages, *msg1)
	testMessages = append(testMessages, *msg)
	// loop runs once every 24 hours, exits if StopRunning is set to true
	for !StopRunning {
		exams, err := db.GetExamsDueSoon(gormDB)
		if err != nil {
			log.Println(err)
			log.Println("Failed getting exams, retrying in 1 hour...")
			time.Sleep(time.Hour * 1)
			continue // error; jump to top of loop, and sleep for 24hrs
		}
		// if exams array is empty, nothing to send.
		// sleep and retry tommorow
		if len(exams) <= 0 {
			log.Println("No exams due soon, retrying in 24 hours...")
			time.Sleep(time.Hour * 24)
			continue // jump to top of loop
		}
		fmt.Println("Exams due soon: ", exams)
		pushMessages, err := createExpoPushMessages(exams)
		if err != nil {
			log.Println("No messages to send, retrying in 24 hours...")
			time.Sleep(time.Hour * 24)
			continue
		}
		examSendPushMessages(pushMessages, 1)
	}

	return nil
}

// creates and return a string that says
// in how many days the exam is starting and what date and time.
func createDateTimeString(exam db.Exam) string {
	days := int(math.Round(time.Until(exam.StartDate).Hours() / 24))
	_, startMonth, startDay := exam.StartDate.Date()
	startHour, startMin, _ := exam.StartDate.Clock()
	startTime := fmt.Sprintf("%d:%d", startHour, startMin)
	if days > 1 {
		return fmt.Sprintf("Exam starts in %d days, %d %s at %s", days, startDay, startMonth.String(), startTime)
	}
	return fmt.Sprintf("Exam starts in %d day, %d %s at %s", days, startDay, startMonth.String(), startTime)
}

// createExpoPushMessages function
// Makes a `expo.PushMessage` for each exam and
// calls getExpoPushTokens for each exam to get the `expo.ExponentPushToken`s
// that the `expo.PushMessage` should be sent to.
//
// Return slice of `expo.PushMessage`s
// Or an error if no pushMessages were created.
func createExpoPushMessages(exams []db.Exam) ([]expo.PushMessage, error) {
	var pushMessages []expo.PushMessage
	for _, exam := range exams {
		pushTokens, err := getExpoPushTokens(exam)
		// error means no push tokens for this exam.
		if err != nil {
			continue // jump to start of loop, skip creating message for this exam
		}
		// create message to send to all tokens in pushTokens
		pushMsg := expo.PushMessage{
			To:        pushTokens,
			Title:     fmt.Sprintf("%s: %s", exam.CourseCode, exam.Name),
			Body:      createDateTimeString(exam),
			Sound:     "default",
			Priority:  expo.DefaultPriority,
			ChannelID: "exams",
		}
		pushMessages = append(pushMessages, pushMsg)
	}
	if pushMessages != nil {
		return pushMessages, nil
	}
	return nil, errors.New("Error: no pushMessages to be sent")
}

// getExpoPushTokens function
// validates and returns all `expo.ExponentPushToken`s for ONE exam.
//
// Returns an error if no user `expo.ExponentPushToken`s are found for the exam
func getExpoPushTokens(exam db.Exam) ([]expo.ExponentPushToken, error) {
	// To check the token is valid
	var expoPushTokens []expo.ExponentPushToken
	// for each users tokens, validate and add them to expoPushTokens
	for _, user := range exam.Users {
		for _, token := range user.Tokens {
			pushToken, err := expo.NewExponentPushToken(token.ExpoPushToken)
			if err != nil {
				log.Println(err)
			} else {
				expoPushTokens = append(expoPushTokens, pushToken)
			}
		}
	}
	if expoPushTokens != nil {
		return expoPushTokens, nil
	}
	return nil, errors.New("Error: No tokens found for exam")
}

// examSendPushMessages function
//
func examSendPushMessages(messages []expo.PushMessage, tryNumber uint) {
	// Create a new Expo SDK client
	client := expo.NewPushClient(nil)

	// Publish message
	responses, err := client.PublishMultiple(messages)
	// Check errors
	if err != nil {
		log.Println(err, "\nAn error occured sending message, sleeping for 1 hour before retrying.")
		time.Sleep(time.Hour * 1)
		retrySendingMessages(messages, tryNumber)
	}
	// Validate responses
	// save failed responses in array and call this function again.
	// when it returns
	var failedMessages []expo.PushMessage
	for _, response := range responses {
		if response.ValidateResponse() != nil {
			failedMessages = append(failedMessages, response.PushMessage)
		}
	}
	// retry sending failed messages, if any failed
	if failedMessages != nil {
		log.Println("we have some failed messages")
		for key, msg := range failedMessages {
			log.Printf("--------Failed message[%d]:\n%#v", key, msg)
		}
		retrySendingMessages(failedMessages, tryNumber)
	}
	// print response
	log.Println("Expo response:", responses)
}

// helper function to retry sending messages
//
// sleep for tryNumber of minutes, and then try to send messages.
// if tryNumber is >= 5; Return without retrying.
func retrySendingMessages(messages []expo.PushMessage, tryNumber uint) {
	if tryNumber > maxRetries {
		return
	}
	log.Println("Some messages failed to send, will retry in ", tryNumber, " minutes...")
	sleepTime := time.Minute * time.Duration(tryNumber)
	time.Sleep(sleepTime)
	examSendPushMessages(messages, tryNumber+1)
}
