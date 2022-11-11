package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type dbExamMock struct {
	mock.Mock
}

func (m *dbExamMock) Create(exam *Exam) *gorm.DB {
	args := m.Called(exam)
	return args.Get(0).(*gorm.DB)
}

func TestExam_Create_success(t *testing.T) {
	// assert equality
	assert := assert.New(t)
	dbMock := new(dbExamMock)
	var exam Exam
	var result gorm.DB
	result.Error = nil
	dbMock.On("Create", &exam).Return(&result)
	res := CreateExam(&exam)
	assert.Equal(res, nil, "")
}

func TestExam_Create_failure(t *testing.T) {
	// assert equality
	assert := assert.New(t)
	dbMock := new(dbExamMock)
	var exam *Exam
	var result *gorm.DB
	result.Error = db.Error
	dbMock.On("Create", exam).Return(result)
	res := CreateExam(exam)
	assert.Equal(res, nil, "")
}
