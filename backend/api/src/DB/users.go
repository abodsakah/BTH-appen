package db

import (
	"errors"
	"time"

  models "github.com/abodsakah/BTH-appen/backend/api/src/Models"
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
func CreateUser(db *gorm.DB, user *models.User) error {
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

// IsRole function
//
// Tests if user has the given role
func IsRole(db *gorm.DB, id uint, role string) (bool, error) {
	var user models.User
	err := db.Where("id = ?", id).Find(&user).Error
	if err != nil || user.Role != role {
		return false, err
	}
	return true, nil
}

// GetAllUsers function
func GetAllUsers(db *gorm.DB) ([]models.User, error) {
	// get users from database
	var users []models.User

	err := db.Omit("password").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

// GetAllUserTokens function
// Wont get deleted users token.
// And later if implemented, should
// respect user settings incase they dont want notifications.
//
// Returns all users expo tokens.
func GetAllUserTokens(db *gorm.DB) ([]models.Token, error) {
	// get users from database
	var users []models.User
	var tokens []models.Token

	err := db.Preload("Tokens").Find(&users).Error
	if err != nil {
		return nil, err
	}

	// append all user tokens to one large tokens slice.
	for _, user := range users {
		tokens = append(tokens, user.Tokens...)
	}

	return tokens, nil
}

// GetUser function
func GetUser(db *gorm.DB, userID uint) (models.User, error) {
	// get user from database
	var user models.User

	err := db.Omit("password").Where("id = ?", userID).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// AuthUser function
func AuthUser(db *gorm.DB, username string, password string) (userID uint, err error) {
	// check username and password length
	err = checkInputLength(username, password)
	if err != nil {
		return 0, err
	}

	// get user from database
	var user models.User
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

// AddExpoPushToken function
// Adds a expo push token to the user.
// ExpoPushToken example: `ExponentPushToken[xxxxxxxxxxxxxxxxxxxxxx]`
//
// Or returns an error.
func AddExpoPushToken(db *gorm.DB, userID uint, pushToken string) (models.User, error) {
	// find user
	var user models.User
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	// update user exams with exam
	err = db.Model(&user).Association("Tokens").Append(&models.Token{ExpoPushToken: pushToken})
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
