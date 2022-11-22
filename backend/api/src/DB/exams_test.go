// Package db provides db
package db

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type TestSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock
}
type Student struct {
	ID   uint
	Name string
}

func TestExam_Create_success(t *testing.T) {
	// assert equality
	assert := assert.New(t)
	// setup mock gorm connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db, // mocked database object
		PreferSimpleProtocol: true,
	})

	mockGorm, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	// test
	// exam := Exam{
	// 	Name:       "Some course name",
	// 	CourseCode: "DV1337",
	// 	StartDate:  time.Now().AddDate(0, 0, 5),
	// }
	student := Student{
		ID:   1234,
		Name: "Test Testsson",
	}

	// TODO: test
	mock.ExpectBegin()
	// mock.ExpectExec("INSERT INTO exams").WithArgs(exam.Name, exam.CourseCode, exam.StartDate).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectQuery(`"INSERT INTO "students" ("name","id") VALUES ($1,$2) RETURNING "id""`).
		WithArgs(student.ID, student.Name).WillReturnRows(sqlmock.NewRows([]string{"id"}).
		AddRow(student.ID))
	mock.ExpectCommit()

	mockGorm.Create(&student)
	// resError := CreateExam(mockGorm, &exam)

	assert.Equal(nil, mock.ExpectationsWereMet(), "Error should be nil")
}

func TestExam_Create_failure(t *testing.T) {
	// assert equality
	assert := assert.New(t)
	// setup mock gorm connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_1",
		DriverName:           "postgres",
		Conn:                 db, // mocked database object
		PreferSimpleProtocol: true,
	})

	mockGorm, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}

	// test
	exam := Exam{
		CourseCode: "DV1337",
		StartDate:  time.Now().AddDate(0, 0, 5),
	}

	// TODO: test
	mock.ExpectBegin()
	mock.ExpectCommit()

	resError := CreateExam(mockGorm, &exam)

	assert.Error(resError, "Error should not be nil")
}
