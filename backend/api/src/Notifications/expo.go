// Package notifications provides notifications
package notifications

import (
	"log"
	"time"

	expo "github.com/noahhakansson/exponent-server-sdk-golang/sdk"
)

// sendExpoPushMessages function
//
func sendExpoPushMessages(messages []expo.PushMessage, tryNumber uint) {
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
	sendExpoPushMessages(messages, tryNumber+1)
}
