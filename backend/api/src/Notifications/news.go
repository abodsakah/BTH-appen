// Package notifications provides notifications
package notifications

import (
	"errors"
	"fmt"
	"log"

	"github.com/abodsakah/BTH-appen/backend/api/src/DB"
	expo "github.com/noahhakansson/exponent-server-sdk-golang/sdk"
	"gorm.io/gorm"
)

// SendNewsPushMessage function
func SendNewsPushMessage(gormDB *gorm.DB, news []db.News) error {
	var pushMessages []expo.PushMessage
	// create all push messages and append to slice
	for _, article := range news {
		pushMsg, err := createNewsPushMessage(gormDB, article)
		if err != nil {
			continue
		}
		pushMessages = append(pushMessages, pushMsg)
	}
	// return error if there are new messages to send
	if len(pushMessages) < 1 {
		return errors.New("Notifications; no news push messages to send")
	}
	// send messages
	var tryNumber uint = 1
	err := sendExpoPushMessages(pushMessages, tryNumber)
	if err != nil {
		return err
	}
	return nil
}

// createNewsPushMessages function
// Makes a `expo.PushMessage` for a news object and
// send it to all users that have expo push tokens.
//
// Return slice of `expo.PushMessage`s
// Or an error if no pushMessages were created.
func createNewsPushMessage(gormDB *gorm.DB, news db.News) (expo.PushMessage, error) {
	// create message to send to all tokens in pushTokens
	tokens, err := db.GetAllUserTokens(gormDB)
	if err != nil {
		return expo.PushMessage{}, err
	}

	// create expo.expoPushTokens
	var expoPushTokens []expo.ExponentPushToken
	for _, token := range tokens {
		pushToken, err := expo.NewExponentPushToken(token.ExpoPushToken)
		if err != nil {
			log.Println(err)
		} else {
			expoPushTokens = append(expoPushTokens, pushToken)
		}
	}

	// create expo.PushMessage
	data := map[string]string{
		"news_id": fmt.Sprint(news.ID),
		"link":    news.Link,
	}
	pushMsg := expo.PushMessage{
		To:        expoPushTokens,
		Title:     "BTH - News article published!",
		Body:      news.Title,
		Data:      data,
		Sound:     "default",
		Priority:  expo.DefaultPriority,
		ChannelID: "news",
	}

	return pushMsg, nil
}
