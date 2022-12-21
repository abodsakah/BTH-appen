// Unit tests for route module
package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	db "github.com/abodsakah/BTH-appen/backend/api/src/DB"
	helpers "github.com/abodsakah/BTH-appen/backend/api/src/Helpers"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func startTest(t *testing.T) {
	_, _ = helpers.FixtureWrapCreate(t, helpers.TestExam, helpers.TestNews)
	temp := *helpers.TestUser
	_ = db.CreateUser(helpers.DbGorm, &temp)
}

func setupContext() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// make sure c.Request is not nil
	c.Request = &http.Request{
		Header: make(http.Header),
	}

	return c, w
}

func mockJSONPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func TestDatabaseRoutes(t *testing.T) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		t.Fatal(err)
	}
	dbP, err := db.SetupDatabase()
	if dbP == nil {
		fmt.Printf("Is nil!!!!!!!!!!!!!!!!!!!!!1111\n")
	}
	assert.Nil(t, err, "Database can not be connected to")
	helpers.DbGorm = dbP
  _ = helpers.FixtureWrapNonCreate(t)
}

func TestSetupRoutes(t *testing.T) {
	var err error
	router, err = SetupRoutes(helpers.DbGorm)
	assert.Nil(t, err, "Trying to set up routes shall not cause any errors")
}

func TestLogin1(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"Username": "test",
		"Password": "pass",
	})

	// call API endpoint
	login(c)
	assert.Equal(t, 200, w.Code, "When trying to log in with correct admin credentials it shall return status: 200")
}

func TestLogin2(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"Username": "admin",
		"Password": "incorrect-password",
	})

	// call API endpoint
	login(c)
	assert.NotEqual(t, 200, w.Code, "When trying to log in with incorrect admin credentials it shall not return status: 200")
}

func TestExamCreate1(t *testing.T) {
	_ = helpers.FixtureWrapNonCreate(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"name":        "Test exam",
		"course_code": "something",
		"room":        "something",
		"start_date":  time.Now(),
	})

	// call API endpoint
	createExam(c)
	assert.Equal(t, 200, w.Code, "When trying to call on create Exam API with no duplicates present, it should return status: 200")
}

func TestExamCreate2(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"name":        "Test exam",
		"course_code": "something",
		"room":        "something",
		"start_date":  time.Now(),
	})

	// call API endpoint
	createExam(c)
	assert.NotEqual(t, 200, w.Code, "When trying to call on create Exam API with duplicates present, it should not return status: 200")
}

func TestExamDelete1(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"exam_id": 1,
	})

	// call API endpoint
	deleteExam(c)
	assert.Equal(t, 200, w.Code, "When trying to call on delete Exam API with entry present, it should return status: 200")
}

func TestExamDelete2(t *testing.T) {
	_ = helpers.FixtureWrapNonCreate(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"exam_id": 1,
	})

	// call API endpoint
	deleteExam(c)
	assert.NotEqual(t, 200, w.Code, "When trying to call on delete Exam API with no entry present, it should not return status: 200")
}

func TestListExams1(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"exam_id": 1,
	})

	// call API endpoint
	listExams(c)
	assert.Equal(t, 200, w.Code, "When trying to call on list Exams API with entries present, it should return status: 200")
}

func TestListDueExams1(t *testing.T) {
	startTest(t)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"exam_id": 1,
	})

	// call API endpoint
	listDueExams(c)
	assert.Equal(t, 200, w.Code, "When trying to call on list due Exams API with due exams present, it should return status: 200")
}

func TestListUserExams(t *testing.T) {
	startTest(t)
	_, _ = db.AddUserToExam(helpers.DbGorm, 1, 1)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"userID":  1,
		"exam_id": 1,
	})
	c.Set("UserID", uint(1))
	// call API endpoint
	listUserExams(c)
	assert.Equal(t, 200, w.Code, "When trying to call on list User Exams API with exams present for user, it should return status: 200")
}

func TestListExamUser(t *testing.T) {
	startTest(t)
	_, _ = db.AddUserToExam(helpers.DbGorm, 1, 1)
	c, w := setupContext()
	// Set Body, Header and Content-Type
	mockJSONPost(c, &gin.H{
		"exam_id": 1,
	})
	c.Set("UserID", uint(1))
	// call API endpoint
	listExamUsers(c)
	assert.Equal(t, 200, w.Code, "When trying to call on list Exam Users API with users present for user, it should return status: 200")
}
