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

// startExamServer function
//
// Starts exam notification server.
// Loops and sends message for due exams once a day.
// Failed messages are retried `maxRetries` times.
//
// Since GetExamsDueSoon only gets exams due in ONE and FIVE days,
// no duplicate notifications should be sent
// as they will not be due in ONE or FIVE days after one more day has passed
func startExamServer(gormDB *gorm.DB, StopRunning *bool) error {
	// loop runs once every 24 hours, exits if StopRunning is set to true
	for !(*StopRunning) {
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
		// otherwise create messages from the fetched exams
		pushMessages, err := createExamPushMessages(exams)
		if err != nil {
			log.Println("No messages to send, retrying in 24 hours...")
			time.Sleep(time.Hour * 24)
			continue
		}
		sendExpoPushMessages(pushMessages, 1)
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

// createExamPushMessages function
// Makes a `expo.PushMessage` for each exam and
// calls getExpoPushTokens for each exam to get the `expo.ExponentPushToken`s
// that the `expo.PushMessage` should be sent to.
//
// Return slice of `expo.PushMessage`s
// Or an error if no pushMessages were created.
func createExamPushMessages(exams []db.Exam) ([]expo.PushMessage, error) {
	var pushMessages []expo.PushMessage
	for _, exam := range exams {
		pushTokens, err := getExamUsersPushTokens(exam)
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

// getExamUsersPushTokens function
// validates and returns all registered users `expo.ExponentPushToken`s for ONE exam.
//
// Returns an error if no user `expo.ExponentPushToken`s are found for the exam
func getExamUsersPushTokens(exam db.Exam) ([]expo.ExponentPushToken, error) {
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
