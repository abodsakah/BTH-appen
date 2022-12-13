package db

import (
	"testing"
	"time"

	fixture "github.com/abodsakah/BTH-appen/backend/api/src/Fixture"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB
	additionalTables = []string{"exam_users"}
)

var testUser = &User{
	Name:     "Test Testsson",
	Username: "test",
	Password: "pass",
	Role:     "student",
}

var testExam = &Exam{
	Name:       "test",
	CourseCode: "pa121212",
	StartDate:  time.Now().AddDate(0, 0, 1),
}

var testNews = &News{
	Title:       "Test",
	Date:        time.Now(),
	Description: "A test",
	Link:        "test.com",
}

func createUserWrap() (uint, error) {
	temp := *testUser
	err := CreateUser(db, testUser)
	id := testUser.ID
	*testUser = temp
	if err != nil {
		return 0, err
	}
	return id, nil
}

func fixtureWrapUser(t *testing.T) (uint, error) {
	err := fixture.CleanUp(db, additionalTables, &User{}, &Exam{}, &News{}, &Token{})
	if err != nil {
		t.Fatal(err)
	}
	id, err := createUserWrap()
	return id, err
}

func fixtureWrap(t *testing.T, entry interface{}) (uint, error) {
	err := fixture.CleanUp(db, additionalTables, &User{}, &Exam{}, &News{}, &Token{})
	if err != nil {
		t.Fatal(err)
	}
	err = db.Create(entry).Error
	return testExam.ID, err
}
