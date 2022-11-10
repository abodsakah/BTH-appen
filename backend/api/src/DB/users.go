package db

import (
	"errors"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func checkInputLength(username string, password string) (err error) {
	if len(username) > userMaxLength || len(password) > passMaxLength {
		return errors.New("Error: username or password exceeds max length")
	}
	return nil
}

// CreateUser function
func CreateUser(user *User) error {
	// check username and password length
	err = checkInputLength(user.Username, user.Password)
	if err != nil {
		return err
	}

	// set creation date
	user.CreatedAt = time.Now()

	// hash password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPass)

	// create user in database
	result := db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// AuthUser function
func AuthUser(username string, password string) (userID string, err error) {
	// check username and password length
	err = checkInputLength(username, password)
	if err != nil {
		return "", err
	}

	// get user from database
	var user User
	result := db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return "", result.Error
	}

	// compare password and hashedPassword
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err // return if password doesn't match
	}

	// return user ID
	userID = strconv.Itoa(int(user.ID))
	return userID, nil
}