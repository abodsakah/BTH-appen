package db

import "time"

// CreateExam function
func CreateExam(exam *Exam) error {
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
// returns array with all exams from database
func ListExams() (exams []Exam, err error) {
	now := time.Now()
	result := db.Where("Startdate >= ?", now).Find(&exams)
	if result.Error != nil {
		return nil, result.Error
	}

	return exams, nil
}

// SearchExams function
// returns a single exam object from database
func SearchExams(wildcard string) (exams []Exam, err error) {
	now := time.Now()
	result := db.Where("CourseCode LIKE ? AND StartDate >= ?", wildcard, now).Find(&exams)
	if result.Error != nil {
		return nil, result.Error
	}

	return exams, nil
}

// ApplyExam function
// returns none if successful, returns error if not
func ApplyExam(courseCode string, userEmail string) error {
	var exam *Exam
	result := db.Where("CourseCode = ?", courseCode).Find(&exam).Limit(1)
	if result.Error != nil {
		return result.Error
	}
	var user *User
	result = db.Where("Username = ?", userEmail).Find(&user).Limit(1)
	if result.Error != nil {
		return result.Error
	}
	exam.Users = append(exam.Users, *user)
	db.Save(&exam)
	return nil
}
