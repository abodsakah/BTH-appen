// Package notifications provides notifications
package notifications

import (
// "log"
// "time"

// "github.com/abodsakah/BTH-appen/backend/api/src/DB"
// expo "github.com/noahhakansson/exponent-server-sdk-golang/sdk"
// db "github.com/abodsakah/BTH-appen/backend/api/src/DB"
// "gorm.io/gorm"
)

// TODO: Figure out way to see if new news articles have been added to database.
// Maybe a boolean `new` flag on each news article that gets set to false after notifications have been sent about it.
// Would result in a function that only gets news with the `new` flag set to true,
// which then sets the flag to false after getting it from the database.
// Loops every 3-5 hours to check for new news.

// createNewsPushMessages function
// Makes a `expo.PushMessage` for each news and
// calls getUserExpoPushTokens for each ALL users to get the `expo.ExponentPushToken`s
// that the `expo.PushMessage` should be sent to.
//
// Return slice of `expo.PushMessage`s
// Or an error if no pushMessages were created.
// func createNewsPushMessages(exams []db.News) ([]expo.PushMessage, error) {
// 	var pushMessages []expo.PushMessage
// 	for _, exam := range exams {
// 		pushTokens, err := getExamPushTokens(exam)
// 		// error means no push tokens for this exam.
// 		if err != nil {
// 			continue // jump to start of loop, skip creating message for this exam
// 		}
// 		// create message to send to all tokens in pushTokens
// 		pushMsg := expo.PushMessage{
// 			To:        pushTokens,
// 			Title:     fmt.Sprintf("%s: %s", exam.CourseCode, exam.Name),
// 			Body:      createDateTimeString(exam),
// 			Sound:     "default",
// 			Priority:  expo.DefaultPriority,
// 			ChannelID: "exams",
// 		}
// 		pushMessages = append(pushMessages, pushMsg)
// 	}
// 	if pushMessages != nil {
// 		return pushMessages, nil
// 	}
// 	return nil, errors.New("Error: no pushMessages to be sent")
// }
