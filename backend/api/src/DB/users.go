package db

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func checkInputLength(username string, password string) (err error) {
	if len(username) > userMaxLength || len(password) > passMaxLength {
		return errors.New("Error: username or password exceeds max length")
	}
	return nil
}

// CreateUser function
func CreateUser(db *gorm.DB, user *User) error {
	// check username and password length
	err := checkInputLength(user.Username, user.Password)
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
	err = db.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

// AuthUser function
func AuthUser(db *gorm.DB, username string, password string) (userID uint, err error) {
	// check username and password length
	err = checkInputLength(username, password)
	if err != nil {
		return 0, err
	}

	// get user from database
	var user User
	err = db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return 0, err
	}

	// compare password and hashedPassword
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return 0, err // return if password doesn't match
	}

	// return user ID
	return user.ID, nil
}
