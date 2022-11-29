// Package notifications provides notifications
package notifications

import (
	// "log"
	// "time"

	// "github.com/abodsakah/BTH-appen/backend/api/src/DB"
	// expo "github.com/noahhakansson/exponent-server-sdk-golang/sdk"
	"gorm.io/gorm"
)

// TODO: Figure out way to see if new news articles have been added to database.
// Maybe a boolean `new` flag on each news article that gets set to false after notifications have been sent about it.
// Would result in a function that only gets news with the `new` flag set to true,
// which then sets the flag to false after getting it from the database.
// Loops every 3-5 hours to check for new news.

// startNewsServerfunction
//
// Starts notification server.
// Loops and sends message for due exams once a day.
// since GetExamsDueSoon only gets exams due in ONE and FIVE days,
// no duplicate notifications should be sent
// as they will not be due in ONE or FIVE days after one more day has passed
func startNewsServer(gormDB *gorm.DB, stopRunning *bool) error {
	// loop runs every few hours checking if any new news articles have been posted.
	// If a new article has been added a notification is sent to all users.
	for !(*stopRunning) {
		// TODO: loop logic
	}
	return nil
}
