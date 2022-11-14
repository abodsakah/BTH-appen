package db

import (
	"time"

	"gorm.io/gorm"
)

// CreateExam function
func CreateExam(db *gorm.DB, exam *Exam) error {
	// set creation date
	exam.CreatedAt = time.Now()

	// create exam in database
	result := db.Create(&exam)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ListExams function
// returns array with all exams from database or an error
func ListExams(db *gorm.DB) (exams []Exam, err error) {
	now := time.Now()
	result := db.Where("start_date >= ?", now).Find(&exams)
	if result.Error != nil {
		return nil, result.Error
	}

	return exams, nil
}

// SearchExams function
// returns matching exams from database or an error
func SearchExams(db *gorm.DB, wildcard string) (exams []Exam, err error) {
	now := time.Now()
	result := db.Where("course_code LIKE ? AND start_date >= ?", wildcard, now).Find(&exams)
	if result.Error != nil {
		return nil, result.Error
	}

	return exams, nil
}

// ApplyExam function
// adds a user to an exams list of users, returns error if not
func ApplyExam(db *gorm.DB, courseCode string, userEmail string) error {
	var exam Exam
	result := db.Where("course_code = ?", courseCode).Find(&exam).Limit(1)
	if result.Error != nil {
		return result.Error
	}
	var user User
	result = db.Where("username = ?", userEmail).Find(&user).Limit(1)
	if result.Error != nil {
		return result.Error
	}
	exam.Users = append(exam.Users, &user)
	db.Save(&exam)
	return nil
}
