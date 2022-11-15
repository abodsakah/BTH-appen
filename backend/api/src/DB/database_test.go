// Package db provides db
package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dbMock struct {
	DB mock.Mock
}

func TestExam_Create_success(t *testing.T) {
	// assert equality
	assert := assert.New(t)
	// setup mock gorm connection
	// TODO: get a mocked database object here.
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
	exam := Exam{
		CourseCode: "DV1337",
		StartDate:  time.Now().AddDate(0, 0, 5),
	}

	// TODO: test

	resError := CreateExam(mockGorm, &exam)

	assert.Equal(nil, resError, "Error should be nil")
}

func TestExam_Create_failure(t *testing.T) {
	// assert equality
	assert := assert.New(t)
	// setup mock gorm connection
	// TODO: get a mocked database object here.
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

	resError := CreateExam(mockGorm, &exam)

	assert.Error(resError, "Error should not be nil")
}
