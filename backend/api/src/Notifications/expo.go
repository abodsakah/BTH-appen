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

// StopRunning variable
// Allows stopping the expo notification server gracefully by setting it to true.
var StopRunning = false

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
	examSendPushMessages([]expo.PushMessage{})

	// loop runs once every 24 hours, exits if StopRunning is set to true
	for !StopRunning {
		// 	// examSendPushMessages()
	}

	return nil
}

// creates expo.PushMessage slice with all messages and returns it.
func createPushMessages(exams []db.Exam) {}

// examSendPushMessages function
//
// Loops and sends message for due exams once a day.
// since GetExamsDueSoon only gets exams due in ONE and FIVE days,
// no duplicate notifications should be sent
// as they will not be due in ONE or FIVE days after one more day has passed
func examSendPushMessages(messages []expo.PushMessage) {
	// loop
	for {
		// To check the token is valid
		pushToken, err := expo.NewExponentPushToken("ExponentPushToken[xxxxxxxxxxxxxxxxxxxxxx]")
		if err != nil {
			panic(err)
		}

		// Create a new Expo SDK client
		client := expo.NewPushClient(nil)

		// Publish message
		response, err := client.Publish(
			&expo.PushMessage{
				To:       []expo.ExponentPushToken{pushToken},
				Body:     "This is a test notification",
				Data:     map[string]string{"withSome": "data"},
				Sound:    "default",
				Title:    "Notification Title",
				Priority: expo.DefaultPriority,
			},
		)
		// Check errors
		if err != nil {
			log.Println(err, "\nAn error occured sending message, sleeping a little bit before retry")
			time.Sleep(time.Minute * 1)
			continue // jump to top of loop
		}

		// Validate responses
		if response.ValidateResponse() != nil {
			log.Println(response.PushMessage.To, "failed, sleeping a little bit before retry")
			time.Sleep(time.Minute * 1)
			continue // jump to top of loop
		}
		// sleep for 24 hours if everything was successfully sent
		log.Println("Expo response:", response)
		time.Sleep(time.Hour * 24)
	}
}
