package db

import (
	"time"

	"gorm.io/gorm"
)

// CreateExam function
// Takes a Exam struct and creates a database entry in exams table.
//
// Or an error.
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
// Takes a course_code and deletes the exam from the database.
//
// Or an error.
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
// Or an error.
func ListExams(db *gorm.DB) (exams []Exam, err error) {
	err = db.Find(&exams).Error
	if err != nil {
		return nil, err
	}

	return exams, nil
}

// GetExamsDueSoon function
//
// Returns all exams due in FIVE or ONE days
// with their users preloaded from the database.
// Made with intent to get exams with users to notify them in the app.
//
// Or an error.
func GetExamsDueSoon(db *gorm.DB) (exams []Exam, err error) {
	return nil, nil
}

// SearchExams function
// Returns matching exams from the database.
//
// Or an error.
func SearchExams(db *gorm.DB, wildcard string) (exams []Exam, err error) {
	now := time.Now()
	err = db.Where("course_code LIKE ? AND start_date >= ?", wildcard, now).Find(&exams).Error
	if err != nil {
		return nil, err
	}

	return exams, nil
}

// RegisterToExam function
// Adds a user to an exams list of users.
//
// Or an error.
func RegisterToExam(db *gorm.DB, examID uint, userID uint) error {
	// find exam
	var exam Exam
	result := db.Where("id = ?", examID).First(&exam)
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
// Removes a user from an exams list of users.
//
// Or an error.
func UnregisterFromExam(db *gorm.DB, examID uint, userID uint) error {
	// find exam
	var exam Exam
	result := db.Where("id = ?", examID).First(&exam)
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
