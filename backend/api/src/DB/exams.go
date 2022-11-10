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
	result := db.Find(&exams)
	if result.Error != nil {
		return nil, result.Error
	}

	return exams, nil
}
