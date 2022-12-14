package db

import (
	"errors"
	"time"
  models "github.com/abodsakah/BTH-appen/backend/api/src/Models"
	"gorm.io/gorm"
)

// CreateExam function
// Takes a Exam struct and creates a database entry in exams table.
//
// Or returns an error.
func CreateExam(db *gorm.DB, exam *models.Exam) error {
	// set creation date
	exam.CreatedAt = time.Now()

	// create exam in database
	err := db.Create(&exam).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteExam function
// Takes a exam ID and deletes the exam from the database.
//
// Or returns an error.
func DeleteExam(db *gorm.DB, examID uint) (models.Exam, error) {
	// find exam
	exam := models.Exam{}
	err := db.Where("id = ?", examID).First(&exam).Error
	if err != nil {
		return models.Exam{}, err
	}
	// delete exam from database
	err = db.Delete(&exam).Error
	if err != nil {
		return models.Exam{}, err
	}

	return exam, nil
}

// GetExams function
// Returns an array with all exams
//
// Or returns an error.
func GetExams(db *gorm.DB) ([]models.Exam, error) {
	var exams []models.Exam
	err := db.Order("start_date ASC").Find(&exams).Error
	if err != nil {
		return nil, err
	}

	return exams, nil
}

// GetExamUsers function
// Returns an array with all users registered to an exam.
// Takes an exam ID.
//
// Or returns an error.
func GetExamUsers(db *gorm.DB, examID uint) ([]*models.User, error) {
	var exam models.Exam

	// get exam with users preloaded
	err := db.Where("id = ?", examID).Preload("Users").First(&exam).Error
	if err != nil {
		return nil, err
	}
	// return users
	return exam.Users, nil
}

// GetExamsDueSoon function
//
// Returns all exams due in FIVE or ONE days
// with all registered users ExpoPushTokens
// Made with intent to get exams with users to notify them in the app.
//
// Or returns an error.
func GetExamsDueSoon(db *gorm.DB) ([]models.Exam, error) {
	var exams []models.Exam
	t := time.Now()
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	tomorrow := midnight.AddDate(0, 0, 1)
	inTwoDays := midnight.AddDate(0, 0, 2)
	inFiveDays := midnight.AddDate(0, 0, 5)
	inSixDays := midnight.AddDate(0, 0, 6)

	// get exams with users preloaded
	err := db.Order("start_date ASC").
		Where("start_date BETWEEN ? AND ?", tomorrow, inTwoDays).
		Or("start_date BETWEEN ? AND ?", inFiveDays, inSixDays).
		Preload("Users.Tokens").
		Find(&exams).Error
	if err != nil {
		return nil, err
	}
	return exams, nil
}

// AddUserToExam function
// Adds a user to an exams list of users.
//
// Or returns an error.
func AddUserToExam(db *gorm.DB, examID uint, userID uint) (models.User, error) {
	// find user
	var user models.User
	err := db.Omit("password").Where("id = ?", userID).Preload("Exams").First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	// check user is not already registered to exam
	var currentExam *models.Exam
	for _, e := range user.Exams {
		if e.ID == examID {
			currentExam = e
			break
		}
	}
	if currentExam != nil {
		return models.User{}, errors.New("User is already registered to this exam")
	}
	// find exam
	var exam models.Exam
	err = db.Where("id = ?", examID).First(&exam).Error
	if err != nil {
		return models.User{}, err
	}
	// update user exams with exam
	err = db.Model(&user).Association("Exams").Append([]models.Exam{exam})
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// RemoveUserFromExam function
// Removes a user from an exams list of users.
//
// Or returns an error.
func RemoveUserFromExam(db *gorm.DB, examID uint, userID uint) (models.User, error) {
	// find user
	var user models.User
	err := db.Omit("password").Where("id = ?", userID).Preload("Exams").First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	// find exam
	var exam *models.Exam
	for _, e := range user.Exams {
		if e.ID == examID {
			exam = e
			break
		}
	}
	if exam == nil {
		return models.User{}, gorm.ErrRecordNotFound
	}
	// remove exam from user
	err = db.Model(&user).Association("Exams").Delete([]models.Exam{*exam})
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// GetUserExams function
// Lists the users exams which they have registered to
//
// Or an error
func GetUserExams(db *gorm.DB, userID uint) ([]*models.Exam, error) {
	var user models.User
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}
	err = db.Model(&models.User{}).Preload("Exams").Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user.Exams, nil
}
