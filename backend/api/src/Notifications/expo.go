// Package notifications provides notifications
package notifications

import (
	"errors"
	"log"
	"time"

	expo "github.com/noahhakansson/exponent-server-sdk-golang/sdk"
	"gorm.io/gorm"
)

var (
	maxRetries uint = 5
	// ErrorMaxRetry returned on reaching max retry limit
	ErrMaxRetry = errors.New("Error: Retry time out, reached max try number")
)

// StartServers function
// Starts the notification server's in the notifications package.
//
//	stopRunning *bool
// Allows stopping the expo notification server's gracefully by setting it to true.
//	retries uint
// Max times to retry sending messages before giving up.
func StartServers(gormDB *gorm.DB, stopRunning *bool, retries uint) error {
	var err error
	// set maxRetries
	maxRetries = retries
	// start exam notification server
	if err = startExamServer(gormDB, stopRunning); err != nil {
		log.Fatalln("Something went wrong with the Exam notification server;\n error: ", err)
	}
	return err
}

// sendExpoPushMessages function
// Retries `tryNumber` of times, with increasing wait times.
func sendExpoPushMessages(messages []expo.PushMessage, tryNumber uint) error {
	// Create a new Expo SDK client
	client := expo.NewPushClient(nil)

	// Publish message
	log.Println("tryNumber; ", tryNumber)
	responses, err := client.PublishMultiple(messages)
	// Check errors
	if err != nil {
		log.Println(err, "\nAn error occured sending message, sleeping for 1 hour before retrying.")
		time.Sleep(time.Hour * 1)
		err := retrySendingMessages(messages, tryNumber)
		if err != nil {
			return err
		}
		return nil
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
		log.Println("CALLING retrySendingMessages(); tryNumber: ", tryNumber)
		err := retrySendingMessages(failedMessages, tryNumber)
		if err != nil {
			return err
		}
		return nil
	}
	// print response
	log.Println("Expo response:", responses)
	return nil
}

// helper function to retry sending messages
//
// sleep for tryNumber of minutes, and then try to send messages.
//  example:
// if tryNumber is = 6 and maxRetries = 5; Return without retrying.
func retrySendingMessages(messages []expo.PushMessage, tryNumber uint) error {
	if tryNumber > maxRetries {
		return ErrMaxRetry
	}
	log.Println("Retry: ", tryNumber)
	log.Println("Some messages failed to send, will retry in ", tryNumber, " minutes...")
	sleepTime := time.Second * time.Duration(tryNumber)
	time.Sleep(sleepTime)
	err := sendExpoPushMessages(messages, tryNumber+1)
	if err != nil {
		log.Println(err)
	}
	return nil
}
