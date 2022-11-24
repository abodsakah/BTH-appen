// Package notifications provides notifications
package notifications

import (
	"fmt"
	"log"
	"time"

	"github.com/abodsakah/BTH-appen/backend/api/src/DB"
	expo "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
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

// StartServer function
//
// Starts notification server.
func StartServer(gormObj *gorm.DB) error {
	// setup GORM database object
	gormDB := gormObj

	exams, err := db.GetExamsDueSoon(gormDB)
	if err != nil {
		return err
	}
	fmt.Println("Exams due soon: ", exams)
	examSendPushMessages([]expo.PushMessage{}, 1)

	// loop runs once every 24 hours, exits if StopRunning is set to true
	for !StopRunning {
		// 	// examSendPushMessages()
	}

	return nil
}

// getExpoPushTokens
// validates and returns all expo tokens for one exam.
func getExpoPushTokens(exam db.Exam) []expo.ExponentPushToken {
	// To check the token is valid
	var expoPushTokens []expo.ExponentPushToken
	pushToken, err := expo.NewExponentPushToken("ExponentPushToken[xxxxxxxxxxxxxxxxxxxxxx]")
	if err != nil {
		log.Println(err)
	}
	expoPushTokens = append(expoPushTokens, pushToken)
	return expoPushTokens
}

// examSendPushMessages function
//
// Loops and sends message for due exams once a day.
// since GetExamsDueSoon only gets exams due in ONE and FIVE days,
// no duplicate notifications should be sent
// as they will not be due in ONE or FIVE days after one more day has passed
func examSendPushMessages(messages []expo.PushMessage, tryNumber uint) {
	// Create a new Expo SDK client
	client := expo.NewPushClient(nil)
	msg := &expo.PushMessage{
		To:       []expo.ExponentPushToken{"ExponentPushToken[xxxxxxxxxxxxxxxxxxxxxx]"},
		Body:     "This is a test notification",
		Data:     map[string]string{"withSome": "data"},
		Sound:    "default",
		Title:    "Notification Title",
		Priority: expo.DefaultPriority,
	}
	messages = append(messages, *msg)
	messages = append(messages, *msg)

	// Publish message
	responses, err := client.PublishMultiple(messages)
	// Check errors
	if err != nil {
		log.Println(err, "\nAn error occured sending message, sleeping a little bit before retry")
		retrySendingMessages(messages, tryNumber)
		return
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
		retrySendingMessages(failedMessages, tryNumber)
	}
	// sleep for 24 hours if everything was successfully sent
	log.Println("Expo response:", responses)
	time.Sleep(time.Hour * 24)
}

// helper function to retry sending messages
//
// sleep for tryNumber of minutes, and then try to send messages.
// if tryNumber is >= 5; Return without retrying.
func retrySendingMessages(messages []expo.PushMessage, tryNumber uint) {
	if tryNumber >= maxRetries {
		return
	}
	log.Println("Some messages failed to send, will retry in ", tryNumber, " minutes...")
	sleepTime := time.Minute * time.Duration(tryNumber)
	time.Sleep(sleepTime)
	examSendPushMessages(messages, tryNumber+1)
}
