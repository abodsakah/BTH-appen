package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	additionalTables = []string{"exam_users"}
)

var testUser = &User{
	Name:     "Test Testsson",
	Username: "test",
	Password: "pass",
	Role:     "student",
}

var testExam = &Exam{
  Name: "test",
  CourseCode: "pa121212",
  StartDate: time.Now(),
}

func createUserWrap() (uint ,error) {
	temp := *testUser
	err := CreateUser(db, testUser)
	*testUser = temp
	if err != nil {
		return 0, err
	}
  user, _ := GetUserByName(db, testUser.Username)
	return user.ID ,nil
}

func fixtureWrapUser(t *testing.T) (uint, error) {
	err := cleanUp(db, additionalTables)
	if err != nil {
		t.Fatal(err)
	}
  id , err := createUserWrap()
	return id, err
}

func fixtureWrapExam(t *testing.T) (uint, error) {

	err := cleanUp(db, additionalTables)
	if err != nil {
		t.Fatal(err)
	}
  err = CreateExam(db, testExam)
  id, _ := getExamByName(db, testExam.Name)
	return uint(id), err
}

func assertNoError(t *testing.T, err error, message string) {
	assert.Equal(t, nil, err, message)
}

func assertError(t *testing.T, err error, message string) {
	assert.NotEqual(t, nil, err, message)
}
