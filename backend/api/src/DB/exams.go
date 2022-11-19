package db

import (
	"time"

	"gorm.io/gorm"
)

// CreateExam function
// Takes a Exam struct and creates a database entry in exams table.
//
// Or returns an error.
func CreateExam(db *gorm.DB, exam *Exam) error {
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
func DeleteExam(db *gorm.DB, examID uint) error {
	// delete exam from database
	err := db.Delete(&Exam{}, examID).Error
	if err != nil {
		return err
	}

	return nil
}

// ListExams function
// Returns an array with all exams that have
// a start date after today at midnight from the database.
//
// Or returns an error.
func ListExams(db *gorm.DB) ([]Exam, error) {
	var exams []Exam
	err := db.Find(&exams).Error
	if err != nil {
		return nil, err
	}

	return exams, nil
}

// ListExamUsers function
// Returns an array with all users registered to an exam.
// Takes an exam ID.
//
// Or returns an error.
func ListExamUsers(db *gorm.DB, examID uint) ([]*User, error) {
	var exam Exam
	// get exam with users preloaded
	err := db.Where("id = ?", examID).Preload("Users").First(&exam).Error
	if err != nil {
		return nil, err
	}
	return exam.Users, nil
}

// GetExamsDueSoon function
//
// Returns all exams due in FIVE or ONE days
// with their users preloaded from the database.
// Made with intent to get exams with users to notify them in the app.
//
// Or returns an error.
func GetExamsDueSoon(db *gorm.DB) (exams []Exam, err error) {
	return nil, nil
}

// AddUserToExam function
// Adds a user to an exams list of users.
//
// Or returns an error.
func AddUserToExam(db *gorm.DB, examID uint, userID uint) error {
	// find exam
	var exam Exam
	err := db.Where("id = ?", examID).First(&exam).Error
	if err != nil {
		return err
	}
	// find user
	var user User
	err = db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return err
	}
	// update user exams with exam
	err = db.Model(&user).Association("Exams").Append([]Exam{exam})
	if err != nil {
		return err
	}
	return nil
}

// RemoveUserFromExam function
// Removes a user from an exams list of users.
//
// Or returns an error.
func RemoveUserFromExam(db *gorm.DB, examID uint, userID uint) error {
	// find exam
	var exam Exam
	err := db.Where("id = ?", examID).First(&exam).Error
	if err != nil {
		return err
	}
	// find user
	var user User
	err = db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		return err
	}
	// remove exam from user
	err = db.Model(&user).Association("Exams").Delete([]Exam{exam})
	if err != nil {
		return err
	}
	return nil
}
