// Package notifications provides notifications
package notifications

import (
	// "log"
	// "time"

	// "github.com/abodsakah/BTH-appen/backend/api/src/DB"
	// expo "github.com/noahhakansson/exponent-server-sdk-golang/sdk"
	"gorm.io/gorm"
)

// TODO: Go through each exam in exams slice and
// create a slice of expo.PushMessage structs
// each being a message for one exam that should be sent to all registered users of that exam.

// startNewsServerfunction
//
// Starts notification server.
// Loops and sends message for due exams once a day.
// since GetExamsDueSoon only gets exams due in ONE and FIVE days,
// no duplicate notifications should be sent
// as they will not be due in ONE or FIVE days after one more day has passed
func startNewsServer(gormDB *gorm.DB, StopRunning *bool) error {
	// loop runs every few hours checking if any new news articles have been posted.
	// If a new article has been added a notification is sent to all users.
	for !(*StopRunning) {
		// TODO: loop logic
	}
	return nil
}
