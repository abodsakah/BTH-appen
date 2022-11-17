package db

import (
	"time"

	"gorm.io/gorm"
)

// CreateExam function
// Takes a Exam struct and creates a database entry in exams table.
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

// ListExams function
// returns array with all exams from database or an error
func ListExams(db *gorm.DB) (exams []Exam, err error) {
	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	err = db.Where("start_date >= ?", midnight).Find(&exams).Error
	if err != nil {
		return nil, err
	}

	return exams, nil
}

// SearchExams function
// returns matching exams from database or an error
func SearchExams(db *gorm.DB, wildcard string) (exams []Exam, err error) {
	now := time.Now()
	err = db.Where("course_code LIKE ? AND start_date >= ?", wildcard, now).Find(&exams).Error
	if err != nil {
		return nil, err
	}

	return exams, nil
}

// RegisterToExam function
// Adds a user to an exams list of users, returns error if not
func RegisterToExam(db *gorm.DB, courseCode string, userID uint) error {
	// find exam
	var exam Exam
	result := db.Where("course_code = ?", courseCode).First(&exam)
	if result.Error != nil {
		return result.Error
	}
	// find user
	var user User
	err := db.Where("id = ?", userID).First(&user).Error
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

// UnregisterFromExam function
// Removes a user from an exams list of users, returns error if not
func UnregisterFromExam(db *gorm.DB, courseCode string, userID uint) error {
	// find exam
	var exam Exam
	result := db.Where("course_code = ?", courseCode).First(&exam)
	if result.Error != nil {
		return result.Error
	}
	// find user
	var user User
	err := db.Where("id = ?", userID).First(&user).Error
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
